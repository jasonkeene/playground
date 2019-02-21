package bst

type SearchTree interface {
	Search(key float64) *Node
	Minimum() *Node
	Maximum() *Node
	Successor() *Node
	Predecessor() *Node

	Insert(*Node) *Node
	Delete() *Node
}

type Node struct {
	Key    float64
	Value  interface{}
	Parent *Node
	Left   *Node
	Right  *Node
}

func (n *Node) Search(key float64) *Node {
	if n == nil || key == n.Key {
		return n
	}
	if key < n.Key {
		return n.Left.Search(key)
	}
	return n.Right.Search(key)
}

func (n *Node) Minimum() *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (n *Node) Maximum() *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (n *Node) Successor() *Node {
	if n.Right != nil {
		return n.Right.Minimum()
	}
	prev := n
	for n.Parent != nil && n.Left != prev {
		prev = n
		n = n.Parent
	}
	return n
}

func (n *Node) Predecessor() *Node {
	if n.Left != nil {
		return n.Left.Maximum()
	}
	prev := n
	for n.Parent != nil && n.Right != prev {
		prev = n
		n = n.Parent
	}
	return n
}

func (n *Node) Insert(x *Node) *Node {
	if n == nil {
		return x
	}
	if x.Key == n.Key {
		n.Value = x.Value
		return n
	}
	x.Parent = n
	if x.Key < n.Key {
		n.Left = n.Left.Insert(x)
		return n
	}
	n.Right = n.Right.Insert(x)
	return n
}

func (n *Node) Delete() *Node {
	if n.Right == nil {
		transplant(n, n.Left)
		return n.Left
	}
	if n.Left == nil {
		transplant(n, n.Right)
		return n.Right
	}

	successor := n.Successor()

	if n.Right != successor {
		transplant(successor, successor.Right)
		successor.Right = n.Right
		successor.Right.Parent = successor
	}
	transplant(n, successor)
	successor.Left = n.Left
	successor.Left.Parent = successor
	return successor
}

func transplant(a, b *Node) {
	if a.Parent != nil {
		if a.Parent.Left == a {
			a.Parent.Left = b
		} else {
			a.Parent.Right = b
		}
	}

	if b != nil {
		b.Parent = a.Parent
	}
}
