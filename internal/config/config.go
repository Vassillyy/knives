package config

type Config struct {
	Port           string
	DatabaseURL    string
	MinIOEndpoint  string
	MinIOAccessKey string
	MinIOSecretKey string
}

func Load() *Config {
	return &Config{
		Port:           "8080",
		DatabaseURL:    "postgres://postgres:postgres@localhost:5432/knives?sslmode=disable",
		MinIOEndpoint:  "localhost:9000",
		MinIOAccessKey: "minioadmin",
		MinIOSecretKey: "minioadmin",
	}
}
