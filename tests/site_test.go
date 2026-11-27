package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"ifuckinghatesnaps.com/internal/site"
)

type response struct {
	status   int
	body     string
	location string
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
	return response{status: resp.StatusCode, body: string(body), location: resp.Header.Get("Location")}
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
			want:  []string{"<p>&lt;script&gt;alert(1)&lt;/script&gt; <mark>&lt;b&gt;&#34;hi&#34;&lt;/b&gt;</mark></p>"},
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

func TestHomePageShowsLatestQuote(t *testing.T) {
	srv := serve(t, map[string]string{
		"0001.txt": "first",
		"0002.txt": "second",
		"0010.txt": "tenth",
	})
	got := get(t, srv, "/")
	assertStatus(t, "/", got, http.StatusOK)
	assertContains(t, "/", got, `id="quote-10"`)
	assertContains(t, "/", got, "<p>tenth</p>")
}

func TestAnotherOpinionLinksToPreviousQuote(t *testing.T) {
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
		assertContains(t, tc.path, got, `<a href="`+tc.wantPrev+`" rel="prev">`)
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
	if got.location != "/7/" {
		t.Fatalf("GET /7: Location %q, want /7/", got.location)
	}

	assertStatus(t, "/8", get(t, srv, "/8"), http.StatusNotFound)
}

func TestRandomRedirectsToExistingQuote(t *testing.T) {
	srv := serve(t, map[string]string{"1.txt": "a", "2.txt": "b", "9.txt": "c"})
	valid := map[string]bool{"/1/": true, "/2/": true, "/9/": true}
	for range 30 {
		got := get(t, srv, "/random")
		assertStatus(t, "/random", got, http.StatusFound)
		if !valid[got.location] {
			t.Fatalf("GET /random: redirected to %q, not an existing quote", got.location)
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
