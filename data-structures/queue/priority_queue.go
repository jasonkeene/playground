package queue

type Element struct {
	Key   float64
	Value interface{}
}

type Priority interface {
	Empty() bool
	Insert(Element)
	PopMin() Element
	Decrease(Element)
}
