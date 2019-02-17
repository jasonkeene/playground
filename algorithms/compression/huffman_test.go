package compression_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/compression"
)

func TestHuffmanTree(t *testing.T) {
	root := compression.HuffmanTree(
		[]byte{'A', 'C', 'T', 'G'},
		[]int{45, 5, 45, 5},
	)
	expectedRoot := &compression.Node{
		Value: 100,
		Left: &compression.Node{
			Char:  'A',
			Value: 45,
		},
		Right: &compression.Node{
			Value: 55,
			Left: &compression.Node{
				Value: 10,
				Left: &compression.Node{
					Char:  'C',
					Value: 5,
				},
				Right: &compression.Node{
					Char:  'G',
					Value: 5,
				},
			},
			Right: &compression.Node{
				Char:  'T',
				Value: 45,
			},
		},
	}

	if !cmp.Equal(root, expectedRoot) {
		t.Fatal(cmp.Diff(root, expectedRoot))
	}
}
