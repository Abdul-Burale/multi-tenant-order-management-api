package handlers

import (
	"encoding/json"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Users EndPoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
