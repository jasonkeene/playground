package graph

import "sort"

func DetectNegativeWeightCycle(
	s int,
	g Graph,
	shortest []float64,
	prev []int,
) []int {
	var (
		present bool
		v       int
	)
	for u, es := range g.Edges {
		for _, e := range es {
			if shortest[u]+e.Weight < shortest[e.Target] {
				present = true
				v = e.Target
				break
			}
		}
	}
	if !present {
		return nil
	}

	visited := make([]bool, len(shortest))
	for !visited[v] {
		visited[v] = true
		v = prev[v]
	}

	visited = make([]bool, len(shortest))
	invalid := make([]int, 0)
	q := []int{v}
	for len(q) > 0 {
		var x int
		x, q = q[0], q[1:len(q)]
		if visited[x] {
			continue
		}
		visited[x] = true
		invalid = append(invalid, x)

		for _, e := range g.Edges[x] {
			q = append(q, e.Target)
		}
	}

	sort.Ints(invalid)
	return invalid
}
