package graph

import "math"

func FloydWarshall(g Graph) ([][]float64, [][]int) {
	// init shortest
	shortest := make([][]float64, len(g.Nodes))
	for u := range shortest {
		shortest[u] = make([]float64, len(g.Nodes))
		for v := range shortest[u] {
			shortest[u][v] = math.Inf(1)
		}
	}

	// init prev
	prev := make([][]int, len(g.Nodes))
	for u := range prev {
		prev[u] = make([]int, len(g.Nodes))
		for v := range prev[u] {
			prev[u][v] = -1
		}
	}

	// populate existing values for shortest and prev based on edges
	for u := range g.Nodes {
		shortest[u][u] = 0
		for _, e := range g.Edges[u] {
			shortest[u][e.Target] = e.Weight
			prev[u][e.Target] = u
		}
	}

	// consider every path of u->v, if it can be improved by visiting x do so
	for x := range g.Nodes {
		for u := range g.Nodes {
			for v := range g.Nodes {
				// see if going from u->v via x is shorter
				shortestViaX := shortest[u][x] + shortest[x][v]
				if shortestViaX < shortest[u][v] {
					// update the shortest val
					shortest[u][v] = shortestViaX
					// prev should now be prev between x->v
					prev[u][v] = prev[x][v]
				}
			}
		}
	}

	return shortest, prev
}
