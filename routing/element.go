package routing

// Node represents a point in the OSM data
type Node struct {
	ID  int64   `json:"id"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Edge represents a connection between two nodes
type Edge struct {
	From   int64   `json:"from"`
	To     int64   `json:"to"`
	Weight float64 `json:"weight"`
}

// Graph represents the entire network of nodes and edges
type Graph struct {
	Nodes map[int64]Node
	Edges map[int64][]Edge
}
