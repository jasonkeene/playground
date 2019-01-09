package sort_test

import (
	"reflect"
	gosort "sort"
	"testing"

	sort "github.com/jasonkeene/playground/algorithms/sort"
)

func TestCountingCorrectness(t *testing.T) {
	testCorrectnessRange(t, 5, sort.Counting(5))
}

func TestReallySimpleSortCorrectness(t *testing.T) {
	testCorrectnessRange(t, 2, sort.ReallySimpleSort)
}

func testCorrectnessRange(t *testing.T, max int, f func([]int)) {
	cases := map[string][]int{
		"empty":       []int{},
		"single item": []int{max - 1},

		"rand_8":    random(max, 8),
		"rand_64":   random(max, 64),
		"rand_1024": random(max, 1024),
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			expected := make([]int, len(v))
			copy(expected, v)
			gosort.Ints(expected)

			f(v)

			if !reflect.DeepEqual(expected, v) {
				t.Errorf(`%q failed to sort:
	Expected: %v
	Actual: %v`, k, expected, v)
			}
		})
	}
}
