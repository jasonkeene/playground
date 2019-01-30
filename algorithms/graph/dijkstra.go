package graph

import "math"

func Dijkstra(s int, g Graph, q PriorityQueue) ([]float64, []int) {
	for i := range g.Nodes {
		key := math.Inf(1)
		if i == s {
			key = 0
		}
		q.Insert(QueueElement{
			Key:   key,
			Value: i,
		})
	}
	shortest := make([]float64, len(g.Nodes))
	for i := range shortest {
		sh := math.Inf(1)
		if i == s {
			sh = 0
		}
		shortest[i] = sh
	}
	prev := make([]int, len(g.Nodes))
	for i := range prev {
		prev[i] = -1
	}

	for !q.Empty() {
		i := q.PopMin().Value.(int)
		for _, e := range g.Edges[i] {
			tmpShortest := shortest[e.Target]
			relax(i, e.Target, e.Weight, shortest, prev)
			if shortest[e.Target] < tmpShortest {
				q.Decrease(QueueElement{
					Key:   shortest[e.Target],
					Value: e.Target,
				})
			}
		}
	}

	return shortest, prev
}

func relax(u, v int, weight float64, shortest []float64, prev []int) {
	newPath := shortest[u] + weight
	if newPath < shortest[v] {
		shortest[v] = newPath
		prev[v] = u
	}
}
