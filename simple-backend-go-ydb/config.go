package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DSN        string
	PathToKey  string
	TableName  string
}

func LoadConfig() (Config, error) {
	cfg := Config{
		Port:      getenvDefault("PORT", "8080"),
		DSN:       os.Getenv("YDB_DSN"),
		PathToKey: os.Getenv("YDB_KEY_FILE"),
		TableName: getenvDefault("YDB_TABLE", "favorites"),
	}

	if cfg.DSN == "" {
		return Config{}, fmt.Errorf("YDB_DSN is required")
	}

	if cfg.PathToKey == "" {
		return Config{}, fmt.Errorf("YDB_KEY_FILE is required")
	}

	return cfg, nil
}

func getenvDefault(name, fallback string) string {
	v := os.Getenv(name)
	if v == "" {
		return fallback
	}
	return v
}
