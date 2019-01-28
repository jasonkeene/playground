package search_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/search"
)

func TestLinearSearchCorrectness(t *testing.T) {
	data := []int{0, 1, 6, 90, 3, 42, 71, 9, 53, 567}

	t.Run("element exists", func(t *testing.T) {
		result := search.Linear(data, 90)
		if result != 3 {
			t.Fatal(cmp.Diff(result, 3))
		}
	})
	t.Run("element does not exist", func(t *testing.T) {
		result := search.Linear(data, 999)
		if result != -1 {
			t.Fatal(cmp.Diff(result, -1))
		}
	})
}

func BenchmarkLinearBestCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		search.Linear(benchmarkData, 8191711)
	}
}

func BenchmarkLinearAverageCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		search.Linear(benchmarkData, 8348554)
	}
}

func BenchmarkLinearWorstCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		search.Linear(benchmarkData, 6309519)
	}
}
