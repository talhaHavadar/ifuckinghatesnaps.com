package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ifuckinghatesnaps.com/internal/site"
	"ifuckinghatesnaps.com/quotes"
)

func TestPublishedQuotesAreValid(t *testing.T) {
	handler, err := site.New(quotes.FS)
	if err != nil {
		t.Fatalf("a file in quotes/ is invalid: %v", err)
	}
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	assertStatus(t, "/", get(t, srv, "/"), http.StatusOK)
}
