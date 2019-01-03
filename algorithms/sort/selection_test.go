package sort_test

import (
	"testing"

	sort "github.com/jasonkeene/playground/algorithms/sort"
)

func TestSelectionCorrectness(t *testing.T) {
	testCorrectness(t, sort.Selection)
}
