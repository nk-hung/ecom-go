package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, t any) error {
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(t)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
