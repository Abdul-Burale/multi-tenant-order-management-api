package middleware

import (
	"log"
	"net/http"
)

func CorsMiddleware(nxt http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		log.Println("Incoming request: Method -> ", r.Method, " Url -> ", r.URL.Path)
		log.Println("Origin: ", origin)

		if origin == "http://localhost:8080" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			log.Println("CORS allowed for origin", origin)
		} else {
			log.Println("CORS blocked for origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		if r.Method == http.MethodOptions {
			log.Println("Preflight OPTIONS request detected")
			w.WriteHeader(http.StatusOK)
		}

		log.Println("Passing request to next handler")
		nxt(w, r)
	}
}
