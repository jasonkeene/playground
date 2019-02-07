package graph_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/graph"
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
						{1, 1},
						{2, 3},
					},
					{
						{3, 2},
					},
					{
						{4, 4},
					},
					{
						{4, 3},
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
						{1, 1},
						{3, 2},
					},
					{
						{0, 2},
						{2, 1},
						{4, 5},
					},
					{
						{1, 2},
						{5, 1},
					},
					{
						{0, 10},
						{4, 2},
						{6, 2},
					},
					{
						{1, 20},
						{3, 3},
						{5, 2},
						{7, 5},
					},
					{
						{2, 30},
						{4, 3},
						{8, 1},
					},
					{
						{3, 10},
						{7, 3},
					},
					{
						{4, 20},
						{6, 4},
						{8, 3},
					},
					{
						{5, 8},
						{7, 4},
					},
				},
			},
			Start:            7,
			ExpectedShortest: []float64{24, 25, 26, 14, 14, 11, 4, 0, 3},
			ExpectedPrev:     []int{3, 0, 1, 6, 5, 8, 7, -1, 7},
		},
	}

	algoFuncs := map[string](func(int, graph.Graph, graph.PriorityQueue) ([]float64, []int)){
		"dijkstra": graph.Dijkstra,
		"bellman-ford": func(s int, g graph.Graph, q graph.PriorityQueue) ([]float64, []int) {
			return graph.BellmanFord(s, g)
		},
	}

	queueFuncs := map[string](func() graph.PriorityQueue){
		"heap": func() graph.PriorityQueue {
			return graph.NewHeapQueue()
		},
		"array": func() graph.PriorityQueue {
			return graph.NewArrayQueue()
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
