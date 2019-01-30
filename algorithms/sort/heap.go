package sort

import "github.com/jasonkeene/playground/algorithms/graph"

func Heap(a []int) {
	q := graph.NewHeapQueue()
	for _, v := range a {
		q.Insert(graph.QueueElement{Key: float64(v), Value: v})
	}
	for i := 0; !q.Empty(); i++ {
		a[i] = q.PopMin().Value.(int)
	}
}
