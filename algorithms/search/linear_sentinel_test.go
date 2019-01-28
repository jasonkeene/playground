package search_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jasonkeene/playground/algorithms/search"
)

func TestLinearSentinelSearchCorrectness(t *testing.T) {
	data := []int{0, 1, 6, 90, 3, 42, 71, 9, 53, 567}

	t.Run("element exists", func(t *testing.T) {
		result := search.LinearSentinel(data, 90)
		if result != 3 {
			t.Fatal(cmp.Diff(result, 3))
		}
	})
	t.Run("element does not exist", func(t *testing.T) {
		result := search.LinearSentinel(data, 999)
		if result != -1 {
			t.Fatal(cmp.Diff(result, -1))
		}
	})
}

func BenchmarkLinearSentinelBestCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		search.LinearSentinel(benchmarkData, 8191711)
	}
}

func BenchmarkLinearSentinelAverageCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		search.LinearSentinel(benchmarkData, 8348554)
	}
}

func BenchmarkLinearSentinelWorstCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		search.LinearSentinel(benchmarkData, 6309519)
	}
}
