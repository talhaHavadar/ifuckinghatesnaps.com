package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"testing/fstest"

	"ifuckinghatesnaps.com/internal/site"
)

// With FromDisk, an edit to a template or a static file shows on the next
// request, without building a new handler.
func TestFromDiskPicksUpEditsWithoutRestart(t *testing.T) {
	dir := t.TempDir()
	if err := os.CopyFS(dir, os.DirFS(filepath.Join("..", "internal", "site"))); err != nil {
		t.Fatal(err)
	}
	handler, err := site.New(fstest.MapFS{"1.txt": {Data: []byte("first")}}, site.FromDisk(dir))
	if err != nil {
		t.Fatalf("site.New: %v", err)
	}
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	before := get(t, srv, "/")
	assertStatus(t, "/", before, http.StatusOK)
	assertContains(t, "/", before, "Whack a snap.")

	landing := filepath.Join(dir, "templates", "landing.html")
	page, err := os.ReadFile(landing)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(landing, []byte(strings.Replace(string(page), "Whack a snap.", "Whack it harder.", 1)), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "static", "whack.js"), []byte("// edited\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	after := get(t, srv, "/")
	assertContains(t, "/", after, "Whack it harder.")

	// The script's version changes with its content, so browsers fetch it again.
	script := regexp.MustCompile(`/static/whack\.js\?v=[0-9a-f]+`)
	if script.FindString(before.body) == script.FindString(after.body) {
		t.Fatalf("GET /: whack.js version did not change after the file changed")
	}
	assertContains(t, "/static/whack.js", get(t, srv, "/static/whack.js"), "// edited")
}
