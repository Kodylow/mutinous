package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Domain     string
	DBPath     string
	Port       string
	ClnRpcPath string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("env")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Domain:     os.Getenv("DOMAIN"),
		DBPath:     os.Getenv("DB_PATH"),
		Port:       os.Getenv("PORT"),
		ClnRpcPath: os.Getenv("CLN_RPC_PATH"),
	}

	// check that everything is set
	if cfg.Domain == "" || cfg.DBPath == "" || cfg.Port == "" || cfg.ClnRpcPath == "" {
		return nil, errors.New("missing config, make sure you have a .env file with the following variables: DOMAIN, DB_PATH, PORT, CLN_RPC_PATH")
	}

	return cfg, nil
}
