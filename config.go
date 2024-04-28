package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	MongoURI   string
	ListenAddr string
}

func loadConfig() (config, error) {
	err := godotenv.Load()
	if err != nil {
		return config{}, errors.New("couldn't load env")
	}
	var cfg config
	cfg.ListenAddr = os.Getenv("LISTEN_ADDR")
	cfg.MongoURI = os.Getenv("MONGO_URI")

	return cfg, nil
}
