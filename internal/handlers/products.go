package handlers

import (
	"encoding/json"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Products EndPoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
