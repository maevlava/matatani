package http

import (
	"log"
	"net/http"
)

type MockResponse struct {
	MockUsers []struct {
		Name string `json:"name"`
	} `json:"users"`
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	response := MockResponse{
		MockUsers: []struct {
			Name string `json:"name"`
		}{
			{Name: "Raffael"},
			{Name: "Hizqya"},
			{Name: "Bakhtiar"},
		},
	}

	err := RespondWithJSON(w, http.StatusOK, response)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
}
