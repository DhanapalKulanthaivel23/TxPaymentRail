package controller

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCRUD(t *testing.T) {
	api := New()
	for _, step := range []struct {
		method, path, body string
		status             int
	}{
		{"GET", "/health", "", 200},
		{"POST", "/records", `{"name":"sample"}`, 201},
		{"GET", "/records/1", "", 200},
		{"PUT", "/records/1", `{"name":"updated"}`, 200},
		{"GET", "/records", "", 200},
		{"DELETE", "/records/1", "", 204},
		{"GET", "/records/1", "", 404},
	} {
		req := httptest.NewRequest(step.method, step.path, strings.NewReader(step.body))
		w := httptest.NewRecorder()
		api.ServeHTTP(w, req)
		if w.Code != step.status {
			t.Fatalf("%s %s: got %d want %d", step.method, step.path, w.Code, step.status)
		}
		if w.Code != 204 {
			var result any
			if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
				t.Fatal(err)
			}
		}
	}
}
func TestInvalidInput(t *testing.T) {
	for _, body := range []string{`{}`, `{"name":" "}`, `{"name":"sample","extra":true}`, `{"name":"sample"} {}`, strings.Repeat("x", 2000)} {
		w := httptest.NewRecorder()
		New().ServeHTTP(w, httptest.NewRequest("POST", "/records", strings.NewReader(body)))
		if w.Code != 400 {
			t.Fatalf("invalid input returned %d", w.Code)
		}
	}
}
