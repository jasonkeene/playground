package queue_test

import (
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/data-structures/queue"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestHeapConstructor(t *testing.T) {
	testConstructor(t, func(qes ...queue.Element) queue.Priority {
		return queue.NewHeap(qes...)
	})
}

func TestArrayConstructor(t *testing.T) {
	testConstructor(t, func(qes ...queue.Element) queue.Priority {
		return queue.NewArray(qes...)
	})
}

func testConstructor(t *testing.T, f func(...queue.Element) queue.Priority) {
	h := f([]queue.Element{
		{Key: 0},
		{Key: 1},
		{Key: 2},
		{Key: 3},
		{Key: math.Inf(1)},
	}...)
	h.Insert(queue.Element{Key: 1.5})
	h.Insert(queue.Element{Key: 2.5})
	h.Insert(queue.Element{Key: 9})
	h.Insert(queue.Element{Key: 2.1})

	result := make([]float64, 0, 6)
	for !h.Empty() {
		result = append(result, h.PopMin().Key)
	}

	expected := []float64{0, 1, 1.5, 2, 2.1, 2.5, 3, 9, math.Inf(1)}
	if !cmp.Equal(result, expected) {
		t.Fatal(cmp.Diff(result, expected))
	}
}

func TestHeapReturnsMins(t *testing.T) {
	testReturnsMins(t, queue.NewHeap(), 100000)
}

func TestArrayReturnsMins(t *testing.T) {
	testReturnsMins(t, queue.NewArray(), 10000)
}

func testReturnsMins(t *testing.T, q queue.Priority, count int) {
	nums := make([]float64, count)
	for i := 0; i < count; i++ {
		nums[i] = rand.Float64()
	}

	for _, v := range nums {
		q.Insert(queue.Element{Key: v})
	}

	sort.Float64s(nums)

	result := make([]float64, 0, count)
	for !q.Empty() {
		result = append(result, q.PopMin().Key)
	}

	if !reflect.DeepEqual(result, nums) {
		t.Fatalf("Priority did not return mins in the correct order")
	}
}

func TestHeapDecrease(t *testing.T) {
	testDecrease(t, queue.NewHeap())
}

func TestArrayDecrease(t *testing.T) {
	testDecrease(t, queue.NewArray())
}

func testDecrease(t *testing.T, q queue.Priority) {
	q.Insert(queue.Element{
		Key:   1,
		Value: 1,
	})
	q.Insert(queue.Element{
		Key:   3,
		Value: 3,
	})
	q.Insert(queue.Element{
		Key:   6,
		Value: 6,
	})
	q.Insert(queue.Element{
		Key:   7,
		Value: 7,
	})
	q.Insert(queue.Element{
		Key:   10,
		Value: 10,
	})
	q.Insert(queue.Element{
		Key:   11,
		Value: 11,
	})

	const v = "some uniqe value"
	q.Insert(queue.Element{
		Key:   10,
		Value: v,
	})
	q.Decrease(queue.Element{
		Key:   7,
		Value: v,
	})
	q.Decrease(queue.Element{
		Key:   4,
		Value: v,
	})

	result := make([]queue.Element, 0, 7)
	for !q.Empty() {
		result = append(result, q.PopMin())
	}

	if result[2].Key != 4 || result[2].Value.(string) != v {
		t.Fatalf("Priority did not return correct element: %v", result[2])
	}
}
