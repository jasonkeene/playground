package queue

import (
	"log"
	"math"
)

type Heap struct {
	values  []Element
	indices map[interface{}]int
}

func NewHeap(v ...Element) *Heap {
	h := &Heap{
		values:  make([]Element, len(v), len(v)*2),
		indices: make(map[interface{}]int, len(v)),
	}
	prev := math.Inf(-1)
	for i, he := range v {
		if prev > he.Key {
			log.Fatalf("invalid heap init: %v", v)
		}
		prev = he.Key
		h.values[i] = he
		h.indices[he.Value] = i
	}
	return h
}

func (h *Heap) Empty() bool {
	return len(h.values) == 0
}

func (h *Heap) Insert(he Element) {
	h.values = append(h.values, he)
	lastIdx := len(h.values) - 1
	h.indices[he.Value] = lastIdx
	h.up(lastIdx)
}

func (h *Heap) up(idx int) {
	he := h.values[idx]
	for {
		parentIdx := (idx - 1) / 2
		parent := h.values[parentIdx]

		if he.Key >= parent.Key {
			break
		}

		h.values[idx], h.values[parentIdx] = h.values[parentIdx], h.values[idx]
		h.indices[he.Value] = parentIdx
		h.indices[parent.Value] = idx

		idx = parentIdx
	}
}

func (h *Heap) PopMin() Element {
	min := h.values[0]
	delete(h.indices, min.Value)

	lastIdx := len(h.values) - 1
	last := h.values[lastIdx]
	h.values[0] = last
	h.indices[last.Value] = 0
	h.values = h.values[:lastIdx]

	if len(h.values) != 0 {
		h.down(0)
	}

	return min
}

func (h *Heap) down(idx int) {
	he := h.values[idx]
	for {
		leftIdx := idx*2 + 1
		rightIdx := leftIdx + 1

		if leftIdx > len(h.values)-1 {
			break
		}

		childIdx := leftIdx
		if rightIdx <= len(h.values)-1 {
			if h.values[rightIdx].Key < h.values[leftIdx].Key {
				childIdx = rightIdx
			}
		}

		child := h.values[childIdx]
		if he.Key <= child.Key {
			break
		}
		h.values[idx], h.values[childIdx] = h.values[childIdx], h.values[idx]
		h.indices[he.Value] = childIdx
		h.indices[child.Value] = idx

		idx = childIdx
	}
}

func (h *Heap) Decrease(he Element) {
	idx, ok := h.indices[he.Value]
	if !ok {
		log.Fatalf("Value was not in heap: %v", he.Value)
	}

	h.values[idx].Key = he.Key
	h.up(idx)
}
