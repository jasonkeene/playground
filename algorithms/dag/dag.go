package dag

type Vertex struct {
	Label string
	Edges []Edge
}

type Edge struct {
	Weight float64
	Target int
}
