package tree

import "testing"

// TestIntAvlTree tests the very simple AVL tree with int type.
func TestIntAvlTree(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)

	// WHEN
	tree.traverse()

	// THEN
	if tree.key != 20 {
		t.Errorf("tree.key != 20")
	}
	if tree.height() != 2 {
		t.Errorf("tree.height != 2")
	}
	if condition := tree.left.key == 10 && tree.right.key == 30; !condition {
		t.Errorf("tree.left.key == 10 && tree.right.key == 30")
	}

}

func TestIntAvlTreeWhenHasMoreNodes(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)
	tree = tree.Insert(40)
	tree = tree.Insert(50)
	tree = tree.Insert(5)
	tree = tree.Insert(1)

	// WHEN
	tree.traverse()

	// THEN
	if tree.key != 20 {
		t.Errorf("tree.key != 20")
	}
	if tree.height() != 3 {
		t.Errorf("tree.height != 3")
	}
	if condition := tree.left.key == 5 && tree.right.key == 40; !condition {
		t.Errorf("tree.left.key == 5 && tree.right.key == 40")
	}
	if condition := tree.left.left.key == 1 && tree.left.right.key == 10; !condition {
		t.Errorf("tree.left.left.key == 1 && tree.left.right.key == 10")
	}
	if condition := tree.right.left.key == 30 && tree.right.right.key == 50; !condition {
		t.Errorf("tree.right.left.key == 30 && tree.right.right.key == 50")
	}
}

func TestDeleteIntAvlTreeWhenTargetHasOnlyOneChild(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)
	tree = tree.Insert(40)
	tree = tree.Insert(50)
	tree = tree.Insert(5)
	tree = tree.Insert(1)
	tree = tree.Insert(12)

	// WHEN
	tree.Delete(1)

	// THEN
	if tree.key != 20 {
		t.Errorf("tree.key != 20")
	}
	if tree.height() != 3 {
		t.Errorf("tree.height != 3")
	}
	if condition := tree.left.key == 10 && tree.right.key == 40; !condition {
		t.Errorf("tree.left.key == 5 && tree.right.key == 40")
	}
	if condition := tree.left.left.key == 5 && tree.left.right.key == 12; !condition {
		t.Errorf("tree.left.left.key == 1 && tree.left.right.key == 10")
	}
	if condition := tree.right.left.key == 30 && tree.right.right.key == 50; !condition {
		t.Errorf("tree.right.left.key == 30 && tree.right.right.key == 50")
	}
}

func TestDeleteIntAvlTreeWhenTargetNodeHasTwoChildren(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)
	tree = tree.Insert(40)
	tree = tree.Insert(50)
	tree = tree.Insert(5)
	tree = tree.Insert(1)

	// WHEN
	tree.Delete(5)

	// THEN
	if tree.key != 20 {
		t.Errorf("tree.key != 20")
	}
	if tree.height() != 3 {
		t.Errorf("tree.height != 3")
	}
	if condition := (tree.left.key == 1 || tree.left.key == 10) && tree.right.key == 40; !condition {
		t.Errorf("tree.left.key == 1|10 && tree.right.key == 40")
	}
	if condition := tree.left.left != nil; !condition {
		t.Errorf("tree.left.left.key == 1 && tree.left.right.key == 10")
	}
}

func TestDeleteIntAvlTreeWhenNotContainTargetNode(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)
	tree = tree.Insert(40)
	tree = tree.Insert(50)
	tree = tree.Insert(5)
	tree = tree.Insert(1)

	// WHEN
	tree.Delete(22)

	// THEN
	if tree.key != 20 {
		t.Errorf("tree.key != 20")
	}
	if tree.height() != 3 {
		t.Errorf("tree.height != 3")
	}
}

//
func TestSearchIntAvlTreeWhenContainsTargetNode(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)
	tree = tree.Insert(40)
	tree = tree.Insert(50)
	tree = tree.Insert(5)
	tree = tree.Insert(1)

	// WHEN
	node := tree.Search(5)

	// THEN
	if node.key != 5 {
		t.Errorf("node.key != 5")
	}
}

func TestSearchIntAvlTreeWhenNotContainTargetNode(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)
	tree = tree.Insert(40)
	tree = tree.Insert(50)
	tree = tree.Insert(5)
	tree = tree.Insert(1)

	// WHEN
	node := tree.Search(2)

	// THEN
	if node != nil {
		t.Errorf("node != nil")
	}
}

func TestMinIntAvlTree(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(10)
	tree = tree.Insert(20)
	tree = tree.Insert(30)
	tree = tree.Insert(40)
	tree = tree.Insert(50)
	tree = tree.Insert(5)
	tree = tree.Insert(1)

	// WHEN
	node := tree.min()

	// THEN
	if node.key != 1 {
		t.Errorf("node.key != 1")
	}
}

func TestLeftRotateIntTree(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(1)
	tree.right = NewAVLNode(2)
	tree.right.right = NewAVLNode(3)

	// WHEN
	tree = tree.leftRotate()

	// THEN
	if tree.key != 2 {
		t.Errorf("tree.key != 2")
	}
	if tree.height() != 2 {
		t.Errorf("tree.height != 2")
	}
	if condition := tree.left.key == 1 && tree.right.key == 3; !condition {
		t.Errorf("tree.left.key == 1 && tree.right.key == 3")
	}
}

func TestRightRotateIntTree(t *testing.T) {
	// GIVEN
	tree := NewAVLNode(3)
	tree.left = NewAVLNode(2)
	tree.left.left = NewAVLNode(1)

	// WHEN
	tree = tree.rightRotate()

	// THEN
	if tree.key != 2 {
		t.Errorf("tree.key != 2")
	}
	if tree.height() != 2 {
		t.Errorf("tree.height != 2")
	}
	if condition := tree.left.key == 1 && tree.right.key == 3; !condition {
		t.Errorf("tree.left.key == 1 && tree.right.key == 3")
	}
}

// TestStringAvlTree tests the AVL tree with string type.
func TestStringAvlTree(t *testing.T) {
	// GIVEN
	tree := NewAVLNode("a")
	tree = tree.Insert("b")
	tree = tree.Insert("c")
	tree = tree.Insert("d")
	tree = tree.Insert("e")
	tree = tree.Insert("f")
	tree = tree.Insert("g")

	// WHEN
	tree.traverse()

	// THEN
	if tree.key != "d" {
		t.Errorf("tree.key != d")
	}
	if tree.height() != 3 {
		t.Errorf("tree.height != 3")
	}
	if condition := tree.left.key == "b" && tree.right.key == "f"; !condition {
		t.Errorf("tree.left.key == b && tree.right.key == f")
	}
	if condition := tree.left.left.key == "a" && tree.left.right.key == "c"; !condition {
		t.Errorf("tree.left.left.key == a && tree.left.right.key == c")
	}
	if condition := tree.right.left.key == "e" && tree.right.right.key == "g"; !condition {
		t.Errorf("tree.right.left.key == e && tree.right.right.key == g")
	}
}
