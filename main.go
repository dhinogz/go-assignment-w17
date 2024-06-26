package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/dhinogz/go-assignment-w17/db"
	"github.com/dhinogz/go-assignment-w17/handlers"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		slog.Error("could not load env", "err", err)
		os.Exit(1)
	}

	store, err := db.NewMongoStore(cfg.mongoURI)
	if err != nil {
		slog.Error("could not initialize mongo store", "err", err)
		os.Exit(1)
	}

	inMemoryStorage := make(map[string]string)

	handlers := handlers.New(store, inMemoryStorage)

	mux := NewMux(handlers)

	srv := http.Server{
		Addr:    cfg.listenAddr,
		Handler: mux,
	}

	slog.Info("Starting server", "listenAddr", cfg.listenAddr)
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Error in server, will end process", "err", err)
		os.Exit(1)
	}
}
