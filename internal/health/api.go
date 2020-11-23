package health

import (
	"net/http"

	"github.com/go-chi/chi"
)

func HealthRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getHealth)
	return r
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
