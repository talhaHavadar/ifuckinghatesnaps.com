// Package site serves the quotes over HTTP.
package site

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"strconv"
)

const supportURL = "https://buymeacoffee.com/uphoenix"

//go:embed templates/*.html
var templateFS embed.FS

type server struct {
	quotes collection
	pages  *template.Template
}

type quotePage struct {
	Quote      quote
	PreviousID int
	SupportURL string
}

type notFoundPage struct {
	SupportURL string
}

// New returns the site's HTTP handler. It returns an error if any file in
// quotes is not a valid quote, so a broken quote stops the server at startup.
func New(quotes fs.FS) (http.Handler, error) {
	c, err := loadQuotes(quotes)
	if err != nil {
		return nil, err
	}
	pages, err := template.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("parsing templates: %w", err)
	}

	s := &server{quotes: c, pages: pages}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", s.handleLatest)
	mux.HandleFunc("GET /{id}/{$}", s.handleQuote)
	mux.HandleFunc("GET /{id}", s.handleBareID)
	mux.HandleFunc("GET /random", s.handleRandom)
	mux.HandleFunc("GET /healthz", handleHealth)
	mux.HandleFunc("GET /", s.handleNotFound)
	return mux, nil
}

func (s *server) handleLatest(w http.ResponseWriter, r *http.Request) {
	s.renderQuote(w, s.quotes.latest())
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
	s.render(w, http.StatusNotFound, "notfound.html", notFoundPage{SupportURL: supportURL})
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
	s.render(w, http.StatusOK, "quote.html", quotePage{
		Quote:      q,
		PreviousID: s.quotes.previous(q.ID),
		SupportURL: supportURL,
	})
}

func (s *server) render(w http.ResponseWriter, status int, name string, data any) {
	var buf bytes.Buffer
	if err := s.pages.ExecuteTemplate(&buf, name, data); err != nil {
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
