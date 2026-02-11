package handlers

import (
	"encoding/json"
	"net/http"
)

func TentantHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Tenants EndPoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
