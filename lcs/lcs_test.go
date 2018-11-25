package lcs_test

import (
	"fmt"
	"testing"

	"github.com/jasonkeene/playground/lcs"
)

var fixtures = map[string]struct {
	A      string
	B      string
	Result string
	DiffA  string
	DiffB  string
}{
	"sub": {
		A:      "I am waiting.",
		B:      "I am always waiting.",
		Result: "I am waiting.",
		DiffA:  "I am waiting.",
		DiffB:  "I am +<al>wa+<ys wa>iting.",
	},
	"sub_small": {
		A:      ".",
		B:      "I.",
		Result: ".",
		DiffA:  ".",
		DiffB:  "+<I>.",
	},
	"super": {
		A:      "I am waiting.",
		B:      "I wait.",
		Result: "I wait.",
		DiffA:  "I -<am >wait-<ing>.",
		DiffB:  "I wait.",
	},
	"super_small": {
		A:      "I.",
		B:      ".",
		Result: ".",
		DiffA:  "-<I>.",
		DiffB:  ".",
	},
	"diff": {
		A:      "I am waiting.",
		B:      "I've been waiting.",
		Result: "I  waiting.",
		DiffA:  "I -<am> waiting.",
		DiffB:  "I+<'ve> +<been> waiting.",
	},
}

func test(t *testing.T, tf func(string, string) string) {
	for k, f := range fixtures {
		result := tf(f.A, f.B)
		if result != f.Result {
			t.Fatalf("Bad result for test %s: %q != %q", k, result, f.Result)
		}
	}
}

func TestRecursive(t *testing.T) {
	test(t, lcs.Recursive)
}

func TestRecursiveMemoized(t *testing.T) {
	test(t, func(a, b string) string {
		return lcs.RecursiveMemoized(a, b, map[lcs.MemKey]string{})
	})
}

func TestRecursiveIndexes(t *testing.T) {
	test(t, func(a, b string) string {
		return lcs.RecursiveIndexes(a, b, 0, 0)
	})
}

func TestRecursiveIndexesMemoized(t *testing.T) {
	test(t, func(a, b string) string {
		return lcs.RecursiveIndexesMemoized(a, b, 0, 0, map[lcs.MemKeyIndexes]string{})
	})
}

const (
	benchA = "I am waiting."
	benchB = "I've been waiting."
)

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lcs.Recursive(benchA, benchB)
	}
}

func BenchmarkRecursiveMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lcs.RecursiveMemoized(benchA, benchB, map[lcs.MemKey]string{})
	}
}

func BenchmarkRecursiveIndexes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lcs.RecursiveIndexes(benchA, benchB, 0, 0)
	}
}

func BenchmarkRecursiveIndexesMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lcs.RecursiveIndexesMemoized(benchA, benchB, 0, 0, map[lcs.MemKeyIndexes]string{})
	}
}

func TestDiff(t *testing.T) {
	for _, f := range fixtures {
		result := lcs.Diff(f.A, f.B)
		fmt.Printf("%#v\n", result)
		if result.A != f.DiffA {
			t.Fatalf("Bad result: %q != %q", result.A, f.DiffA)
		}
		if result.B != f.DiffB {
			t.Fatalf("Bad result: %q != %q", result.B, f.DiffB)
		}
	}
}
