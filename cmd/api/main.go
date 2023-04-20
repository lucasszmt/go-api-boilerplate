package main

import (
	"github.com/go-chi/chi"
	"go-api-boilerplate/handlers"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/health", handlers.Health)
	log.Println("running at :8080")
	http.ListenAndServe(":8080", r)
}
