package http

import (
	"net/http"
	"www.github.com/Maevlava/Matatani/backend/internal/app"
)

func NewRouter(app *app.Application) http.Handler {
	mux := http.NewServeMux()
	clientMux := clientServerMux(app)

	clientHandler := http.StripPrefix("/client", clientMux)
	mux.Handle("/client/", clientHandler)

	return mux
}

func clientServerMux(app *app.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", MainPageHandler)

	return mux
}
