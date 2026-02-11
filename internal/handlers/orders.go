package handlers

import (
	"encoding/json"
	"net/http"
)

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Orders EndPoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
