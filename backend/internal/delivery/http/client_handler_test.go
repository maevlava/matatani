package http_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"www.github.com/Maevlava/Matatani/backend/internal/app"
	"www.github.com/Maevlava/Matatani/backend/internal/config"
	httpdelivery "www.github.com/Maevlava/Matatani/backend/internal/delivery/http"
)

func setupRouter() http.Handler {
	conf := config.Load()
	appInstance := app.NewApplication(conf)
	test_mux := httpdelivery.NewRouter(appInstance)

	return test_mux
}

func TestMainPageHandler(t *testing.T) {
	test_mux := setupRouter()

	t.Run("Exact Response", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/client/", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		test_mux.ServeHTTP(w, req)

		expectedResponse := httpdelivery.MockResponse{
			MockUsers: []struct {
				Name string `json:"name"`
			}{
				{Name: "Raffael"},
				{Name: "Hizqya"},
				{Name: "Bakhtiar"},
			},
		}
		var actualResponse httpdelivery.MockResponse
		decoder := json.NewDecoder(w.Body)
		err = decoder.Decode(&actualResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedResponse, actualResponse)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Invalid Method", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/client/", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		test_mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})

}
