package graph

import "log"

type ArrayQueue struct {
	values []QueueElement
}

func NewArrayQueue(qes ...QueueElement) *ArrayQueue {
	return &ArrayQueue{
		values: qes,
	}
}

func (a *ArrayQueue) Empty() bool {
	return len(a.values) == 0
}

func (a *ArrayQueue) Insert(qe QueueElement) {
	for i, v := range a.values {
		if v.Key >= qe.Key {
			a.values = append(a.values, QueueElement{})
			copy(a.values[i+1:], a.values[i:])
			a.values[i] = qe
			return
		}
	}
	a.values = append(a.values, qe)
}

func (a *ArrayQueue) PopMin() QueueElement {
	min := a.values[0]
	a.values = a.values[1:]
	return min
}

func (a *ArrayQueue) Decrease(qe QueueElement) {
	for i, v := range a.values {
		if v.Value == qe.Value {
			a.values[i].Key = qe.Key
			j := i - 1
			for {
				if j < 0 || a.values[j].Key <= a.values[i].Key {
					return
				}

				a.values[i], a.values[j] = a.values[j], a.values[i]
				i, j = j, j-1
			}
		}
	}
	log.Fatalf("value was not found for: %v", qe)
}
