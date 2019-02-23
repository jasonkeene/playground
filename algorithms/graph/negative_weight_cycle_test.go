package graph_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/graph"
)

func TestNegativeWeightCycles(t *testing.T) {
	t.Run("present", func(t *testing.T) {
		g := graph.Graph{
			Nodes: []graph.Node{
				{"A"},
				{"B"},
				{"C"},
				{"D"},
				{"E"},
				{"F"},
				{"G"},
			},
			Edges: [][]graph.Edge{
				{
					{6, 5, 0},
					{2, 1, 0},
				},
				{},
				{
					{3, 8, 0},
					{1, 5, 0},
				},
				{
					{4, -6, 0},
				},
				{
					{2, -3, 0},
					{5, 2, 0},
				},
				{},
				{},
			},
		}

		shortest, prev := graph.BellmanFord(0, g)
		badVertices := graph.DetectNegativeWeightCycle(0, g, shortest, prev)
		expectedBadVertices := []int{1, 2, 3, 4, 5}

		if shortest[0] != 0 {
			t.Fatalf("shortest[0] != 0: %f", shortest[0])
		}
		if shortest[6] != 5 {
			t.Fatalf("shortest[6] != 5: %f", shortest[6])
		}
		if prev[0] != -1 {
			t.Fatalf("prev[0] != -1: %d", prev[0])
		}
		if prev[6] != 0 {
			t.Fatalf("prev[6] != 0: %d", prev[6])
		}

		if !cmp.Equal(badVertices, expectedBadVertices) {
			t.Fatal(cmp.Diff(badVertices, expectedBadVertices))
		}
	})

	t.Run("not present", func(t *testing.T) {
		g := graph.Graph{
			Nodes: []graph.Node{
				{"A"},
				{"B"},
				{"C"},
			},
			Edges: [][]graph.Edge{
				{
					{1, 1, 0},
				},
				{
					{2, 1, 0},
				},
				{
					{0, 1, 0},
				},
			},
		}

		shortest, prev := graph.BellmanFord(0, g)
		badVertices := graph.DetectNegativeWeightCycle(0, g, shortest, prev)
		var expectedBadVertices []int

		if !cmp.Equal(badVertices, expectedBadVertices) {
			t.Fatal(cmp.Diff(badVertices, expectedBadVertices))
		}
	})
}
