package graph

import (
	"math"
)

func BellmanFord(s int, g Graph) ([]float64, []int) {
	shortest := make([]float64, len(g.Nodes))
	for i := range shortest {
		if i != s {
			shortest[i] = math.Inf(1)
		}
	}
	prev := make([]int, len(g.Nodes))
	for i := range prev {
		prev[i] = -1
	}

	for range g.Nodes {
		for i, es := range g.Edges {
			for _, e := range es {
				relax(i, e.Target, e.Weight, shortest, prev)
			}
		}
	}

	return shortest, prev
}
