package main

import (
	"log"
	"net/http"
	"os"
	"se4458-go-gateway/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	chatAPI := os.Getenv("CHAT_API_URL")
	gsmAPI := os.Getenv("GSM_API_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	r.Handle("/chat/*", handlers.ReverseProxy(chatAPI, "/chat"))
	r.Handle("/gsm/*", handlers.ReverseProxy(gsmAPI, "/gsm"))

	log.Printf("🚀 API Gateway running on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}
