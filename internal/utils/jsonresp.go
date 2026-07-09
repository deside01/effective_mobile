package utils

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Error any `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, statusCode int, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res := httpError{
		Error: data,
	}

	json.NewEncoder(w).Encode(res)
}
