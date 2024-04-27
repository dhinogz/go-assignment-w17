package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	listenAddr := ":4000"
	ctx := context.Background()

	mongoClient, err := mongo.Connect(ctx)
	if err != nil {
		slog.Error("could not initialize new mongo client", "err", err)
		os.Exit(1)
	}

	store, err := NewMongoStore(mongoClient)
	if err != nil {
		slog.Error("could not initialize mongo store", "err", err)
		os.Exit(1)
	}

	handlers := NewHandlers(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.handleHealth)
	mux.HandleFunc("/records", handlers.handleRecords)

	srv := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	slog.Info("Starting server", "listenAddr", listenAddr)
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Error in server, will end process", "err", err)
		os.Exit(1)
	}
}
