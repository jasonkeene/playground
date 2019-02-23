package graph

func relax(u, v int, weight float64, shortest []float64, prev []int) {
	newPath := shortest[u] + weight
	if newPath < shortest[v] {
		shortest[v] = newPath
		prev[v] = u
	}
}
