package tree

import (
	"crypto/md5"
	"testing"
)

// TestCreateMerkleTree tests creating a merkle tree.
func TestCreateMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
		[]byte("e"),
		[]byte("f"),
		[]byte("g"),
		[]byte("h"),
	}
	tree := NewMerkleTree(data)
	if tree.Root == nil {
		t.Error("expected root node not nil")
	}
}

// TestMerkleTreeWithOneNode tests creating a merkle tree with only one node.
func TestMerkleTreeWithOneNode(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
	}
	tree := NewMerkleTree(data)
	if tree.Root == nil {
		t.Error("expected root node not nil")
	}
	hash := md5.Sum(data[0])
	if hash != [16]byte(tree.Root.Data) {
		t.Error("expected root node data is not correct")
	}
}

// TestMerkleTreeShouldBeSame tests creating a merkle tree with the same data.
func TestMerkleTreeShouldBeSame(t *testing.T) {
	// GIVEN
	data1 := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}
	data2 := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	// WHEN
	tree1 := NewMerkleTree(data1)
	tree2 := NewMerkleTree(data2)

	// THEN
	if tree1.Root == nil {
		t.Error("expected root node not nil")
	}
	if tree2.Root == nil {
		t.Error("expected root node not nil")
	}
	if string(tree1.Root.Data) != string(tree2.Root.Data) {
		t.Error("expected root node data is not correct")
	}
}

// TestMerkleTreeShouldBeDifferent tests creating a merkle tree with different
// data.
func TestMerkleTreeShouldBeDifferent(t *testing.T) {
	// GIVEN
	data1 := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
		[]byte("e"),
		[]byte("f"),
		[]byte("g"),
		[]byte("h"),
	}
	data2 := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
		[]byte("e"),
		[]byte("f"),
		[]byte("g"),
		[]byte("k"),
	}

	// WHEN
	tree1 := NewMerkleTree(data1)
	tree2 := NewMerkleTree(data2)

	// THEN
	if tree1.Root == nil {
		t.Error("expected root node not nil")
	}
	if tree2.Root == nil {
		t.Error("expected root node not nil")
	}
	if string(tree1.Root.Data) == string(tree2.Root.Data) {
		t.Error("expected root node data is not correct")
	}
}
