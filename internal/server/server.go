package server

import (
	"log"
	"net/http"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/handlers"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/middleware"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tenants", middleware.CorsMiddleware(handlers.TentantHandler))
	mux.HandleFunc("/users", middleware.CorsMiddleware(handlers.OrdersHandler))
	mux.HandleFunc("/products", middleware.CorsMiddleware(handlers.ProductsHandler))
	mux.HandleFunc("/orders", middleware.CorsMiddleware(handlers.OrdersHandler))

	addr := ":3000"
	log.Println("Starting server on", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
