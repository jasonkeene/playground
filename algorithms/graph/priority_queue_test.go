package graph_test

import (
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/graph"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestHeapQueueConstructor(t *testing.T) {
	testQueueConstructor(t, func(qes ...graph.QueueElement) graph.PriorityQueue {
		return graph.NewHeapQueue(qes...)
	})
}

func TestArrayQueueConstructor(t *testing.T) {
	testQueueConstructor(t, func(qes ...graph.QueueElement) graph.PriorityQueue {
		return graph.NewArrayQueue(qes...)
	})
}

func testQueueConstructor(t *testing.T, f func(...graph.QueueElement) graph.PriorityQueue) {
	h := f([]graph.QueueElement{
		{Key: 0},
		{Key: 1},
		{Key: 2},
		{Key: 3},
		{Key: math.Inf(1)},
	}...)
	h.Insert(graph.QueueElement{Key: 1.5})
	h.Insert(graph.QueueElement{Key: 2.5})
	h.Insert(graph.QueueElement{Key: 9})
	h.Insert(graph.QueueElement{Key: 2.1})

	result := make([]float64, 0, 6)
	for !h.Empty() {
		result = append(result, h.PopMin().Key)
	}

	expected := []float64{0, 1, 1.5, 2, 2.1, 2.5, 3, 9, math.Inf(1)}
	if !cmp.Equal(result, expected) {
		t.Fatal(cmp.Diff(result, expected))
	}
}

func TestHeapQueueReturnsMins(t *testing.T) {
	testQueueReturnsMins(t, graph.NewHeapQueue(), 100000)
}

func TestArrayQueueReturnsMins(t *testing.T) {
	testQueueReturnsMins(t, graph.NewArrayQueue(), 10000)
}

func testQueueReturnsMins(t *testing.T, q graph.PriorityQueue, count int) {
	nums := make([]float64, count)
	for i := 0; i < count; i++ {
		nums[i] = rand.Float64()
	}

	for _, v := range nums {
		q.Insert(graph.QueueElement{Key: v})
	}

	sort.Float64s(nums)

	result := make([]float64, 0, count)
	for !q.Empty() {
		result = append(result, q.PopMin().Key)
	}

	if !reflect.DeepEqual(result, nums) {
		t.Fatalf("PriorityQueue did not return mins in the correct order")
	}
}

func TestHeapQueueDecrease(t *testing.T) {
	testQueueDecrease(t, graph.NewHeapQueue())
}

func TestArrayQueueDecrease(t *testing.T) {
	testQueueDecrease(t, graph.NewArrayQueue())
}

func testQueueDecrease(t *testing.T, q graph.PriorityQueue) {
	q.Insert(graph.QueueElement{
		Key:   1,
		Value: 1,
	})
	q.Insert(graph.QueueElement{
		Key:   3,
		Value: 3,
	})
	q.Insert(graph.QueueElement{
		Key:   6,
		Value: 6,
	})
	q.Insert(graph.QueueElement{
		Key:   7,
		Value: 7,
	})
	q.Insert(graph.QueueElement{
		Key:   10,
		Value: 10,
	})
	q.Insert(graph.QueueElement{
		Key:   11,
		Value: 11,
	})

	const v = "some uniqe value"
	q.Insert(graph.QueueElement{
		Key:   10,
		Value: v,
	})
	q.Decrease(graph.QueueElement{
		Key:   7,
		Value: v,
	})
	q.Decrease(graph.QueueElement{
		Key:   4,
		Value: v,
	})

	result := make([]graph.QueueElement, 0, 7)
	for !q.Empty() {
		result = append(result, q.PopMin())
	}

	if result[2].Key != 4 || result[2].Value.(string) != v {
		t.Fatalf("PriorityQueue did not return correct element: %v", result[2])
	}
}
