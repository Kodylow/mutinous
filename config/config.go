package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Domain string
	DBPath string
	Port   string
	ClnRpcPath string
}

func LoadConfig() (*Config, error) {
	return &Config{
		Domain: os.Getenv("DOMAIN"),
		DBPath: os.Getenv("DB_PATH"),
		Port:   os.Getenv("PORT"),
		ClnRpcPath: os.Getenv("CLN_RPC_PATH"),
	}, nil
}
