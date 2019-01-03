package sort_test

import (
	"testing"

	sort "github.com/jasonkeene/playground/algorithms/sort"
)

func TestQuickCorrectness(t *testing.T) {
	testCorrectness(t, sort.Quick)
}
