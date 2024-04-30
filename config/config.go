package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	ListenAddr string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, errors.New("couldn't load env")
	}
	var cfg Config
	cfg.ListenAddr = os.Getenv("LISTEN_ADDR")
	cfg.MongoURI = os.Getenv("MONGO_URI")

	return cfg, nil
}
