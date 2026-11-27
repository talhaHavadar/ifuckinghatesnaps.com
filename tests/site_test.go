package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"testing/fstest"

	"ifuckinghatesnaps.com/internal/site"
)

type response struct {
	status int
	body   string
	header http.Header
}

func serve(t *testing.T, files map[string]string) *httptest.Server {
	t.Helper()
	fsys := fstest.MapFS{}
	for name, content := range files {
		fsys[name] = &fstest.MapFile{Data: []byte(content)}
	}
	handler, err := site.New(fsys)
	if err != nil {
		t.Fatalf("site.New: %v", err)
	}
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return srv
}

func request(t *testing.T, srv *httptest.Server, method, path string) response {
	t.Helper()
	client := &http.Client{
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}
	req, err := http.NewRequest(method, srv.URL+path, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("%s %s: %v", method, path, err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	return response{status: resp.StatusCode, body: string(body), header: resp.Header}
}

func get(t *testing.T, srv *httptest.Server, path string) response {
	t.Helper()
	return request(t, srv, http.MethodGet, path)
}

func assertStatus(t *testing.T, path string, got response, want int) {
	t.Helper()
	if got.status != want {
		t.Fatalf("GET %s: status %d, want %d", path, got.status, want)
	}
}

func assertContains(t *testing.T, path string, got response, want string) {
	t.Helper()
	if !strings.Contains(got.body, want) {
		t.Fatalf("GET %s: body does not contain %q:\n%s", path, want, got.body)
	}
}

func TestQuotePageHighlightsMarkedText(t *testing.T) {
	cases := []struct {
		name  string
		quote string
		want  []string
	}{
		{
			name:  "single highlight",
			quote: "Snaps **auto-refresh** whenever they like.",
			want:  []string{"<p>Snaps <mark>auto-refresh</mark> whenever they like.</p>"},
		},
		{
			name:  "several highlights in one paragraph",
			quote: "**Slow** to start and **slow** to update.",
			want:  []string{"<p><mark>Slow</mark> to start and <mark>slow</mark> to update.</p>"},
		},
		{
			name:  "no highlight",
			quote: "Just sad.",
			want:  []string{"<p>Just sad.</p>"},
		},
		{
			name:  "paragraphs split on blank lines",
			quote: "First.\n\n**Second.**\r\n\r\nThird.",
			want:  []string{"<p>First.</p>", "<p><mark>Second.</mark></p>", "<p>Third.</p>"},
		},
		{
			name:  "html in the quote is escaped",
			quote: `<script>alert(1)</script> **<b>"hi"</b>**`,
			// Straight quotes become curly ones before escaping (see the
			// typography case below), so "hi" comes out as “hi”.
			want: []string{"<p>&lt;script&gt;alert(1)&lt;/script&gt; <mark>&lt;b&gt;“hi”&lt;/b&gt;</mark></p>"},
		},
		{
			name:  "straight quotes and apostrophes become curly",
			quote: `"It's 'fine'," he said. **Don't.**`,
			want:  []string{"<p>“It’s ‘fine’,” he said. <mark>Don’t.</mark></p>"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := serve(t, map[string]string{"1.txt": tc.quote})
			got := get(t, srv, "/1/")
			assertStatus(t, "/1/", got, http.StatusOK)
			for _, want := range tc.want {
				assertContains(t, "/1/", got, want)
			}
		})
	}
}

func TestInvalidQuotesStopTheSite(t *testing.T) {
	cases := []struct {
		name  string
		files map[string]string
	}{
		{"no quotes", map[string]string{}},
		{"empty quote", map[string]string{"1.txt": " \n\n "}},
		{"unclosed highlight", map[string]string{"1.txt": "I **hate snaps."}},
		{"highlight across paragraphs", map[string]string{"1.txt": "I **hate\n\nsnaps**."}},
		{"empty highlight", map[string]string{"1.txt": "I **** snaps."}},
		{"file name is not a number", map[string]string{"one.txt": "Snaps."}},
		{"wrong extension", map[string]string{"1.md": "Snaps."}},
		{"quote number zero", map[string]string{"0.txt": "Snaps."}},
		{"same number twice", map[string]string{"1.txt": "Snaps.", "001.txt": "Snaps again."}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fsys := fstest.MapFS{}
			for name, content := range tc.files {
				fsys[name] = &fstest.MapFile{Data: []byte(content)}
			}
			if _, err := site.New(fsys); err == nil {
				t.Fatal("site.New accepted invalid quotes")
			}
		})
	}
}

// The home page used to show the newest quote. It is now the whack-a-snap
// game, and shows no quote at all.
func TestHomePageIsTheWhackASnapGame(t *testing.T) {
	srv := serve(t, map[string]string{
		"0001.txt": "first",
		"0010.txt": "tenth",
	})
	got := get(t, srv, "/")
	assertStatus(t, "/", got, http.StatusOK)
	assertContains(t, "/", got, "<h1>I fucking hate snaps.</h1>")
	assertContains(t, "/", got, "Whack a snap.")
	assertContains(t, "/", got, `<script src="/static/whack.js?v=`)
	// The closing sentence is the page's footer landmark, and hits are read
	// out to screen readers through a status region that starts empty.
	assertContains(t, "/", got, `<footer class="report report-band">`)
	assertContains(t, "/", got, `<p class="sr-only" role="status" id="whack-status"></p>`)
	for _, quote := range []string{"<blockquote", "first", "tenth"} {
		if strings.Contains(got.body, quote) {
			t.Fatalf("GET /: body contains %q, want no quote on the home page", quote)
		}
	}
	// Without JavaScript one snap is up, and it is a plain link to a random
	// quote. The two snaps that are down cannot be reached with the keyboard.
	if n := strings.Count(got.body, `<a class="snap" href="/random"`); n != 1 {
		t.Fatalf("GET /: %d snaps are up, want 1", n)
	}
	if n := strings.Count(got.body, `<a class="snap down rest" href="/random" aria-label="Whack a snap: open a random opinion" tabindex="-1" aria-hidden="true">`); n != 2 {
		t.Fatalf("GET /: %d snaps are down and out of the tab order, want 2", n)
	}
}

func TestHomePageAssetsAreServed(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "first"})
	home := get(t, srv, "/")
	paths := regexp.MustCompile(`(/static/[^"\s,)]+)`).FindAllStringSubmatch(home.body, -1)
	if len(paths) == 0 {
		t.Fatal("GET /: references no /static/ files")
	}
	for _, m := range paths {
		path := m[1]
		got := get(t, srv, path)
		assertStatus(t, path, got, http.StatusOK)
		if got.body == "" {
			t.Fatalf("GET %s: empty body", path)
		}
		// The page links to each file with a version taken from its
		// content, so the browser may keep it for good.
		if !strings.Contains(path, "?v=") {
			t.Fatalf("GET /: %s has no version", path)
		}
		if cc := got.header.Get("Cache-Control"); !strings.Contains(cc, "immutable") {
			t.Fatalf("GET %s: Cache-Control %q, want immutable", path, cc)
		}
	}
	wantType := map[string]string{
		"/static/whack.js":                       "text/javascript",
		"/static/snap-box.webp":                  "image/webp",
		"/static/figtree-latin-800-normal.woff2": "font/woff2",
	}
	for path, want := range wantType {
		got := get(t, srv, path)
		if ct := got.header.Get("Content-Type"); !strings.HasPrefix(ct, want) {
			t.Fatalf("GET %s: Content-Type %q, want %s", path, ct, want)
		}
		// Without a version the file may change under the same URL.
		if cc := got.header.Get("Cache-Control"); strings.Contains(cc, "immutable") {
			t.Fatalf("GET %s: Cache-Control %q, want no immutable without a version", path, cc)
		}
	}
}

func TestQuotePageLinksToPreviousQuote(t *testing.T) {
	srv := serve(t, map[string]string{
		"1.txt": "first",
		"2.txt": "second",
		"5.txt": "fifth",
	})
	cases := []struct {
		path, wantPrev string
	}{
		{"/5/", "/2/"},
		{"/2/", "/1/"},
		{"/1/", "/5/"},
	}
	for _, tc := range cases {
		got := get(t, srv, tc.path)
		assertStatus(t, tc.path, got, http.StatusOK)
		// The page links its previous quote for browsers and crawlers; the
		// visible links are "Add your own rant" and "Another opinion".
		assertContains(t, tc.path, got, `<link rel="prev" href="`+tc.wantPrev+`">`)
	}
}

func TestUnknownPathsAreNotFound(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "first", "3.txt": "third"})
	for _, path := range []string{"/2/", "/4/", "/01/", "/-1/", "/abc/", "/1/extra", "/static/x.css"} {
		got := get(t, srv, path)
		assertStatus(t, path, got, http.StatusNotFound)
		assertContains(t, path, got, "No opinion here.")
	}
}

func TestQuoteURLWithoutTrailingSlashRedirects(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "first", "7.txt": "seventh"})

	got := get(t, srv, "/7")
	assertStatus(t, "/7", got, http.StatusMovedPermanently)
	if got.header.Get("Location") != "/7/" {
		t.Fatalf("GET /7: Location %q, want /7/", got.header.Get("Location"))
	}

	assertStatus(t, "/8", get(t, srv, "/8"), http.StatusNotFound)
}

func TestRandomRedirectsToExistingQuote(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "a", "2.txt": "b", "9.txt": "c"})
	valid := map[string]bool{"/1/": true, "/2/": true, "/9/": true}
	for range 30 {
		got := get(t, srv, "/random")
		assertStatus(t, "/random", got, http.StatusFound)
		if cc := got.header.Get("Cache-Control"); cc != "no-store" {
			t.Fatalf("GET /random: Cache-Control %q, want no-store", cc)
		}
		if !valid[got.header.Get("Location")] {
			t.Fatalf("GET /random: redirected to %q, not an existing quote", got.header.Get("Location"))
		}
	}
}

func TestPagesLinkToBuyMeACoffee(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "first"})
	for _, path := range []string{"/", "/1/", "/missing/"} {
		assertContains(t, path, get(t, srv, path), `href="https://buymeacoffee.com/uphoenix"`)
	}
}

func TestHealthz(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "first"})
	got := get(t, srv, "/healthz")
	assertStatus(t, "/healthz", got, http.StatusOK)
}

func TestOnlyGetIsAllowed(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "first"})
	for _, path := range []string{"/", "/1/", "/random"} {
		got := request(t, srv, http.MethodPost, path)
		if got.status != http.StatusMethodNotAllowed {
			t.Fatalf("POST %s: status %d, want %d", path, got.status, http.StatusMethodNotAllowed)
		}
	}
}

// A quote page leads with its punchline, keeps the whole quote for context,
// and carries what a link preview needs: its own title, the quote as the
// description, a canonical URL and an image.
func TestQuotePageIsShareable(t *testing.T) {
	srv := serve(t, map[string]string{
		"1.txt": "I delete it, it comes back. **It's like a cockroach.**",
		"2.txt": "No highlight here, just pain.",
	})

	got := get(t, srv, "/1/")
	assertStatus(t, "/1/", got, http.StatusOK)
	for _, want := range []string{
		`<h1 class="pull pull-quoted">It’s like a cockroach.</h1>`,
		`<blockquote id="quote-1">`,
		`<title>“It’s like a cockroach.” · I fucking hate snaps.</title>`,
		`<meta property="og:title" content="“It’s like a cockroach.” · I fucking hate snaps.">`,
		`<meta property="og:description" content="I delete it, it comes back. It’s like a cockroach.">`,
		`<link rel="canonical" href="https://ifuckinghatesnaps.com/1/">`,
		`<meta property="og:image" content="https://ifuckinghatesnaps.com/static/og.jpg?v=`,
		// The share button needs the script; without it the link is shown as text.
		`<button class="share" type="button" data-url="https://ifuckinghatesnaps.com/1/"`,
		`<span class="permalink">ifuckinghatesnaps.com/1/</span>`,
		`<footer class="report report-band">`,
		`<a href="https://buymeacoffee.com/uphoenix">Add your own rant</a>`,
		`<a href="/random">Another opinion</a>`,
	} {
		assertContains(t, "/1/", got, want)
	}

	// A quote with no highlight has no punchline to pull out: it is shown on
	// its own, and the page still has a heading for screen readers.
	plain := get(t, srv, "/2/")
	assertContains(t, "/2/", plain, `<h1 class="sr-only">Opinion No. 2</h1>`)
	assertContains(t, "/2/", plain, `<blockquote id="quote-2" class="solo">`)
	assertContains(t, "/2/", plain, `<title>Opinion No. 2 · I fucking hate snaps.</title>`)
}

func TestQuoteAndNotFoundPageAssetsAreServed(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "**first**"})
	for _, page := range []string{"/1/", "/missing/"} {
		body := get(t, srv, page).body
		paths := regexp.MustCompile(`(/static/[^"\s,)]+)`).FindAllStringSubmatch(body, -1)
		if len(paths) == 0 {
			t.Fatalf("GET %s: references no /static/ files", page)
		}
		for _, m := range paths {
			assertStatus(t, m[1], get(t, srv, m[1]), http.StatusOK)
		}
	}
	notFound := get(t, srv, "/missing/")
	assertContains(t, "/missing/", notFound, `<h1 class="pull">No opinion here.</h1>`)
	assertContains(t, "/missing/", notFound, `<a class="pill" href="/random">Random opinion</a>`)
}

// Every page follows the system's light or dark setting, and a picker lets
// the visitor choose. A saved choice is applied in the head, before the page
// paints, so a reload does not flash the other theme.
func TestPagesOfferATheme(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "**first**"})
	for _, path := range []string{"/", "/1/", "/missing/"} {
		got := get(t, srv, path)
		assertContains(t, path, got, `<meta name="color-scheme" content="light dark">`)
		assertContains(t, path, got, `localStorage.getItem("theme")`)
		assertContains(t, path, got, `<button class="theme" type="button" data-choice="system"`)
		assertContains(t, path, got, `<script src="/static/theme.js?v=`)
		head, _, _ := strings.Cut(got.body, "<body>")
		if !strings.Contains(head, `localStorage.getItem("theme")`) {
			t.Fatalf("GET %s: the saved theme is not applied in the head", path)
		}
	}
}
