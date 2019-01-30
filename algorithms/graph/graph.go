package graph

type Graph struct {
	Nodes []Node
	Edges [][]Edge
}

type Node struct {
	Label string
}

type Edge struct {
	Target int
	Weight float64
}
