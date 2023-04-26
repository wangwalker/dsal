package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// RBNode is the node of the red-black tree.
type RBNode[T constraints.Ordered] struct {
	key    T
	value  interface{}
	left   *RBNode[T]
	right  *RBNode[T]
	parent *RBNode[T]
	color  bool // true: red, false: black
}

// NewRBNode creates a new red-black node.
func NewRBNode[T constraints.Ordered](k T, v interface{}) *RBNode[T] {
	return &RBNode[T]{key: k, value: v}
}

// Insert inserts a key into the red-black tree.
func (tree *RBNode[T]) Insert(k T, v interface{}) *RBNode[T] {
	if tree == nil {
		return NewRBNode(k, v)
	}
	if k < tree.key {
		tree.left = tree.left.Insert(k, v)
		tree.left.parent = tree
	} else if k > tree.key {
		tree.right = tree.right.Insert(k, v)
		tree.right.parent = tree
	} else {
		return tree
	}
	// rebalance the tree
	if tree.left != nil && tree.left.color {
		if tree.left.left != nil && tree.left.left.color {
			tree = tree.rightRotate()
		} else if tree.left.right != nil && tree.left.right.color {
			tree.left = tree.left.leftRotate()
			tree = tree.rightRotate()
		}
	} else if tree.right != nil && tree.right.color {
		if tree.right.right != nil && tree.right.right.color {
			tree = tree.leftRotate()
		} else if tree.right.left != nil && tree.right.left.color {
			tree.right = tree.right.rightRotate()
			tree = tree.leftRotate()
		}
	}
	return tree
}

// leftRotate rotates the tree to the left.
func (tree *RBNode[T]) leftRotate() *RBNode[T] {
	right := tree.right
	tree.right = right.left
	right.left = tree
	right.color = tree.color
	tree.color = true
	return right
}

// rightRotate rotates the tree to the right.
func (tree *RBNode[T]) rightRotate() *RBNode[T] {
	left := tree.left
	tree.left = left.right
	left.right = tree
	left.color = tree.color
	tree.color = true
	return left
}

// traverse traverses the tree in-order.
//
//lint:ignore U1000 Ignore unused function temporarily for debugging
func (tree *RBNode[T]) traverse() {
	if tree == nil {
		return
	}
	tree.left.traverse()
	if tree.color {
		fmt.Printf("key: %v, color: %s", tree.key, "red")
	} else {
		fmt.Printf("key: %v, color: %s", tree.key, "black")
	}
	tree.right.traverse()
}
