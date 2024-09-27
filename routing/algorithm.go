package routing

import "math"

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3 // Earth's radius in meters
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c // in meters
	return distance
}

func Dijkstra(start, end int64) ([]int64, float64) {
	distances := make(map[int64]float64)
	previous := make(map[int64]int64)
	pq := make(PriorityQueue, 0)

	for id := range GlbGraph.Nodes {
		if id == start {
			distances[id] = 0
			pq.push(&Item{value: id, priority: 0})
		} else {
			distances[id] = math.Inf(1)
			pq.push(&Item{value: id, priority: math.Inf(1)})
		}
	}

	for pq.Len() > 0 {
		current := pq.pop().value

		if current == end {
			break
		}

		for _, edge := range GlbGraph.Edges[current] {
			newDist := distances[current] + edge.Weight
			if newDist < distances[edge.To] {
				distances[edge.To] = newDist
				previous[edge.To] = current
				pq.update(edge.To, newDist)
			}
		}
	}

	path := []int64{}
	current := end
	for current != start {
		prev, exists := previous[current]
		if !exists {
			// No path found
			return []int64{}, math.Inf(1)
		}
		path = append([]int64{current}, path...)
		current = prev
	}
	path = append([]int64{start}, path...)

	return path, distances[end]
}
