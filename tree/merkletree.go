package tree

import (
	"crypto/md5"
)

// MerkleTreeNode is the node of the merkle tree, left and right are the child
// nodes, data is the hash of the node.
type MerkleTreeNode struct {
	Left  *MerkleTreeNode
	Right *MerkleTreeNode
	Data  []byte
}

// MerkleTree is the merkle tree, root is the root node of the tree.
type MerkleTree struct {
	Root *MerkleTreeNode
}

// NewNode returns a new merkle tree node with the given left and right child
// nodes and data, if the node is a leaf node, the data is the hash of the data,
// otherwise the data is the hash of the left and right child nodes.
func NewNode(left, right *MerkleTreeNode, data []byte) *MerkleTreeNode {
	node := &MerkleTreeNode{
		Left:  left,
		Right: right,
		Data:  data,
	}
	if node.Left == nil && node.Right == nil {
		hash := md5.Sum(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(node.Left.Data, node.Right.Data...)
		hash := md5.Sum(prevHashes)
		node.Data = hash[:]
	}
	return node
}

// NewMerkleTree returns a new merkle tree with the given data, the data is the
// hash of the leaf nodes.
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []*MerkleTreeNode
	for _, d := range data {
		node := NewNode(nil, nil, d)
		nodes = append(nodes, node)
	}
	for len(nodes) > 1 {
		var newLevel []*MerkleTreeNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			var right *MerkleTreeNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			}
			newNode := NewNode(left, right, nil)
			newLevel = append(newLevel, newNode)
		}
		nodes = newLevel
	}
	return &MerkleTree{Root: nodes[0]}
}
