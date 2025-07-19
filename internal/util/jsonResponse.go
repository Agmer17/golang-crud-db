package util

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
