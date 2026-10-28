package base

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConnectionUI(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte("connection page"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("UI_DIR", dir)
	h := New()
	r := httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/", nil))
	if r.Code != 200 || !strings.Contains(r.Body.String(), "connection page") {
		t.Fatalf("UI: %d %s", r.Code, r.Body.String())
	}
	if r.Header().Get("Referrer-Policy") != "no-referrer" {
		t.Fatal("missing referrer policy")
	}
	r = httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/README.md", nil))
	if r.Code != 404 {
		t.Fatal("unexpected public file exposure")
	}
	r = httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest(http.MethodPost, "/", nil))
	if r.Code != 405 {
		t.Fatal("unexpected UI mutation")
	}
	r = httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/health", nil))
	if r.Code != 200 {
		t.Fatal("synthetic API unavailable")
	}
}
