package config

import (
    "os"
)

type Config struct {
    MongoURI    string
    MongoDB     string
    Port        string
    FrontendURL string
}

func Load() (*Config, error) {
    return &Config{
        MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
        MongoDB:     getEnv("MONGO_DB", "todo"),
        Port:        getEnv("PORT", "8080"),
        FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
    }, nil
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}