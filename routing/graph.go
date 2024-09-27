package routing

import "golang.org/x/exp/rand"

// Global variable to store our graph
var GlbGraph Graph

func initGraph() {
	GlbGraph = Graph{
		Nodes: make(map[int64]Node),
		Edges: make(map[int64][]Edge),
	}
}

func addNode(n Node) {
	GlbGraph.Nodes[n.ID] = n
}

func addEdge(e Edge) {
	GlbGraph.Edges[e.From] = append(GlbGraph.Edges[e.From], e)
}

func GenerateRandomGraph(n int) {
	initGraph()

	// Sweden's approximate boundaries
	minLat, maxLat := 55.0, 69.0
	minLon, maxLon := 11.0, 24.0

	// Generate n random nodes
	for i := int64(1); i <= int64(n); i++ {
		addNode(Node{
			ID: i,
			// Lat: rand.Float64()*180 - 90,  // Random latitude between -90 and 90
			// Lon: rand.Float64()*360 - 180, // Random longitude between -180 and 180
			Lat: rand.Float64()*(maxLat-minLat) + minLat, // Random latitude within Sweden
			Lon: rand.Float64()*(maxLon-minLon) + minLon, // Random longitude within Sweden
		})
	}

	// Ensure the graph is connected by creating a path through all nodes
	for i := int64(1); i < int64(n); i++ {
		fromNode := GlbGraph.Nodes[i]
		toNode := GlbGraph.Nodes[i+1]
		weight := haversine(fromNode.Lat, fromNode.Lon, toNode.Lat, toNode.Lon)
		addEdge(Edge{From: i, To: i + 1, Weight: weight})
		addEdge(Edge{From: i + 1, To: i, Weight: weight}) // Undirected graph
	}

	// Add some random additional edges to make the graph more interesting
	extraEdges := n / 2
	for i := 0; i < extraEdges; i++ {
		from := rand.Int63n(int64(n)) + 1
		to := rand.Int63n(int64(n)) + 1
		if from != to {
			fromNode := GlbGraph.Nodes[from]
			toNode := GlbGraph.Nodes[to]
			weight := haversine(fromNode.Lat, fromNode.Lon, toNode.Lat, toNode.Lon)
			addEdge(Edge{From: from, To: to, Weight: weight})
			addEdge(Edge{From: to, To: from, Weight: weight}) // Undirected graph
		}
	}
}
