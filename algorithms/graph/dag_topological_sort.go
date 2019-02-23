package graph

func DAGTopologicalSort(g Graph) []int {
	next := make([]int, 0, len(g.Nodes)/4)

	// compute in-degrees
	inDegree := make([]int, len(g.Nodes))
	for _, edges := range g.Edges {
		for _, e := range edges {
			inDegree[e.Target]++
		}
	}

	// start with nodes that have in-degree of 0
	for i, degree := range inDegree {
		if degree == 0 {
			next = append(next, i)
		}
	}

	sorted := make([]int, 0, len(g.Nodes))
	for len(next) != 0 {
		n := next[0]
		next = next[1:]

		sorted = append(sorted, n)
		for _, e := range g.Edges[n] {
			inDegree[e.Target]--
			if inDegree[e.Target] == 0 {
				next = append(next, e.Target)
			}
		}
	}

	return sorted
}
