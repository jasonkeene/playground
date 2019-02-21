package bst_test

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/jasonkeene/playground/data-structures/bst"
)

var _ bst.SearchTree = &bst.Node{}

func TestBST(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		nodeToFind := &bst.Node{
			Key:   25,
			Value: "test-value",
		}
		tree := ensureCorrectTree(&bst.Node{
			Key: 20,
			Left: &bst.Node{
				Key: 10,
			},
			Right: &bst.Node{
				Key:  30,
				Left: nodeToFind,
				Right: &bst.Node{
					Key: 35,
				},
			},
		})

		found := tree.Search(nodeToFind.Key)

		if !cmp.Equal(found, nodeToFind, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(found, nodeToFind), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("search for value that does not exist", func(t *testing.T) {
		nodeToFind := &bst.Node{
			Key:   25,
			Value: "test-value",
		}
		tree := ensureCorrectTree(&bst.Node{
			Key: 20,
			Left: &bst.Node{
				Key: 10,
			},
			Right: &bst.Node{
				Key: 30,
				Right: &bst.Node{
					Key: 35,
				},
			},
		})

		found := tree.Search(nodeToFind.Key)

		var emptyNode *bst.Node
		if !cmp.Equal(found, emptyNode, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(found, emptyNode), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("search empty node", func(t *testing.T) {
		var emptyNode *bst.Node

		found := emptyNode.Search(25)

		if !cmp.Equal(found, emptyNode, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(found, emptyNode), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("insert", func(t *testing.T) {
		var tree *bst.Node

		tree = tree.Insert(&bst.Node{
			Key: 20,
		})
		tree = tree.Insert(&bst.Node{
			Key: 30,
		})
		tree = tree.Insert(&bst.Node{
			Key: 10,
		})

		expectedTree := ensureCorrectTree(&bst.Node{
			Key: 20,
			Left: &bst.Node{
				Key: 10,
			},
			Right: &bst.Node{
				Key: 30,
			},
		})

		if !cmp.Equal(tree, expectedTree, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(tree, expectedTree), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("inserts new value when key already exists", func(t *testing.T) {
		newNode := &bst.Node{
			Key:   20,
			Value: "new-value",
		}
		tree := ensureCorrectTree(&bst.Node{
			Key:   20,
			Value: "old-value",
			Left: &bst.Node{
				Key: 10,
			},
			Right: &bst.Node{
				Key: 30,
				Left: &bst.Node{
					Key: 25,
				},
				Right: &bst.Node{
					Key: 35,
				},
			},
		})
		tree = tree.Insert(newNode)

		if !cmp.Equal(tree.Value, newNode.Value) {
			t.Fatal(cmp.Diff(tree.Value, newNode.Value))
		}
	})

	t.Run("minimum", func(t *testing.T) {
		expectedMinNode := &bst.Node{
			Key:   5,
			Value: "min-node",
		}
		tree := ensureCorrectTree(&bst.Node{
			Key: 20,
			Left: &bst.Node{
				Key:  10,
				Left: expectedMinNode,
			},
			Right: &bst.Node{
				Key: 30,
				Left: &bst.Node{
					Key: 25,
				},
				Right: &bst.Node{
					Key: 35,
				},
			},
		})

		minNode := tree.Minimum()

		if !cmp.Equal(minNode, expectedMinNode, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(minNode, expectedMinNode), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("maximum", func(t *testing.T) {
		expectedMaxNode := &bst.Node{
			Key:   50,
			Value: "max-node",
		}
		tree := ensureCorrectTree(&bst.Node{
			Key: 20,
			Left: &bst.Node{
				Key: 10,
			},
			Right: &bst.Node{
				Key: 30,
				Left: &bst.Node{
					Key: 25,
				},
				Right: &bst.Node{
					Key:   35,
					Right: expectedMaxNode,
				},
			},
		})

		maxNode := tree.Maximum()

		if !cmp.Equal(maxNode, expectedMaxNode, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(maxNode, expectedMaxNode), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("successor in subtree", func(t *testing.T) {
		expectedSuccessor := &bst.Node{
			Key:   9,
			Value: "successor",
		}

		tree := ensureCorrectTree(&bst.Node{
			Key: 7,
			Right: &bst.Node{
				Key:  13,
				Left: expectedSuccessor,
			},
		})

		successor := tree.Successor()

		if !cmp.Equal(successor, expectedSuccessor, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(successor, expectedSuccessor), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("successor in ancestry", func(t *testing.T) {
		node := &bst.Node{
			Key: 13,
		}
		expectedSuccessor := &bst.Node{
			Key:   15,
			Value: "successor",
			Left: &bst.Node{
				Key: 6,
				Right: &bst.Node{
					Key:   7,
					Right: node,
				},
			},
		}

		ensureCorrectTree(&bst.Node{
			Key:  27,
			Left: expectedSuccessor,
		})

		successor := node.Successor()

		if !cmp.Equal(successor, expectedSuccessor, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(successor, expectedSuccessor, cmp.Comparer(EqualNodes)))
		}
	})

	t.Run("predecessor in subtree", func(t *testing.T) {
		expectedPredecessor := &bst.Node{
			Key:   4,
			Value: "predecessor",
		}

		tree := ensureCorrectTree(&bst.Node{
			Key: 6,
			Left: &bst.Node{
				Key:   3,
				Right: expectedPredecessor,
			},
		})

		predecessor := tree.Predecessor()

		if !cmp.Equal(predecessor, expectedPredecessor, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(predecessor, expectedPredecessor), cmp.Comparer(EqualNodes))
		}
	})

	t.Run("predecessor in ancestry", func(t *testing.T) {
		node := &bst.Node{
			Key: 42,
		}
		expectedPredecessor := &bst.Node{
			Key:   40,
			Value: "predecessor",
			Right: &bst.Node{
				Key: 50,
				Left: &bst.Node{
					Key:  45,
					Left: node,
				},
			},
		}

		ensureCorrectTree(&bst.Node{
			Key:   30,
			Right: expectedPredecessor,
		})

		predecessor := node.Predecessor()

		if !cmp.Equal(predecessor, expectedPredecessor, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(predecessor, expectedPredecessor, cmp.Comparer(EqualNodes)))
		}
	})

	t.Run("delete isolated", func(t *testing.T) {
		tree := &bst.Node{
			Key: 20,
		}

		tree = tree.Delete()

		var emptyNode *bst.Node
		if !cmp.Equal(tree, emptyNode, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(tree, emptyNode, cmp.Comparer(EqualNodes)))
		}
	})

	t.Run("delete with no children", func(t *testing.T) {
		toDelete := &bst.Node{
			Key: 20,
		}

		tree := ensureCorrectTree(&bst.Node{
			Key:   10,
			Right: toDelete,
		})

		replacedBy := toDelete.Delete()

		var emptyNode *bst.Node
		if !cmp.Equal(replacedBy, emptyNode, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(replacedBy, emptyNode, cmp.Comparer(EqualNodes)))
		}
		if !cmp.Equal(tree.Right, emptyNode, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(tree.Right, emptyNode, cmp.Comparer(EqualNodes)))
		}
	})

	t.Run("delete with left child", func(t *testing.T) {
		tree := ensureCorrectTree(&bst.Node{
			Key: 10,
			Right: &bst.Node{
				Key: 20,
				Left: &bst.Node{
					Key: 15,
				},
			},
		})
		expectedTree := ensureCorrectTree(&bst.Node{
			Key: 10,
			Right: &bst.Node{
				Key: 15,
			},
		})

		replacedBy := tree.Right.Delete()

		if !cmp.Equal(replacedBy, expectedTree.Right, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(replacedBy, expectedTree.Right, cmp.Comparer(EqualNodes)))
		}
		if !cmp.Equal(tree, expectedTree, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(tree, expectedTree, cmp.Comparer(EqualNodes)))
		}
	})

	t.Run("delete with right child", func(t *testing.T) {
		expectedReplacedBy := &bst.Node{
			Key: 30,
		}
		toDelete := &bst.Node{
			Key:   20,
			Right: expectedReplacedBy,
		}

		tree := ensureCorrectTree(&bst.Node{
			Key:   10,
			Right: toDelete,
		})

		replacedBy := toDelete.Delete()

		if !cmp.Equal(replacedBy, expectedReplacedBy, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(replacedBy, expectedReplacedBy, cmp.Comparer(EqualNodes)))
		}
		if !cmp.Equal(tree.Right, expectedReplacedBy, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(tree.Right, expectedReplacedBy, cmp.Comparer(EqualNodes)))
		}
	})

	t.Run("delete with both children and successor is a child", func(t *testing.T) {
		expectedReplacedBy := &bst.Node{
			Key: 30,
		}
		toDelete := &bst.Node{
			Key:   20,
			Right: expectedReplacedBy,
			Left: &bst.Node{
				Key: 15,
			},
		}

		tree := ensureCorrectTree(&bst.Node{
			Key:   10,
			Right: toDelete,
		})

		replacedBy := toDelete.Delete()

		if !cmp.Equal(replacedBy, expectedReplacedBy, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(replacedBy, expectedReplacedBy, cmp.Comparer(EqualNodes)))
		}
		if !cmp.Equal(tree.Right, expectedReplacedBy, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(tree.Right, expectedReplacedBy, cmp.Comparer(EqualNodes)))
		}
	})

	t.Run("delete with both children and successor is not a child", func(t *testing.T) {
		tree := ensureCorrectTree(&bst.Node{
			Key: 10,
			Right: &bst.Node{
				Key:   20,
				Value: "getting-deleted",
				Left: &bst.Node{
					Key: 15,
				},
				Right: &bst.Node{
					Key: 50,
					Left: &bst.Node{
						Key:   30,
						Value: "replacing",
						Right: &bst.Node{
							Key: 40,
						},
					},
				},
			},
		})
		expectedReplacedBy := ensureCorrectTree(&bst.Node{
			Key:   30,
			Value: "replacing",
			Left: &bst.Node{
				Key: 15,
			},
			Right: &bst.Node{
				Key: 50,
				Left: &bst.Node{
					Key: 40,
				},
			},
		})
		expectedTree := ensureCorrectTree(&bst.Node{
			Key:   10,
			Right: expectedReplacedBy,
		})

		replacedBy := tree.Right.Delete()

		if !cmp.Equal(replacedBy, expectedReplacedBy, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(replacedBy, expectedReplacedBy, cmp.Comparer(EqualNodes)))
		}
		if !cmp.Equal(tree, expectedTree, cmp.Comparer(EqualNodes)) {
			t.Fatal(cmp.Diff(tree, expectedTree, cmp.Comparer(EqualNodes)))
		}
	})
}

func ensureCorrectTree(n *bst.Node) *bst.Node {
	if n.Left != nil {
		n.Left.Parent = n
		if n.Left.Key >= n.Key {
			log.Fatal("incorrect tree")
		}
		ensureCorrectTree(n.Left)
	}
	if n.Right != nil {
		n.Right.Parent = n
		if n.Right.Key <= n.Key {
			log.Fatal("incorrect tree")
		}
		ensureCorrectTree(n.Right)
	}
	return n
}

func EqualNodes(x, y *bst.Node) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	eq := cmp.Equal(x, y, cmpopts.IgnoreFields(*x, "Parent", "Left", "Right"))
	if !eq {
		return false
	}

	if x.Parent == nil && y.Parent != nil {
		return false
	}
	if x.Parent != nil && y.Parent == nil {
		return false
	}
	parentEq := cmp.Equal(x.Parent, y.Parent, cmpopts.IgnoreFields(*x, "Parent", "Left", "Right"))
	if !parentEq {
		return false
	}

	if !EqualNodes(x.Left, y.Left) {
		return false
	}
	return EqualNodes(x.Right, y.Right)
}
