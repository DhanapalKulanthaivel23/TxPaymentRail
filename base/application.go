package base

import (
	"github.com/DhanapalKulanthaivel23/addressrail-ai/controller"
	"net/http"
)

func New() http.Handler { return controller.New() }
