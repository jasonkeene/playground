package dag

import (
	"math"
)

type nextElement struct {
	total float64
	v     Vertex
	path  []Vertex
}

func ShortestPath(graph []Vertex, start, end int) []Vertex {
	shortest := make([]float64, len(graph))
	for i := range shortest {
		shortest[i] = math.Inf(1)
	}
	shortest[start] = 0
	pred := make([]*int, len(graph))

	sorted := TopologicalSort(graph)
	for vi, v := range sorted {
		for _, e := range v.Edges {
			relax(shortest, pred, vi, e.Target, e.Weight)
		}
	}

	path := []Vertex{
		graph[end],
	}
	next := end
	for {
		if next == start {
			break
		}
		p := pred[next]
		path = append([]Vertex{graph[*p]}, path...)
		next = *p
	}
	return path
}

func relax(shortest []float64, pred []*int, u, v int, weight float64) {
	if shortest[u]+weight < shortest[v] {
		shortest[v] = shortest[u] + weight
		pred[v] = &u
	}
}

func ShortestPathDijkstras(graph []Vertex, start, end int) []Vertex {
	next := make([]nextElement, 0, len(graph)/4)
	next = append(next, nextElement{
		v: graph[start],
	})

	for len(next) > 0 {
		// pop off v
		elem := next[0]
		next = next[1:]

		// see if we are at the end vertex
		if elem.v.Label == graph[end].Label {
			return append(elem.path, elem.v)
		}

		// add edges to the next list
		for _, e := range elem.v.Edges {
			path := make([]Vertex, len(elem.path)+1)
			copy(path, elem.path)
			path[len(path)-1] = elem.v
			ne := nextElement{
				total: e.Weight + elem.total,
				v:     graph[e.Target],
				path:  path,
			}
			next = insert(next, ne)
		}
	}
	return nil
}

func insert(s []nextElement, ne nextElement) []nextElement {
	for i, elem := range s {
		if ne.total < elem.total {
			s = append(s, nextElement{})
			copy(s[i+1:], s[i:])
			s[i] = ne
			return s
		}
	}
	s = append(s, ne)
	return s
}
