package sort_test

import (
	"testing"

	sort "github.com/jasonkeene/playground/algorithms/sort"
)

func TestHeapCorrectness(t *testing.T) {
	testCorrectness(t, sort.Heap)
}
