package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/dhinogz/go-assignment-w17/config"
	"github.com/dhinogz/go-assignment-w17/db"
	"github.com/dhinogz/go-assignment-w17/handlers"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("could not load env", "err", err)
		os.Exit(1)
	}

	store, err := db.NewMongoStore(cfg.MongoURI)
	if err != nil {
		slog.Error("could not initialize mongo store", "err", err)
		os.Exit(1)
	}

	handlers := handlers.NewHandlers(store)

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
