package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"www.github.com/Maevlava/Matatani/backend/internal/app"
	"www.github.com/Maevlava/Matatani/backend/internal/config"
	httpdelivery "www.github.com/Maevlava/Matatani/backend/internal/delivery/http"
)

const developmentPort = "3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		_ = fmt.Errorf("error loading .env file")
	}

	cfg := config.Load()
	appInstance := app.NewApplication(cfg)

	router := httpdelivery.NewRouter(appInstance)
	server := http.Server{
		Addr:    ":" + developmentPort,
		Handler: router,
	}

	fmt.Println("Server listening on port ", developmentPort)
	log.Fatal(server.ListenAndServe())
}
