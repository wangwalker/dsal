package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// AVLNode is the node of the AVL tree.
type AVLNode[T constraints.Ordered] struct {
	key   T
	level int
	left  *AVLNode[T]
	right *AVLNode[T]
}

// height returns the level of the AVL tree when node is not nil.
func (tree *AVLNode[T]) height() int {
	if tree == nil {
		return 0
	}
	return tree.level
}

// NewAVLNode creates a new AVL node.
func NewAVLNode[T constraints.Ordered](k T) *AVLNode[T] {
	return &AVLNode[T]{key: k, level: 1}
}

// Insert inserts a key into the AVL tree.
func (tree *AVLNode[T]) Insert(k T) *AVLNode[T] {
	if tree == nil {
		return NewAVLNode(k)
	}
	if k < tree.key {
		tree.left = tree.left.Insert(k)
	} else if k > tree.key {
		tree.right = tree.right.Insert(k)
	} else {
		return tree
	}
	// rebalance the tree
	tree.level = max(tree.left.height(), tree.right.height()) + 1
	balance := tree.left.height() - tree.right.height()
	if balance > 1 {
		// left-left case
		if k < tree.left.key {
			return tree.rightRotate()
		} else {
			// left-right case
			tree.left = tree.left.leftRotate()
			return tree.rightRotate()
		}
	} else if balance < -1 {
		// right-right case
		if k > tree.right.key {
			return tree.leftRotate()
		} else {
			// right-left case
			tree.right = tree.right.rightRotate()
			return tree.leftRotate()
		}
	}
	return tree
}

// Delete deletes a key from the AVL tree.
func (tree *AVLNode[T]) Delete(k T) *AVLNode[T] {
	if tree == nil {
		return nil
	}
	if k < tree.key {
		tree.left = tree.left.Delete(k)
	} else if k > tree.key {
		tree.right = tree.right.Delete(k)
	} else {
		if tree.left == nil {
			return tree.right
		} else if tree.right == nil {
			return tree.left
		}
		// the deleted node has two children
		min := tree.right.min()
		tree.key = min.key
		tree.right = tree.right.Delete(min.key)
	}
	// rebalance the tree
	tree.level = max(tree.left.height(), tree.right.height()) + 1
	balance := tree.left.height() - tree.right.height()
	if balance > 1 {
		if tree.left.left.height() >= tree.left.right.height() {
			return tree.rightRotate()
		} else {
			tree.left = tree.left.leftRotate()
			return tree.rightRotate()
		}
	} else if balance < -1 {
		if tree.right.right.height() >= tree.right.left.height() {
			return tree.leftRotate()
		} else {
			tree.right = tree.right.rightRotate()
			return tree.leftRotate()
		}
	}
	return tree
}

// Search searches key in the AVL tree.
func (tree *AVLNode[T]) Search(k T) *AVLNode[T] {
	if tree == nil {
		return nil
	}
	if k < tree.key {
		return tree.left.Search(k)
	} else if k > tree.key {
		return tree.right.Search(k)
	}
	return tree
}

// min returns the minimum node in the AVL tree.
func (tree *AVLNode[T]) min() *AVLNode[T] {
	if tree.left == nil {
		return tree
	}
	return tree.left.min()
}

// leftRotate performs a left rotation.
func (tree *AVLNode[T]) leftRotate() *AVLNode[T] {
	x := tree.right
	tree.right = x.left
	x.left = tree
	tree.level = max(tree.left.height(), tree.right.height()) + 1
	x.level = max(x.left.height(), x.right.height()) + 1
	return x
}

// rightRotate performs a right rotation.
func (tree *AVLNode[T]) rightRotate() *AVLNode[T] {
	x := tree.left
	tree.left = x.right
	x.right = tree
	tree.level = max(tree.left.height(), tree.right.height()) + 1
	x.level = max(x.left.height(), x.right.height()) + 1
	return x
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// traverse traverses the AVL tree in-order.
func (tree *AVLNode[T]) traverse() {
	if tree == nil {
		return
	}
	tree.left.traverse()
	fmt.Printf("%v at level %v", tree.key, tree.height())
	tree.right.traverse()
}
