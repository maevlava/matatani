package config

import "os"

type APIConfig struct {
	JWTSecret string
}

func Load() *APIConfig {
	JWTSecret := os.Getenv("JWT_SECRET")
	
	return &APIConfig{
		JWTSecret: JWTSecret,
	}
}
