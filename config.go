package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	mongoURI   string
	listenAddr string
}

func LoadConfig() (config, error) {
	err := godotenv.Load()
	if err != nil {
		return config{}, errors.New("couldn't load env")
	}
	var cfg config
	cfg.listenAddr = os.Getenv("LISTEN_ADDR")
	cfg.mongoURI = os.Getenv("MONGO_URI")

	return cfg, nil
}
