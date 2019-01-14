package dag_test

import (
	"testing"

	"github.com/jasonkeene/playground/algorithms/dag"
)

func TestTopologicalSort(t *testing.T) {
	graph := testDAG()
	indices := indices(graph)

	ordered := dag.TopologicalSort(graph)

	counts := make([]int, len(graph))
	for i, v := range ordered {
		counts[indices[v.Label]]++
		for j := i + 1; j < len(ordered); j++ {
			assertNoReference(t, graph, indices[v.Label], ordered[j])
		}
	}
	for i, v := range counts {
		if v != 1 {
			t.Fatalf("vertex %d was counted %d times", i, v)
		}
	}
}

func assertNoReference(
	t *testing.T,
	graph []dag.Vertex,
	ref int,
	v dag.Vertex,
) {
	if len(v.Edges) == 0 {
		return
	}
	for _, e := range v.Edges {
		if e.Target == ref {
			t.Fatalf("verted: %#v contains edge to ref: %d", v, ref)
		}
		assertNoReference(t, graph, ref, graph[e.Target])
	}
}

func testDAG() []dag.Vertex {
	return []dag.Vertex{
		{
			Label: "pants",
			Edges: []dag.Edge{
				{Target: 5},
				{Target: 9},
			},
		},
		{
			Label: "chest pad",
			Edges: []dag.Edge{
				{Target: 5},
			},
		},
		{
			Label: "compression shorts",
			Edges: []dag.Edge{
				{Target: 7},
				{Target: 12},
			},
		},
		{
			Label: "catch glove",
			Edges: []dag.Edge{
				{Target: 11},
			},
		},
		{
			Label: "mask",
			Edges: []dag.Edge{
				{Target: 3},
			},
		},
		{
			Label: "sweater",
			Edges: []dag.Edge{
				{Target: 4},
			},
		},
		{
			Label: "t-shirt",
			Edges: []dag.Edge{
				{Target: 1},
			},
		},
		{
			Label: "hose",
			Edges: []dag.Edge{
				{Target: 0},
			},
		},
		{
			Label: "leg pads",
			Edges: []dag.Edge{
				{Target: 3},
			},
		},
		{
			Label: "skates",
			Edges: []dag.Edge{
				{Target: 8},
			},
		},
		{
			Label: "undershorts",
			Edges: []dag.Edge{
				{Target: 2},
			},
		},
		{
			Label: "blocker",
		},
		{
			Label: "cup",
			Edges: []dag.Edge{
				{Target: 0},
			},
		},
		{
			Label: "socks",
			Edges: []dag.Edge{
				{Target: 7},
			},
		},
	}
}

func indices(graph []dag.Vertex) map[string]int {
	indices := make(map[string]int, len(graph))
	for i, v := range graph {
		indices[v.Label] = i
	}
	return indices
}
