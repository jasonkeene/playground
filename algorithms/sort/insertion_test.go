package sort_test

import (
	"testing"

	sort "github.com/jasonkeene/playground/algorithms/sort"
)

func TestInsertionCorrectness(t *testing.T) {
	testCorrectness(t, sort.Insertion)
}
