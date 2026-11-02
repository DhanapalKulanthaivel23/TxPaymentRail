package presentation

import (
	"bytes"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHostedAndEmbeddedClientFilesMatch(t *testing.T) {
	for _, name := range []string{"index.html", "app.js", "style.css", "connection.mjs", "graph-view.mjs", "client-config.json"} {
		published, e := os.ReadFile("../docs/" + name)
		if e != nil {
			t.Fatal(e)
		}
		embedded, e := assets.ReadFile("assets/" + name)
		if e != nil || !bytes.Equal(published, embedded) {
			t.Fatal("client divergence", name)
		}
	}
}
func TestPublicServerDoesNotExposeCoreSource(t *testing.T) {
	for _, path := range []string{"/catalog", "/execute", "/mcp", "/.env", "/AGENTS.md", "/server.py"} {
		w := httptest.NewRecorder()
		New().ServeHTTP(w, httptest.NewRequest("GET", path, nil))
		if w.Code != 404 {
			t.Fatal(path, w.Code)
		}
	}
}
