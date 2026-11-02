package presentation

import (
	"embed"
	"encoding/json"
	"net/http"
	"os"
)

//go:embed assets/*
var assets embed.FS

// New serves the same client files as GitHub Pages. Core requests are made by the browser to its configured API.
func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; connect-src 'self' https: http://localhost:* http://127.0.0.1:*; frame-ancestors 'none'; base-uri 'none'")
		if r.Method != "GET" && r.Method != "HEAD" {
			if r.URL.Path == "/api/demo-trades" {
				http.NotFound(w, r)
			} else {
				http.Error(w, "Method not permitted", 405)
			}
			return
		}
		if r.URL.Path == "/health" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
			return
		}
		file := ""
		typ := "text/html; charset=utf-8"
		switch r.URL.Path {
		case "/", "/index.html", "/workbench.html", "/collections", "/mcp-tools", "/files", "/graph", "/questions":
			file = "index.html"
		case "/app.js":
			file = "app.js"
			typ = "text/javascript"
		case "/connection.mjs":
			file = "connection.mjs"
			typ = "text/javascript"
		case "/graph-view.mjs":
			file = "graph-view.mjs"
			typ = "text/javascript"
		case "/style.css":
			file = "style.css"
			typ = "text/css"
		case "/client-config.json":
			file = "client-config.json"
			typ = "application/json"
		default:
			http.NotFound(w, r)
			return
		}
		raw, err := assets.ReadFile("assets/" + file)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		if file == "client-config.json" && os.Getenv("CORE_API_URL") != "" {
			var config map[string]any
			if json.Unmarshal(raw, &config) == nil {
				config["apiBaseUrl"] = os.Getenv("CORE_API_URL")
				raw, _ = json.Marshal(config)
			}
		}
		w.Header().Set("Content-Type", typ)
		w.Write(raw)
	})
}
