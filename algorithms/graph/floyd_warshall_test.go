package graph_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/graph"
)

func TestFloydWarshall(t *testing.T) {
	g := graph.Graph{
		Nodes: []graph.Node{
			{"A"},
			{"B"},
			{"C"},
			{"D"},
		},
		Edges: [][]graph.Edge{
			{
				{1, 3},
				{2, 8},
			},
			{
				{3, 1},
			},
			{
				{1, 4},
			},
			{
				{0, 2},
				{2, -5},
			},
		},
	}

	shortest, prev := graph.FloydWarshall(g)
	expectedShortest := [][]float64{
		{0, 3, -1, 4},
		{3, 0, -4, 1},
		{7, 4, 0, 5},
		{2, -1, -5, 0},
	}
	expectedPrev := [][]int{
		{-1, 0, 3, 1},
		{3, -1, 3, 1},
		{3, 2, -1, 1},
		{3, 2, 3, -1},
	}

	if !cmp.Equal(shortest, expectedShortest) {
		t.Fatal(cmp.Diff(shortest, expectedShortest))
	}
	if !cmp.Equal(prev, expectedPrev) {
		t.Fatal(cmp.Diff(prev, expectedPrev))
	}
}
