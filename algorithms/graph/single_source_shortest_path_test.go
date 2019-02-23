package graph_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/graph"
	"github.com/jasonkeene/playground/data-structures/queue"
)

func TestSingleSourceShortestPath(t *testing.T) {
	testCases := map[string]struct {
		Graph            graph.Graph
		Start            int
		ExpectedShortest []float64
		ExpectedPrev     []int
	}{
		"basic graph": {
			Graph: graph.Graph{
				Nodes: []graph.Node{
					{"A"},
					{"B"},
					{"C"},
					{"D"},
					{"E"},
				},
				Edges: [][]graph.Edge{
					{
						{1, 1, 0},
						{2, 3, 0},
					},
					{
						{3, 2, 0},
					},
					{
						{4, 4, 0},
					},
					{
						{4, 3, 0},
					},
					{},
				},
			},
			Start:            0,
			ExpectedShortest: []float64{0, 1, 3, 3, 6},
			ExpectedPrev:     []int{-1, 0, 0, 1, 3},
		},
		"traffic graph": {
			Graph: graph.Graph{
				Nodes: []graph.Node{
					{"A"},
					{"B"},
					{"C"},
					{"D"},
					{"E"},
					{"F"},
					{"G"},
					{"H"},
					{"I"},
				},
				Edges: [][]graph.Edge{
					{
						{1, 1, 0},
						{3, 2, 0},
					},
					{
						{0, 2, 0},
						{2, 1, 0},
						{4, 5, 0},
					},
					{
						{1, 2, 0},
						{5, 1, 0},
					},
					{
						{0, 10, 0},
						{4, 2, 0},
						{6, 2, 0},
					},
					{
						{1, 20, 0},
						{3, 3, 0},
						{5, 2, 0},
						{7, 5, 0},
					},
					{
						{2, 30, 0},
						{4, 3, 0},
						{8, 1, 0},
					},
					{
						{3, 10, 0},
						{7, 3, 0},
					},
					{
						{4, 20, 0},
						{6, 4, 0},
						{8, 3, 0},
					},
					{
						{5, 8, 0},
						{7, 4, 0},
					},
				},
			},
			Start:            7,
			ExpectedShortest: []float64{24, 25, 26, 14, 14, 11, 4, 0, 3},
			ExpectedPrev:     []int{3, 0, 1, 6, 5, 8, 7, -1, 7},
		},
	}

	algoFuncs := map[string](func(int, graph.Graph, queue.Priority) ([]float64, []int)){
		"dijkstra": graph.Dijkstra,
		"bellman-ford": func(s int, g graph.Graph, q queue.Priority) ([]float64, []int) {
			return graph.BellmanFord(s, g)
		},
	}

	queueFuncs := map[string](func() queue.Priority){
		"heap": func() queue.Priority {
			return queue.NewHeap()
		},
		"array": func() queue.Priority {
			return queue.NewArray()
		},
	}

	for ak, af := range algoFuncs {
		t.Run(ak, func(t *testing.T) {
			for qk, qf := range queueFuncs {
				t.Run(qk, func(t *testing.T) {
					for k, tc := range testCases {
						t.Run(k, func(t *testing.T) {
							shortest, prev := af(
								tc.Start,
								tc.Graph,
								qf(),
							)

							if !cmp.Equal(shortest, tc.ExpectedShortest) {
								t.Error(cmp.Diff(shortest, tc.ExpectedShortest))
							}
							if !cmp.Equal(prev, tc.ExpectedPrev) {
								t.Error(cmp.Diff(prev, tc.ExpectedPrev))
							}
						})
					}
				})
			}
		})
	}
}
