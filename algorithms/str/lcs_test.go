package str_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/str"
)

func TestLCS(t *testing.T) {
	testCases := map[string]struct {
		a   string
		b   string
		lcs map[string]struct{}
	}{
		"single char": {
			a: "abcdef",
			b: "zyqwck",
			lcs: map[string]struct{}{
				"c": struct{}{},
			},
		},
		"adjacent characters": {
			a: "abcdef",
			b: "qzbcay",
			lcs: map[string]struct{}{
				"bc": struct{}{},
			},
		},
		"disjoint characters": {
			a: "abcdef",
			b: "azczez",
			lcs: map[string]struct{}{
				"ace": struct{}{},
			},
		},
		"multiple solutions": {
			a: "abcdef",
			b: "defabc",
			lcs: map[string]struct{}{
				"abc": struct{}{},
				"def": struct{}{},
			},
		},
		"dna": {
			a: "CATCGA",
			b: "GTACCGTCA",

			lcs: map[string]struct{}{
				"TCGA": struct{}{},
				"ACGA": struct{}{},
				"ATCA": struct{}{},
				"CCGA": struct{}{},
				"CTCA": struct{}{},
			},
		},
	}

	impls := map[string](func(string, string) string){
		"LCS":      str.LCS,
		"LCSTable": str.LCSTable,
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			for fname, f := range impls {
				t.Run(fname, func(t *testing.T) {
					result := f(tc.a, tc.b)
					if _, ok := tc.lcs[result]; !ok {
						t.Fatalf("invalid result: %q", result)
					}
				})
			}
		})

		t.Run("all "+name, func(t *testing.T) {
			result := str.AllLCS(tc.a, tc.b)

			if !cmp.Equal(result, tc.lcs) {
				t.Fatal(cmp.Diff(result, tc.lcs))
			}
		})
	}
}

func TestLCSLengthTable(t *testing.T) {
	a := "CATCGA"
	b := "GTACCGTCA"

	result := str.LCSLengthTable(a, b)
	expected := [][]int{
		{0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 2},
		{0, 1, 1, 1, 1, 1, 2, 2, 2},
		{0, 1, 1, 2, 2, 2, 2, 3, 3},
		{1, 1, 1, 2, 2, 3, 3, 3, 3},
		{1, 1, 2, 2, 2, 3, 3, 3, 4},
	}

	if !cmp.Equal(result, expected) {
		t.Fatal(cmp.Diff(result, expected))
	}
}
