package health

import (
	"net/http"

	"github.com/go-chi/chi"
)

func HealthRoutes(r chi.Router) {
	r.Get("/", getHealth)
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
