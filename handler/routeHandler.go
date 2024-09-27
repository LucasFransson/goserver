package handler

import (
	"encoding/json"
	"goserver/routing"
	"log"
	"net/http"
	"strconv"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Route handler called")
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		log.Printf("Invalid start node: %v", err)
		http.Error(w, "Invalid start node", http.StatusBadRequest)
		return
	}

	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		log.Printf("Invalid end node: %v", err)
		http.Error(w, "Invalid end node", http.StatusBadRequest)
		return
	}

	path, distance := routing.Dijkstra(start, end)

	result := struct {
		Path     []int64 `json:"path"`
		Distance float64 `json:"distance"`
	}{
		Path:     path,
		Distance: distance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	log.Printf("Route found: %+v", result)
}
