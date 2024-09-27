package main

import (
	"goserver/handler"
	"goserver/routing"
	"log"
	"net/http"

	"golang.org/x/exp/rand"
)

func main() {
	rand.Seed(42) // Set a seed for reproducibility

	// Initialize the graph with a small number of nodes
	routing.GenerateRandomGraph(10)
	log.Println("Initial graph generated with 10 nodes")

	http.HandleFunc("/", handler.MapHandler)
	http.HandleFunc("/route", handler.RouteHandler)
	http.HandleFunc("/generate", handler.GenerateHandler)
	http.HandleFunc("/graph-data", handler.GraphDataHandler)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
