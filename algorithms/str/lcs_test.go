package str_test

import (
	"testing"

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

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := str.LCS(tc.a, tc.b)
			if _, ok := tc.lcs[result]; !ok {
				t.Fatalf("invalid result: %q", result)
			}
		})
	}
}
