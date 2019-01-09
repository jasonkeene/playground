package sort_test

import (
	"fmt"
	"math/rand"
	"reflect"
	gosort "sort"
	"testing"

	sort "github.com/jasonkeene/playground/algorithms/sort"
)

func TestRadixCorrectness(t *testing.T) {
	cases := map[string][]string{
		"empty":              []string{},
		"single item":        []string{"ABC123"},
		"multiple items":     []string{"ABC123", "QAD613", "ABC612"},
		"duplicate items":    []string{"ABC123", "WAB614", "ABC123"},
		"all the same items": []string{"ABC123", "ABC123", "ABC123"},
		"already sorted":     []string{"ABC123", "ABC612", "QAD613"},
		"reversed sorted":    []string{"QAD613", "ABC612", "ABC123"},

		"rand_8":    randomString(8),
		"rand_64":   randomString(64),
		"rand_1024": randomString(1024),
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			expected := make([]string, len(v))
			copy(expected, v)
			gosort.Strings(expected)

			sort.Radix(v, 6)

			if !reflect.DeepEqual(expected, v) {
				t.Errorf(`%q failed to sort:
	Expected: %v
	Actual: %v`, k, expected, v)
			}
		})
	}
}

func randomString(s int) []string {
	a := make([]string, s)
	for i := 0; i < s; i++ {
		a[i] = fmt.Sprintf(
			"%s%s%s%d%d%d",
			string(rand.Intn(26)+1|0x40),
			string(rand.Intn(26)+1|0x40),
			string(rand.Intn(26)+1|0x40),
			rand.Intn(9),
			rand.Intn(9),
			rand.Intn(9),
		)
	}
	return a
}
