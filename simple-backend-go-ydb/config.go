package main

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port                string
	DSN                 string
	ServiceAccountKey   string
	UseMetadataAuth     bool
	TableName           string
}

func LoadConfig() (Config, error) {
	keyFromYDB := os.Getenv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS")
	if keyFromYDB == "" {
		// Backward compatibility with previous workshop env name.
		keyFromYDB = os.Getenv("YDB_KEY_FILE")
	}

	cfg := Config{
		Port:              getenvDefault("PORT", "8080"),
		DSN:               os.Getenv("YDB_DSN"),
		ServiceAccountKey: keyFromYDB,
		UseMetadataAuth:   isTruthy(os.Getenv("YDB_METADATA_CREDENTIALS")),
		TableName:         getenvDefault("YDB_TABLE", "favorites"),
	}

	if cfg.DSN == "" {
		return Config{}, fmt.Errorf("YDB_DSN is required")
	}

	if cfg.ServiceAccountKey == "" && !cfg.UseMetadataAuth {
		return Config{}, fmt.Errorf("set either YDB_METADATA_CREDENTIALS=1 or YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS")
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

func isTruthy(v string) bool {
	v = strings.TrimSpace(strings.ToLower(v))
	return v == "1" || v == "true" || v == "yes"
}
