package middleware

import (
	"encoding/json"
	"net/http"
)

//JSON tries to decode the body if Content-Type is application/json
//A 404 will be returned for failure to do so
func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType == "application/json" {
			var body interface{}
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("JSON Error"))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
