package utils

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func ParseJSONBody(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		slog.Error("error parsing the json body", "request", v, "err", err.Error())
		return errors.New("error parsing the request body")
	}
	return nil
}

func SendJSON(status int, w http.ResponseWriter, data any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		slog.Error("error sending the json", "err", err.Error())
		return err
	}
	return nil
}
