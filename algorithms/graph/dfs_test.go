package graph_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/graph"
)

func TestDFS(t *testing.T) {
	g := graph.Graph{
		Nodes: []graph.Node{
			{"y"}, // 0
			{"z"}, // 1
			{"s"}, // 2
			{"t"}, // 3
			{"x"}, // 4
			{"w"}, // 5
			{"v"}, // 6
			{"u"}, // 7
		},
		Edges: [][]graph.Edge{
			{
				{4, 1, 0},
			},
			{
				{0, 1, 0},
				{5, 1, 0},
			},
			{
				{1, 1, 0},
				{5, 1, 0},
			},
			{
				{6, 1, 0},
				{7, 1, 0},
			},
			{
				{1, 1, 0},
			},
			{
				{4, 1, 0},
			},
			{
				{2, 1, 0},
				{5, 1, 0},
			},
			{
				{3, 1, 0},
				{6, 1, 0},
			},
		},
	}

	var order []graph.Node
	graph.DFS(g, 2, func(n graph.Node) {
		order = append(order, n)
	})

	expectedOrder := []graph.Node{
		{"x"},
		{"y"},
		{"w"},
		{"z"},
		{"s"},
		{"v"},
		{"u"},
		{"t"},
	}
	expectedEdges := [][]graph.Edge{
		{
			{4, 1, graph.TreeEdge},
		},
		{
			{0, 1, graph.TreeEdge},
			{5, 1, graph.TreeEdge},
		},
		{
			{1, 1, graph.TreeEdge},
			{5, 1, graph.ForwardEdge},
		},
		{
			{6, 1, graph.TreeEdge},
			{7, 1, graph.TreeEdge},
		},
		{
			{1, 1, graph.BackwardEdge},
		},
		{
			{4, 1, graph.CrossEdge},
		},
		{
			{2, 1, graph.CrossEdge},
			{5, 1, graph.CrossEdge},
		},
		{
			{3, 1, graph.BackwardEdge},
			{6, 1, graph.CrossEdge},
		},
	}

	if !cmp.Equal(g.Edges, expectedEdges) {
		t.Fatal(cmp.Diff(g.Edges, expectedEdges))
	}
	if !cmp.Equal(order, expectedOrder) {
		t.Fatal(cmp.Diff(order, expectedOrder))
	}
}
