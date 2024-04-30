package main

import (
	"net/http"

	"github.com/dhinogz/go-assignment-w17/handlers"
)

func NewMux(h handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy"))
	})

	mux.HandleFunc("POST /records", h.HandleRecords)
	mux.HandleFunc("GET /in-memory", h.HandleInMemory)
	mux.HandleFunc("POST /in-memory", h.HandleCreateInMemory)

	return mux
}
