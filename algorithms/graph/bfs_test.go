package graph_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/graph"
)

func TestBFS(t *testing.T) {
	g := graph.Graph{
		Nodes: []graph.Node{
			{"r"}, // 0
			{"s"}, // 1
			{"t"}, // 2
			{"u"}, // 3
			{"v"}, // 4
			{"w"}, // 5
			{"x"}, // 6
			{"y"}, // 7
		},
		Edges: [][]graph.Edge{
			{
				{1, 1},
				{4, 1},
			},
			{
				{0, 1},
				{5, 1},
			},
			{
				{5, 1},
				{6, 1},
				{3, 1},
			},
			{
				{2, 1},
				{6, 1},
				{7, 1},
			},
			{
				{0, 1},
			},
			{
				{1, 1},
				{2, 1},
				{6, 1},
			},
			{
				{5, 1},
				{2, 1},
				{3, 1},
				{7, 1},
			},
			{
				{6, 1},
				{3, 1},
			},
		},
	}

	order1 := make([]graph.Node, 0, len(g.Nodes))
	graph.BFS(g, 1, func(n graph.Node) {
		order1 = append(order1, n)
	})
	order2 := make([]graph.Node, 0, len(g.Nodes))
	tree := graph.BFSTree(g, 1, func(n graph.Node) {
		order2 = append(order2, n)
	})

	expectedOrder := []graph.Node{
		{
			Label: "s",
		},
		{
			Label: "r",
		},
		{
			Label: "w",
		},
		{
			Label: "v",
		},
		{
			Label: "t",
		},
		{
			Label: "x",
		},
		{
			Label: "u",
		},
		{
			Label: "y",
		},
	}
	expectedTree := &graph.BFSTreeNode{
		Node: graph.Node{
			Label: "s",
		},
		Children: []*graph.BFSTreeNode{
			{
				Node: graph.Node{
					Label: "r",
				},
				Children: []*graph.BFSTreeNode{
					{
						Node: graph.Node{
							Label: "v",
						},
					},
				},
			},
			{
				Node: graph.Node{
					Label: "w",
				},
				Children: []*graph.BFSTreeNode{
					{
						Node: graph.Node{
							Label: "t",
						},
						Children: []*graph.BFSTreeNode{
							{
								Node: graph.Node{
									Label: "u",
								},
							},
						},
					},
					{
						Node: graph.Node{
							Label: "x",
						},
						Children: []*graph.BFSTreeNode{
							{
								Node: graph.Node{
									Label: "y",
								},
							},
						},
					},
				},
			},
		},
	}

	if !cmp.Equal(order1, expectedOrder) {
		t.Fatal(cmp.Diff(order1, expectedOrder))
	}
	if !cmp.Equal(order2, expectedOrder) {
		t.Fatal(cmp.Diff(order2, expectedOrder))
	}
	if !cmp.Equal(tree, expectedTree) {
		t.Fatal(cmp.Diff(tree, expectedTree))
	}
}
