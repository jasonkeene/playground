package sort

import "github.com/jasonkeene/playground/data-structures/queue"

func Heap(a []int) {
	q := queue.NewHeap()
	for _, v := range a {
		q.Insert(queue.Element{Key: float64(v), Value: v})
	}
	for i := 0; !q.Empty(); i++ {
		a[i] = q.PopMin().Value.(int)
	}
}
