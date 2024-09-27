package handler

import (
	"encoding/json"
	"goserver/routing"
	"log"
	"net/http"
)

func GraphDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Graph data handler called")
	nodes := make([]routing.Node, 0, len(routing.GlbGraph.Nodes))
	for id, node := range routing.GlbGraph.Nodes {
		nodes = append(nodes, node)
		log.Printf("Node: ID=%d, Lat=%f, Lon=%f", id, node.Lat, node.Lon)
	}

	edges := make([]routing.Edge, 0)
	for from, edgeList := range routing.GlbGraph.Edges {
		for _, edge := range edgeList {
			edges = append(edges, edge)
			log.Printf("Edge: From=%d, To=%d, Weight=%f", from, edge.To, edge.Weight)
		}
	}

	data := struct {
		Nodes []routing.Node `json:"nodes"`
		Edges []routing.Edge `json:"edges"`
	}{
		Nodes: nodes,
		Edges: edges,
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Error marshaling graph data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("Graph data to be sent:\n%s", string(jsonData))

	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing graph data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("Graph data sent: %d nodes, %d edges", len(nodes), len(edges))
}
