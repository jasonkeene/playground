package dag_test

import (
	"testing"

	"github.com/jasonkeene/playground/algorithms/dag"
)

func TestShortestPath(t *testing.T) {
	g := testWeightedDAG()

	path := dag.ShortestPath(g, 0, 4)

	expectedPath := []string{"A", "C", "D", "E"}
	if len(path) != len(expectedPath) {
		t.Fatalf("shortest path is not of the right length: %v", path)
	}
	for i, v := range path {
		if v.Label != expectedPath[i] {
			t.Fatalf("shortest path is not correct: %v pos: %d", path, i)
		}
	}
}

func TestShortestPathDijkstras(t *testing.T) {
	g := testWeightedDAG()

	path := dag.ShortestPathDijkstras(g, 0, 4)

	expectedPath := []string{"A", "C", "D", "E"}
	if len(path) != len(expectedPath) {
		t.Fatalf("shortest path is not of the right length: %v", path)
	}
	for i, v := range path {
		if v.Label != expectedPath[i] {
			t.Fatalf("shortest path is not correct: %v pos: %d", path, i)
		}
	}
}

func testWeightedDAG() []dag.Vertex {
	return []dag.Vertex{
		{
			Label: "A",
			Edges: []dag.Edge{
				{Target: 1, Weight: 5},
				{Target: 2, Weight: 3},
			},
		},
		{
			Label: "B",
			Edges: []dag.Edge{
				{Target: 3, Weight: 10},
			},
		},
		{
			Label: "C",
			Edges: []dag.Edge{
				{Target: 3, Weight: 8},
				{Target: 4, Weight: 10},
			},
		},
		{
			Label: "D",
			Edges: []dag.Edge{
				{Target: 4, Weight: 1},
			},
		},
		{
			Label: "E",
		},
	}
}
