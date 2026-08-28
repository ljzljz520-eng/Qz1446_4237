package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func Fail(w http.ResponseWriter, status int, msg string) { JSON(w, status, ErrorResponse{Error: msg}) }
func Method(r *http.Request, allowed string) bool        { return r.Method == allowed }
func Query(r *http.Request, key string) string           { return r.URL.Query().Get(key) }
