package handlers

import (
    "encoding/json"
    "log"
    "net/http"
)

type ErrorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(v); err != nil {
        log.Printf("Error encoding JSON response: %v", err)
        http.Error(w, `{"code":"encode_failed","message":"Error encoding response"}`, http.StatusInternalServerError)
    }
}
