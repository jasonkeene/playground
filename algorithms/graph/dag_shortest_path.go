package graph

import (
	"math"
)

type nextElement struct {
	total float64
	i     int
	path  []int
}

func DAGShortestPathTopologicalSort(g Graph, start, end int) []int {
	shortest := make([]float64, len(g.Nodes))
	for i := range shortest {
		shortest[i] = math.Inf(1)
	}
	shortest[start] = 0
	pred := make([]int, len(g.Nodes))
	for i := range shortest {
		pred[i] = -1
	}

	sorted := DAGTopologicalSort(g)
	for _, n := range sorted {
		for _, e := range g.Edges[n] {
			relax(n, e.Target, e.Weight, shortest, pred)
		}
	}

	path := []int{end}
	next := end
	for {
		if next == start {
			break
		}
		next = pred[next]
		path = append([]int{next}, path...)
	}
	return path
}

func DAGShortestPathDijkstra(g Graph, start, end int) []int {
	next := make([]nextElement, 0, len(g.Nodes)/4)
	next = append(next, nextElement{
		i: start,
	})

	for len(next) > 0 {
		// pop off nextElement
		elem := next[0]
		next = next[1:]

		// see if we are at the end vertex
		if elem.i == end {
			return append(elem.path, elem.i)
		}

		// add edges to the next list
		for _, e := range g.Edges[elem.i] {
			path := make([]int, len(elem.path)+1)
			copy(path, elem.path)
			path[len(path)-1] = elem.i
			ne := nextElement{
				total: e.Weight + elem.total,
				i:     e.Target,
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
