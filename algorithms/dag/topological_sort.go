package dag

func TopologicalSort(graph []Vertex) []Vertex {
	next := make([]int, 0, len(graph)/4)

	inDegree := make([]int, len(graph))
	for _, v := range graph {
		for _, e := range v.Edges {
			inDegree[e.Target]++
		}
	}
	for i, v := range inDegree {
		if v == 0 {
			next = append(next, i)
		}
	}

	sorted := make([]Vertex, 0, len(graph))
	for len(next) != 0 {
		v := graph[next[0]]
		next = next[1:]

		sorted = append(sorted, v)
		for _, e := range v.Edges {
			inDegree[e.Target]--
			if inDegree[e.Target] == 0 {
				next = append(next, e.Target)
			}
		}
	}
	return sorted
}
