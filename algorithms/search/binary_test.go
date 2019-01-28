package search_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jasonkeene/playground/algorithms/search"
)

func TestBinarySearchCorrectness(t *testing.T) {
	tests := map[string]struct {
		data []int
		tcs  map[string]struct {
			search   int
			expected int
		}
	}{
		"ordered slice of integers": {
			[]int{0, 1, 3, 6, 9, 42, 53, 71, 90, 567, 998, 1000, 10001},
			map[string]struct{ search, expected int }{
				"element exists":                      {90, 8},
				"element does not exist":              {999, -1},
				"element is smaller than 0th element": {-20, -1},
				"element is larger than nth element":  {99999, -1},
			},
		},
		"ordered two element slice": {
			[]int{55, 90},
			map[string]struct{ search, expected int }{
				"element exists":         {90, 1},
				"element does not exist": {999, -1},
			},
		},
		"single element slice": {
			[]int{90},
			map[string]struct{ search, expected int }{
				"element exists":         {90, 0},
				"element does not exist": {999, -1},
			},
		},
		"empty slice": {
			[]int{},
			map[string]struct{ search, expected int }{
				"element does not exist": {90, -1},
			},
		},
		"nil": {
			nil,
			map[string]struct{ search, expected int }{
				"element does not exist": {90, -1},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for tcName, tc := range test.tcs {
				t.Run(tcName, func(t *testing.T) {
					result := search.BinarySearch(test.data, tc.search)

					if result != tc.expected {
						t.Fatal(cmp.Diff(result, tc.expected))
					}
				})
			}
		})
	}
}
