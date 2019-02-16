package str_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/str"
)

func TestTransform(t *testing.T) {
	testCases := map[string]struct {
		a   string
		b   string
		ops []str.Operation
	}{
		"copies": {
			a: "abc",
			b: "abc",

			ops: []str.Operation{
				{
					Type: str.Copy,
					A:    'a',
				},
				{
					Type: str.Copy,
					A:    'b',
				},
				{
					Type: str.Copy,
					A:    'c',
				},
			},
		},
		"replaces": {
			a: "abc",
			b: "123",

			ops: []str.Operation{
				{
					Type: str.Replace,
					A:    'a',
					B:    '1',
				},
				{
					Type: str.Replace,
					A:    'b',
					B:    '2',
				},
				{
					Type: str.Replace,
					A:    'c',
					B:    '3',
				},
			},
		},
		"deletes": {
			a: "abc",
			b: "",

			ops: []str.Operation{
				{
					Type: str.Delete,
					A:    'a',
				},
				{
					Type: str.Delete,
					A:    'b',
				},
				{
					Type: str.Delete,
					A:    'c',
				},
			},
		},
		"inserts": {
			a: "",
			b: "123",

			ops: []str.Operation{
				{
					Type: str.Insert,
					A:    '1',
				},
				{
					Type: str.Insert,
					A:    '2',
				},
				{
					Type: str.Insert,
					A:    '3',
				},
			},
		},
		"dna": {
			a: "ACAAGC",
			b: "CCGT",

			ops: []str.Operation{
				{
					Type: str.Delete,
					A:    'A',
				},
				{
					Type: str.Copy,
					A:    'C',
				},
				{
					Type: str.Delete,
					A:    'A',
				},
				{
					Type: str.Replace,
					A:    'A',
					B:    'C',
				},
				{
					Type: str.Copy,
					A:    'G',
				},
				{
					Type: str.Replace,
					A:    'C',
					B:    'T',
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := str.Transform(tc.a, tc.b)
			if !cmp.Equal(result, tc.ops) {
				t.Fatalf(cmp.Diff(result, tc.ops))
			}
		})
	}
}
