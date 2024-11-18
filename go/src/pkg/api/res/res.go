package res

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

// WriteJSON записывает Response в формате JSON в http.ResponseWriter
func WriteJSON(w http.ResponseWriter, statusCode int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}

// WriteError записывает Response в формате JSON в http.ResponseWriter
func WriteError(w http.ResponseWriter, statusCode int, err error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := Error(message)
	_ = json.NewEncoder(w).Encode(res)
}

// ReadJSON считывает body из http.Request в структуру
func ReadJSON[T any](r *http.Request) (T, error) {
	var body T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	return body, err
}
