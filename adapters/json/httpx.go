package json

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{
		"error": msg,
	})
}

func DecodeJSON[T any](r *http.Request) (*T, error) {
	defer r.Body.Close()

	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

func DecodeJSONOrFail[T any](w http.ResponseWriter, r *http.Request) (*T, bool) {
	body, err := DecodeJSON[T](r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid request body")
		return nil, false
	}
	return body, false

}
