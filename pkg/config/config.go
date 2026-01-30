package config

import "os"

type Config struct {
	HTTPPort string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

func Load() Config {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		HTTPPort: port,
		DBHost:   getenv("DB_HOST", "localhost"),
		DBPort:   getenv("DB_PORT", "5432"),
		DBUser:   getenv("DB_USER", "platform"),
		DBPass:   getenv("DB_PASS", "platform"),
		DBName:   getenv("DB_NAME", "platform"),
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
