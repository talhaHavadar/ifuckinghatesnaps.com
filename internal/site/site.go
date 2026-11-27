// Package site serves the quotes over HTTP.
package site

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	supportURL = "https://buymeacoffee.com/uphoenix"
	// siteURL is the public origin, for canonical links and link previews,
	// which must be absolute.
	siteURL  = "https://ifuckinghatesnaps.com"
	siteName = "I fucking hate snaps."
)

// assets holds the page templates (templates/) and the files served under
// /static/ (static/).
//
//go:embed templates/*.html static/*.js static/*.css static/*.png static/*.jpg static/*.webp static/*.woff2 static/*.txt
var assets embed.FS

type server struct {
	quotes collection
	static fs.FS
	// pages returns the parsed templates: the ones parsed at startup, or with
	// FromDisk, freshly parsed ones.
	pages func() (*template.Template, error)
}

// An Option changes how New builds the site.
type Option func(*options)

type options struct {
	assets fs.FS
	reload bool
}

// FromDisk serves templates and static files from dir, which holds the same
// templates/ and static/ directories as internal/site, and reads them again
// on every request. It is for development: a template edit shows on the next
// page load without restarting the server.
func FromDisk(dir string) Option {
	return func(o *options) {
		o.assets = os.DirFS(dir)
		o.reload = true
	}
}

// meta is what every page's head and link preview needs.
type meta struct {
	Title       string
	Description string
	// URL is the page's absolute canonical URL; empty for pages that have
	// none, such as the 404.
	URL string
	// Prev is the previous quote's path, for a rel="prev" link; empty when
	// the page is not a quote.
	Prev string
}

type quotePage struct {
	meta
	Quote quote
	// Link is the quote's URL as shown on the page, without the scheme.
	Link       string
	SupportURL string
}

type notFoundPage struct {
	meta
	SupportURL string
}

type landingPage struct {
	meta
	// Holes says which of the three holes shows the snap before the game
	// script runs, and when it never does.
	Holes      []bool
	SupportURL string
}

// New returns the site's HTTP handler. It returns an error if any file in
// quotes is not a valid quote, so a broken quote stops the server at startup.
func New(quotes fs.FS, opts ...Option) (http.Handler, error) {
	o := options{assets: assets}
	for _, opt := range opts {
		opt(&o)
	}
	c, err := loadQuotes(quotes)
	if err != nil {
		return nil, err
	}
	static, err := fs.Sub(o.assets, "static")
	if err != nil {
		return nil, fmt.Errorf("opening static files: %w", err)
	}
	pages, err := parsePages(o.assets, static)
	if err != nil {
		return nil, err
	}

	s := &server{quotes: c, static: static}
	s.pages = func() (*template.Template, error) { return pages, nil }
	if o.reload {
		s.pages = func() (*template.Template, error) { return parsePages(o.assets, static) }
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", s.handleLanding)
	mux.HandleFunc("GET /{id}/{$}", s.handleQuote)
	mux.HandleFunc("GET /{id}", s.handleBareID)
	mux.HandleFunc("GET /random", s.handleRandom)
	mux.HandleFunc("GET /static/{file}", s.handleStatic)
	mux.HandleFunc("GET /healthz", handleHealth)
	mux.HandleFunc("GET /", s.handleNotFound)
	return mux, nil
}

func (s *server) handleLanding(w http.ResponseWriter, r *http.Request) {
	s.render(w, http.StatusOK, "landing.html", landingPage{
		meta: meta{
			Title:       siteName,
			Description: "Real opinions from real people about Ubuntu snaps. Whack a snap to read one.",
			URL:         siteURL + "/",
		},
		Holes:      []bool{false, true, false},
		SupportURL: supportURL,
	})
}

func (s *server) handleStatic(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("file")
	if info, err := fs.Stat(s.static, name); err != nil || info.IsDir() {
		s.handleNotFound(w, r)
		return
	}
	// Pages link to static files with a version taken from their content, so
	// a versioned URL never changes and can be cached for good.
	if r.URL.Query().Has("v") {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	} else {
		w.Header().Set("Cache-Control", "public, max-age=3600")
	}
	http.ServeFileFS(w, r, s.static, name)
}

// parsePages parses the templates and gives them a static function that
// returns a file's versioned URL.
func parsePages(assets, static fs.FS) (*template.Template, error) {
	urls, err := staticURLs(static)
	if err != nil {
		return nil, err
	}
	funcs := template.FuncMap{
		// abs makes a site path absolute, for link previews.
		"abs": func(path string) string { return siteURL + path },
		"static": func(name string) (string, error) {
			u, ok := urls[name]
			if !ok {
				return "", fmt.Errorf("no static file %q", name)
			}
			return u, nil
		},
	}
	pages, err := template.New("").Funcs(funcs).ParseFS(assets, "templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("parsing templates: %w", err)
	}
	return pages, nil
}

// staticURLs maps each static file name to its URL, versioned by a hash of
// its content so that browsers fetch it again when it changes.
func staticURLs(static fs.FS) (map[string]string, error) {
	entries, err := fs.ReadDir(static, ".")
	if err != nil {
		return nil, fmt.Errorf("listing static files: %w", err)
	}
	urls := make(map[string]string, len(entries))
	for _, e := range entries {
		data, err := fs.ReadFile(static, e.Name())
		if err != nil {
			return nil, fmt.Errorf("reading static file: %w", err)
		}
		sum := sha256.Sum256(data)
		urls[e.Name()] = "/static/" + e.Name() + "?v=" + hex.EncodeToString(sum[:4])
	}
	return urls, nil
}

func (s *server) handleQuote(w http.ResponseWriter, r *http.Request) {
	q, ok := s.lookup(r.PathValue("id"))
	if !ok {
		s.handleNotFound(w, r)
		return
	}
	s.renderQuote(w, q)
}

func (s *server) handleBareID(w http.ResponseWriter, r *http.Request) {
	q, ok := s.lookup(r.PathValue("id"))
	if !ok {
		s.handleNotFound(w, r)
		return
	}
	http.Redirect(w, r, quotePath(q.ID), http.StatusMovedPermanently)
}

func (s *server) handleRandom(w http.ResponseWriter, r *http.Request) {
	id := s.quotes.ids[rand.IntN(len(s.quotes.ids))]
	w.Header().Set("Cache-Control", "no-store")
	http.Redirect(w, r, quotePath(id), http.StatusFound)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok\n"))
}

func (s *server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	s.render(w, http.StatusNotFound, "notfound.html", notFoundPage{
		meta: meta{
			Title:       "No opinion here · " + siteName,
			Description: "Real opinions from real people about Ubuntu snaps.",
		},
		SupportURL: supportURL,
	})
}

// lookup accepts only the canonical form of a quote number, so /05/ is a 404
// and every quote has exactly one URL.
func (s *server) lookup(raw string) (quote, bool) {
	id, err := strconv.Atoi(raw)
	if err != nil || strconv.Itoa(id) != raw {
		return quote{}, false
	}
	q, ok := s.quotes.byID[id]
	return q, ok
}

func (s *server) renderQuote(w http.ResponseWriter, q quote) {
	title := "Opinion No. " + strconv.Itoa(q.ID) + " · " + siteName
	if q.Punchline != "" {
		title = "\u201c" + q.Punchline + "\u201d · " + siteName
	}
	s.render(w, http.StatusOK, "quote.html", quotePage{
		meta: meta{
			Title:       title,
			Description: q.Text,
			URL:         siteURL + quotePath(q.ID),
			Prev:        quotePath(s.quotes.previous(q.ID)),
		},
		Quote:      q,
		Link:       strings.TrimPrefix(siteURL, "https://") + quotePath(q.ID),
		SupportURL: supportURL,
	})
}

func (s *server) render(w http.ResponseWriter, status int, name string, data any) {
	pages, err := s.pages()
	if err != nil {
		slog.Error("loading templates", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	var buf bytes.Buffer
	if err := pages.ExecuteTemplate(&buf, name, data); err != nil {
		slog.Error("rendering page", "template", name, "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, _ = buf.WriteTo(w)
}

func quotePath(id int) string {
	return "/" + strconv.Itoa(id) + "/"
}
