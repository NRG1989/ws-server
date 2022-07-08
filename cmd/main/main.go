package main

import (
	"log"
	"net/http"

	"ws-server/internal/handlers"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	return mux
}

func main() {
	mux := routes()
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", mux)
}
