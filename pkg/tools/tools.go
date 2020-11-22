package tools

import (
	"encoding/json"
	"net/http"
)

func Bind(r *http.Request, m interface{}) error {
	return json.NewDecoder(r.Body).Decode(&m)
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}

func StringReponse(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
