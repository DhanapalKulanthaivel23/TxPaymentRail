package main

import (
	"github.com/DhanapalKulanthaivel23/TxPaymentRail/presentation"
	"log"
	"net/http"
	"os"
	"time"
)

func app() http.Handler { return presentation.New() }
func main() {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	server := &http.Server{Addr: addr, Handler: app(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 30 * time.Second, IdleTimeout: 60 * time.Second}
	log.Fatal(server.ListenAndServe())
}
