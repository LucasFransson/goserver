package handler

import (
	"encoding/json"
	"goserver/routing"
	"log"
	"net/http"
	"strconv"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Generate handler called")
	sizeStr := r.URL.Query().Get("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		log.Printf("Invalid size parameter: %v", err)
		http.Error(w, "Invalid size parameter", http.StatusBadRequest)
		return
	}

	routing.GenerateRandomGraph(size)

	response := struct {
		Message string `json:"message"`
		Size    int    `json:"size"`
	}{
		Message: "Random graph generated successfully",
		Size:    size,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("Graph generated with size: %d", size)
}
