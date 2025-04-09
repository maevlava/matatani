package app

import "www.github.com/Maevlava/Matatani/backend/internal/config"

type Application struct {
	Config *config.APIConfig
}

func NewApplication(config *config.APIConfig) *Application {
	return &Application{
		Config: config,
	}
}
