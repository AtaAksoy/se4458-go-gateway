package main

import (
	"log"
	"net/http"
	"os"
	"se4458-go-gateway/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	authAPI := os.Getenv("AUTH_API_URL")
	jobPostingAPI := os.Getenv("JOB_POSTING_API_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	r.Handle("/auth/*", handlers.ReverseProxy(authAPI, "/auth"))
	r.Handle("/jobs/*", handlers.ReverseProxy(jobPostingAPI, "/jobs"))
	r.Handle("/job-search/*", handlers.ReverseProxy(jobPostingAPI, "/jobs/search"))

	log.Printf("🚀 API Gateway running on :%s\n", port)
	http.ListenAndServe(":"+port, corsMiddleware(r))
}
