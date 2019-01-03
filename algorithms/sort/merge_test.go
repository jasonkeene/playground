package sort_test

import (
	"testing"

	sort "github.com/jasonkeene/playground/algorithms/sort"
)

func TestMergeCorrectness(t *testing.T) {
	testCorrectness(t, sort.Merge)
}
