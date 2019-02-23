package graph_test

import (
	"testing"

	"github.com/jasonkeene/playground/algorithms/graph"
)

func TestDAGShortestPathTopologicalSort(t *testing.T) {
	testCases := map[string]func(graph.Graph, int, int) []int{
		"topological_sort": graph.DAGShortestPathTopologicalSort,
		"dijkstra_sort":    graph.DAGShortestPathDijkstra,
	}

	for name, f := range testCases {
		t.Run(name, func(t *testing.T) {
			g := testWeightedDAG()

			path := f(g, 0, 4)

			expectedPath := []string{"A", "C", "D", "E"}
			if len(path) != len(expectedPath) {
				t.Fatalf("shortest path is not of the right length: %v", path)
			}

			for i, n := range path {
				if g.Nodes[n].Label != expectedPath[i] {
					t.Fatalf("shortest path is not correct: %v pos: %d", path, i)
				}
			}
		})
	}
}

// DAG (direction of edges is downward:
//
//   A
//   | \
//   5  3
//   |    \
//   B     C
//   |   / |
//  10  8  |
//   |/    |
//   D     |
//   |    /
//   1  10
//   | /
//   E
//
func testWeightedDAG() graph.Graph {
	return graph.Graph{
		Nodes: []graph.Node{
			{"A"}, // 0
			{"B"}, // 1
			{"C"}, // 2
			{"D"}, // 3
			{"E"}, // 4
		},
		Edges: [][]graph.Edge{
			{ // 0
				{1, 5, 0},
				{2, 3, 0},
			},
			{ // 1
				{3, 10, 0},
			},
			{ // 2
				{3, 8, 0},
				{4, 10, 0},
			},
			{ // 3
				{4, 1, 0},
			},
			{}, // 4
		},
	}
}
