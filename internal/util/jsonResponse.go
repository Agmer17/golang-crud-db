package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Println(err)
		return
	}
}
