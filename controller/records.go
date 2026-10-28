package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/DhanapalKulanthaivel23/addressrail-ai/libs"
)

type API struct {
	mu      sync.RWMutex
	records map[int64]libs.Record
	next    int64
}

func New() http.Handler {
	api := &API{records: make(map[int64]libs.Record)}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { write(w, 200, map[string]string{"status": "ok"}) })
	mux.HandleFunc("GET /records", api.list)
	mux.HandleFunc("POST /records", api.create)
	mux.HandleFunc("GET /records/{id}", api.get)
	mux.HandleFunc("PUT /records/{id}", api.update)
	mux.HandleFunc("DELETE /records/{id}", api.remove)
	return mux
}

func write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
func invalid(w http.ResponseWriter) { write(w, 400, map[string]string{"error": "invalid_request"}) }
func missing(w http.ResponseWriter) { write(w, 404, map[string]string{"error": "not_found"}) }
func read(w http.ResponseWriter, r *http.Request) (string, bool) {
	var input struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1024))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		invalid(w)
		return "", false
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		invalid(w)
		return "", false
	}
	name := strings.TrimSpace(input.Name)
	if name == "" || len(name) > 120 {
		invalid(w)
		return "", false
	}
	return name, true
}
func (a *API) list(w http.ResponseWriter, r *http.Request) {
	a.mu.RLock()
	values := make([]libs.Record, 0, len(a.records))
	for _, v := range a.records {
		values = append(values, v)
	}
	a.mu.RUnlock()
	sort.Slice(values, func(i, j int) bool { return values[i].ID < values[j].ID })
	write(w, 200, values)
}
func (a *API) create(w http.ResponseWriter, r *http.Request) {
	name, ok := read(w, r)
	if !ok {
		return
	}
	a.mu.Lock()
	a.next++
	record := libs.Record{ID: a.next, Name: name}
	a.records[record.ID] = record
	a.mu.Unlock()
	w.Header().Set("Location", "/records/"+strconv.FormatInt(record.ID, 10))
	write(w, 201, record)
}
func id(r *http.Request) (int64, bool) {
	value, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return value, err == nil && value > 0
}
func (a *API) get(w http.ResponseWriter, r *http.Request) {
	key, ok := id(r)
	if !ok {
		missing(w)
		return
	}
	a.mu.RLock()
	record, exists := a.records[key]
	a.mu.RUnlock()
	if !exists {
		missing(w)
		return
	}
	write(w, 200, record)
}
func (a *API) update(w http.ResponseWriter, r *http.Request) {
	key, ok := id(r)
	if !ok {
		missing(w)
		return
	}
	name, ok := read(w, r)
	if !ok {
		return
	}
	a.mu.Lock()
	_, exists := a.records[key]
	record := libs.Record{ID: key, Name: name}
	if exists {
		a.records[key] = record
	}
	a.mu.Unlock()
	if !exists {
		missing(w)
		return
	}
	write(w, 200, record)
}
func (a *API) remove(w http.ResponseWriter, r *http.Request) {
	key, ok := id(r)
	if !ok {
		missing(w)
		return
	}
	a.mu.Lock()
	_, exists := a.records[key]
	if exists {
		delete(a.records, key)
	}
	a.mu.Unlock()
	if !exists {
		missing(w)
		return
	}
	w.WriteHeader(204)
}
