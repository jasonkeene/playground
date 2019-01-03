package sort_test

import (
	"math"
	"math/rand"
	"testing"
)

type sortingFunc func([]int)

func testCorrectness(t *testing.T, f sortingFunc) {
	cases := map[string][]int{
		"empty":           []int{},
		"single item":     []int{5},
		"multiple items":  []int{5, 8, 16, 2, 71, -13, 61},
		"duplicate items": []int{5, 5, 5, 2, 71, -13, 61},
		"already sorted":  []int{-13, 2, 5, 8, 16, 61, 71},
		"reversed sorted": []int{71, 61, 16, 8, 5, 2, -13},

		"extra_small_rand_8": random(4, 8),
		"small_rand_8":       random(64, 8),
		"med_rand_8":         random(1<<31-1, 8),
		"large_rand_8":       random(math.MaxInt64, 8),

		"extra_small_rand_64": random(4, 64),
		"small_rand_64":       random(64, 64),
		"med_rand_64":         random(1<<31-1, 64),
		"large_rand_64":       random(math.MaxInt64, 64),

		"extra_small_rand_1024": random(4, 1024),
		"small_rand_1024":       random(64, 1024),
		"med_rand_1024":         random(1<<31-1, 1024),
		"large_rand_1024":       random(math.MaxInt64, 1024),
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			f(v)
			problem, ok := sorted(v)
			if !ok {
				t.Errorf(`%q failed to sort:
	Array State: %v
	Problem:
		Index: %d
		Area: %v`, k, v, problem.index, problem.area)
			}
		})
	}
}

type problem struct {
	index int
	area  []int
}

func sorted(a []int) (problem, bool) {
	var last int
	for i, v := range a {
		if i == 0 {
			last = v
			continue
		}

		if v < last {
			return problem{
				index: i,
				area:  area(a, i),
			}, false
		}
	}

	return problem{}, true
}

func area(a []int, i int) []int {
	h := i - 1
	j := i + 1

	if h < 0 {
		h = i
	}
	if j >= len(a) {
		j = i
	}

	return a[h : j+1]
}

func random(n, s int) []int {
	a := make([]int, 0, s)
	for i := 0; i < s; i++ {
		a = append(a, rand.Intn(n))
	}
	return a
}
