package base

import (
	"github.com/DhanapalKulanthaivel23/addressrail-ai/controller"
	"net/http"
	"os"
	"path/filepath"
)

func New() http.Handler {
	api := controller.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assets := map[string]string{"/": "index.html", "/index.html": "index.html", "/app.js": "app.js", "/style.css": "style.css", "/config.json": "config.json"}
		if name, ok := assets[r.URL.Path]; ok {
			if r.Method != "GET" && r.Method != "HEAD" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			dir := os.Getenv("UI_DIR")
			if dir == "" {
				dir = "docs"
			}
			w.Header().Set("Cache-Control", "no-store")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Referrer-Policy", "no-referrer")
			w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; connect-src 'self' https: http://localhost:* http://127.0.0.1:*; frame-ancestors 'none'; base-uri 'none'; form-action 'self'")
			http.ServeFile(w, r, filepath.Join(dir, name))
			return
		}
		api.ServeHTTP(w, r)
	})
}
