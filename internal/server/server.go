package server

import (
	"log"
	"net/http"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/config"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/handlers"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/middleware"
)

func New(cfg *config.Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/tenants", middleware.CorsMiddleware(handlers.TentantHandler))
	mux.HandleFunc("/users", middleware.CorsMiddleware(handlers.OrdersHandler))
	mux.HandleFunc("/products", middleware.CorsMiddleware(handlers.ProductsHandler))
	mux.HandleFunc("/orders", middleware.CorsMiddleware(handlers.OrdersHandler))

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	return srv
}

func Run(srv *http.Server) {
	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
