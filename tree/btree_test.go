package tree

import "testing"

// Tests lt of Key when the type is string.
func TestStringKeyLt(t *testing.T) {
	// GIVEN
	k1 := Key[string]{"a", 1}
	k2 := Key[string]{"b", 2}

	// WHEN
	lt := k1.lt(k2)

	// THEN
	if !lt {
		t.Errorf("k1.lt(k2) = %v, want true", lt)
	}
}

// Tests lt of Key when the type is int.
func TestIntKeyLt(t *testing.T) {
	// GIVEN
	k1 := Key[int]{1, 1}
	k2 := Key[int]{2, 2}

	// WHEN
	lt := k1.lt(k2)

	// THEN
	if !lt {
		t.Errorf("k1.lt(k2) = %v, want true", lt)
	}
}

func TestStringBtree(t *testing.T) {
	// GIVEN
	root := &BTNode[string]{
		keys: []Key[string]{{"e", 0}, {"k", 30}},
		children: []*BTNode[string]{
			{keys: []Key[string]{{"a", 1}, {"b", 2}, {"v", 3}}, isLeaf: true, level: 2},
			{keys: []Key[string]{{"fd", 4}, {"gd", 5}, {"h2", 6}}, isLeaf: true, level: 2},
			{keys: []Key[string]{{"m1", 7}, {"m2", 8}, {"root", 9}}, isLeaf: true, level: 2}},
		isLeaf: false,
		level:  1,
	}
	t.Log("traversing the tree before insertion")
	traverse(root)

	// WHEN
	root.Insert(Key[string]{"food", 10})
	root.Insert(Key[string]{"godd", 11})
	root.Insert(Key[string]{"hi", 12})
	root.Insert(Key[string]{"internet", 13})
	root.Insert(Key[string]{"j", 14})
	root.Insert(Key[string]{"kitty", 15})
	root.Insert(Key[string]{"loop", 16})
	root.Insert(Key[string]{"moon", 17})
	root.Insert(Key[string]{"string", 18})

	// THEN
	t.Log("traversing the tree after inserting food, godd, hi, internet, j, kitty, loop, moon, string")
	traverse(root)

	if len(root.keys) != 3 {
		t.Errorf("root should have 3 keys, but got %d", len(root.keys))
	}
	if len(root.children) != 4 {
		t.Errorf("root should have 4 children, but got %d", len(root.children))
	}
	if v := root.Search("food"); v != 10 {
		t.Error("food should be found")
	}
	if v := root.Search("kitty"); v != 15 {
		t.Error("kitty should be found")
	}
	if v := root.Search("internet"); v != 13 {
		t.Error("internet should be found")
	}
	if v := root.Search("string"); v != 18 {
		t.Error("string should be found")
	}
	if v := root.Search("loop"); v != 16 {
		t.Error("loop should be found")
	}
	if v := root.Search("hi"); v != 12 {
		t.Error("hi should be found")
	}
	if v := root.Search("f"); v != -1 {
		t.Error("f should not be found")
	}
	if v := root.Search("z"); v != -1 {
		t.Error("z should not be found")
	}
}
