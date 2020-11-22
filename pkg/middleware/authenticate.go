package middleware

import (
	"net/http"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/constants"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie(constants.CookieName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	})
}
