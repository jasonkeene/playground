package graph

func BFS(g Graph, s int, f func(Node)) {
	enqueued := make([]bool, len(g.Nodes))
	fifo := make(chan int, len(g.Nodes))
	fifo <- s
	enqueued[s] = true

	for len(fifo) != 0 {
		n := <-fifo
		if f != nil {
			f(g.Nodes[n])
		}
		for _, e := range g.Edges[n] {
			if !enqueued[e.Target] {
				fifo <- e.Target
				enqueued[e.Target] = true
			}
		}
	}
}

type BFSTreeNode struct {
	Node
	Children []*BFSTreeNode
}

func BFSTree(g Graph, s int, f func(Node)) *BFSTreeNode {
	var tree *BFSTreeNode
	bfsTreeNodes := make([]*BFSTreeNode, len(g.Nodes))
	tree = &BFSTreeNode{
		Node: g.Nodes[s],
	}
	bfsTreeNodes[s] = tree

	enqueued := make([]bool, len(g.Nodes))
	fifo := make(chan int, len(g.Nodes))
	fifo <- s
	enqueued[s] = true

	for len(fifo) != 0 {
		n := <-fifo
		if f != nil {
			f(g.Nodes[n])
		}
		for _, e := range g.Edges[n] {
			if !enqueued[e.Target] {
				bfsNode := &BFSTreeNode{
					Node: g.Nodes[e.Target],
				}
				bfsTreeNodes[e.Target] = bfsNode
				bfsTreeNodes[n].Children = append(
					bfsTreeNodes[n].Children,
					bfsNode,
				)
				fifo <- e.Target
				enqueued[e.Target] = true
			}
		}
	}

	return tree
}
