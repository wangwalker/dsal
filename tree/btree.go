package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Using t to represent the half of degree of the B-tree,
// which means the maximum number of keys in a node is 2t-1.
const t = 2

type node[T constraints.Ordered] struct {
	keys     []T
	children []*node[T]
	isLeaf   bool
	level    int
}

// Search key in the B-tree.
func (tree *node[T]) search(key T) *node[T] {
	i := 0
	for i < len(tree.keys) && key > tree.keys[i] {
		i++
	}
	if i < len(tree.keys) && key == tree.keys[i] {
		return tree
	}
	if tree.isLeaf {
		return nil
	}
	return tree.children[i].search(key)
}

// Insert inserts a key into the B-tree, which is the outer interface.
func (tree *node[T]) insert(key T) *node[T] {
	i := len(tree.keys) - 1
	if tree.isLeaf {
		tree.keys = append(tree.keys, key)
		j := i
		for ; j >= 0 && key < tree.keys[j]; j-- {
			tree.keys[j+1] = tree.keys[j]
		}
		tree.keys[j+1] = key
		return tree
	}
	for i >= 0 && key < tree.keys[i] {
		i--
	}
	i++
	if len(tree.children[i].keys) == 2*t-1 {
		tree.splitChild(i, tree.children[i])
		// recalculate the index after split node
		i = len(tree.keys) - 1
		for i >= 0 && key < tree.keys[i] {
			i--
		}
		i++
	}
	return tree.children[i].insert(key)
}

// Split node when the number of the keys = [2*t-1].
// In this case, first split the original child into two pieces with the
// middle key, then constuct a new node with the middle key and two children,
// finally insert the new node into the  parent node.
func (parent *node[T]) splitChild(i int, child *node[T]) {
	// split original child into two pieces with the middle key,
	// child1, child2 = child[:t-1], child[t:]
	var child1, child2 *node[T]
	level := child.level + 1
	if child.isLeaf {
		child1 = &node[T]{
			keys:     child.keys[:t-1],
			children: nil,
			isLeaf:   child.isLeaf,
			level:    level,
		}
		child2 = &node[T]{
			keys:     child.keys[t:],
			children: nil,
			isLeaf:   child.isLeaf,
			level:    level,
		}
	} else {
		child1 = &node[T]{
			keys:     child.keys[:t-1],
			children: child.children[:t-1],
			isLeaf:   child.isLeaf,
			level:    level,
		}
		child2 = &node[T]{
			keys:     child.keys[t:],
			children: child.children[t:],
			isLeaf:   child.isLeaf,
			level:    level,
		}
	}

	subParent := &node[T]{
		keys:     []T{child.keys[t-1]},
		children: []*node[T]{child1, child2},
		isLeaf:   false,
		level:    child.level,
	}
	parent.children[i] = subParent
	parent.merge(subParent, i)
}

// Merge merges parent and child node when the number of parent's keys < 2*t-1.
// child node is the new node after split, so it has just one key and two children.
// It should be called after splitChild to balance tree.
func (parent *node[T]) merge(child *node[T], i int) {
	if len(parent.keys) == 2*t-1 {
		return
	}
	if i == 0 {
		parent.keys = append(child.keys, parent.keys...)
		parent.children = append(child.children, parent.children[1:]...)
	} else if len(parent.keys) > i {
		// splict parent keys two pieces, the middle one will is the only key of child node
		key1, key2 := parent.keys[:i], parent.keys[i:]
		keys := make([]T, 0, len(key1)+len(key2)+1)
		keys = append(keys, key1...)
		keys = append(keys, child.keys[0])
		keys = append(keys, key2...)
		parent.keys = keys
		// splict parent children two pieces, will ignore the middle one
		child1, child2 := parent.children[:i], parent.children[i+1:]
		children := make([]*node[T], 0, len(child1)+len(child2)+1)
		children = append(children, child1...)
		children = append(children, child.children...)
		children = append(children, child2...)
		parent.children = children
	} else {
		// just append the key and children of child node to parent
		parent.keys = append(parent.keys, child.keys[0])
		parent.children = append(parent.children[:i], child.children...)
	}
	// update level of children
	for _, c := range parent.children {
		c.level = parent.level + 1
	}
}

func traverse[T constraints.Ordered](tree *node[T]) {
	fmt.Printf("level = %d, keys = %+v\n", tree.level, tree.keys)
	for i := range tree.children {
		if !tree.isLeaf && tree.children[i] != nil {
			traverse(tree.children[i])
		}
	}
}