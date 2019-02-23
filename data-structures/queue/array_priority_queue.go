package queue

import "log"

type Array struct {
	values []Element
}

func NewArray(qes ...Element) *Array {
	return &Array{
		values: qes,
	}
}

func (a *Array) Empty() bool {
	return len(a.values) == 0
}

func (a *Array) Insert(qe Element) {
	for i, v := range a.values {
		if v.Key >= qe.Key {
			a.values = append(a.values, Element{})
			copy(a.values[i+1:], a.values[i:])
			a.values[i] = qe
			return
		}
	}
	a.values = append(a.values, qe)
}

func (a *Array) PopMin() Element {
	min := a.values[0]
	a.values = a.values[1:]
	return min
}

func (a *Array) Decrease(qe Element) {
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
