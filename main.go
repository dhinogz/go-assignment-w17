package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		slog.Error("could not load env", "err", err)
		os.Exit(1)
	}

	store, err := NewMongoStore(cfg.MongoURI)
	if err != nil {
		slog.Error("could not initialize mongo store", "err", err)
		os.Exit(1)
	}

	handlers := NewHandlers(store)

	mux := NewMux(handlers)

	srv := http.Server{
		Addr:    cfg.ListenAddr,
		Handler: mux,
	}

	slog.Info("Starting server", "listenAddr", cfg.ListenAddr)
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Error in server, will end process", "err", err)
		os.Exit(1)
	}
}

func NewMux(h Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	// Add some middleware for logging

	mux.HandleFunc("/", h.handleHealth)
	mux.HandleFunc("/records", h.handleRecords)
	mux.HandleFunc("GET /in-memory", h.handleInMemory)
	mux.HandleFunc("POST /in-memory", h.handleCreateInMemory)

	return mux
}
