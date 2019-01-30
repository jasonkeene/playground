package graph

type QueueElement struct {
	Key   float64
	Value interface{}
}

type PriorityQueue interface {
	Empty() bool
	Insert(QueueElement)
	PopMin() QueueElement
	Decrease(QueueElement)
}
