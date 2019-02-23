package graph_test

import (
	"testing"

	"github.com/jasonkeene/playground/algorithms/graph"
)

func TestTopologicalSort(t *testing.T) {
	g := testDAG()

	ordered := graph.DAGTopologicalSort(g)

	counts := make([]int, len(g.Nodes))
	for i, n1 := range ordered {
		counts[n1]++
		for j := i + 1; j < len(ordered); j++ {
			n2 := ordered[j]
			assertNoReference(t, g, n1, n2)
		}
	}
	for i, v := range counts {
		if v != 1 {
			t.Fatalf("vertex %d was counted %d times", i, v)
		}
	}
}

func assertNoReference(t *testing.T, g graph.Graph, n1 int, n2 int) {
	for _, e := range g.Edges[n2] {
		if e.Target == n1 {
			t.Fatalf("invalid ordering: %d contains edge to: %d", n2, n1)
		}
		assertNoReference(t, g, n1, e.Target)
	}
}

// DAG (direction of edges is downward:
//
//   6  13  10
//   |   |   |
//   1   |   2
//   |   | / |
//   |   7  12
//   |   |
//   |   0
//   | / |
//   5   9
//   |   |
//   4   8
//   | /
//   3
//   |
//  11
//
func testDAG() graph.Graph {
	return graph.Graph{
		Nodes: []graph.Node{
			{"pants"},              // 0
			{"chest pad"},          // 1
			{"compression shorts"}, // 2
			{"catch glove"},        // 3
			{"mask"},               // 4
			{"sweater"},            // 5
			{"t-shirt"},            // 6
			{"hose"},               // 7
			{"leg pads"},           // 8
			{"skates"},             // 9
			{"undershorts"},        // 10
			{"blocker"},            // 11
			{"cup"},                // 12
			{"socks"},              // 13
		},
		Edges: [][]graph.Edge{
			{ // 0
				{5, 1, 0}, // sweater
				{9, 1, 0}, // skates
			},
			{ // 1
				{5, 1, 0}, // sweater
			},
			{ // 2
				{7, 1, 0},  // hose
				{12, 1, 0}, // cup
			},
			{ // 3
				{11, 1, 0}, // blocker
			},
			{ // 4
				{3, 1, 0}, // catch glove
			},
			{ // 5
				{4, 1, 0}, // mask
			},
			{ // 6
				{1, 1, 0}, // chest pad
			},
			{ // 7
				{0, 1, 0}, // pants
			},
			{ // 8
				{3, 1, 0}, // catch glove
			},
			{ // 9
				{8, 1, 0}, // leg pads
			},
			{ // 10
				{2, 1, 0}, // compression shorts
			},
			{}, // 11
			{ // 12
				{0, 1, 0}, // pants
			},
			{ // 13
				{7, 1, 0}, //hose
			},
		},
	}
}
