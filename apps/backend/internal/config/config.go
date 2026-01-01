package config

import (
	"log/slog"
	"os"
)

type Config struct {
	Port   string
	Logger *slog.Logger
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
