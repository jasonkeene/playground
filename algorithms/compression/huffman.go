package compression

import "github.com/jasonkeene/playground/algorithms/graph"

type Node struct {
	Char  byte
	Value int
	Left  *Node
	Right *Node
}

func HuffmanTree(chars []byte, freqs []int) *Node {
	q := graph.NewHeapQueue()

	for i := range chars {
		q.Insert(graph.QueueElement{
			Key: float64(freqs[i]),
			Value: &Node{
				Char:  chars[i],
				Value: freqs[i],
			},
		})
	}

	var root *Node
	for {
		node1 := q.PopMin().Value.(*Node)
		if q.Empty() {
			root = node1
			break
		}
		node2 := q.PopMin().Value.(*Node)

		node := &Node{
			Value: node1.Value + node2.Value,
			Left:  node1,
			Right: node2,
		}

		q.Insert(graph.QueueElement{
			Key:   float64(node.Value),
			Value: node,
		})
	}

	return root
}
