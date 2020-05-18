package handler

import (
	"net/http"
	"encoding/json"
)

// ResponseData ...
type ResponseData struct {
	Hello string `json:"hello"`
}

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseData{Hello: "world"})
}
