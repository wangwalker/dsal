package linear

import "testing"

type mockDicisionMaker struct {
	shouldInsert bool
}

func (m *mockDicisionMaker) ShouldInsert() bool {
	return m.shouldInsert
}

func TestCreateSkipListNode(t *testing.T) {
	head := NewSkipList(1, "a")
	if head.key != 1 || head.value != "a" || head.right != nil || head.down != nil {
		t.Errorf("NewSkipList(1, \"a\") = %v, want %v", head, SLNode[int]{1, "a", nil, nil, &RandomDicisionMaker{}})
	}
}

func TestInsertAndSearchSkipListNode(t *testing.T) {
	// GIVEN
	head := NewSkipList(1, "a")
	head = head.Insert(2, "b")
	head = head.Insert(3, "c")
	head = head.Insert(4, "d")
	head = head.Insert(5, "e")

	// WHEN
	r1 := head.Search(1)
	r2 := head.Search(2)
	r3 := head.Search(3)
	r4 := head.Search(4)
	r5 := head.Search(5)

	// THEN
	if r1 != "a" {
		t.Errorf("head.Search(1) = %v, want \"a\"", r1)
	}
	if r2 != "b" {
		t.Errorf("head.Search(2) = %v, want \"b\"", r2)
	}
	if r3 != "c" {
		t.Errorf("head.Search(3) = %v, want \"c\"", r3)
	}
	if r4 != "d" {
		t.Errorf("head.Search(4) = %v, want \"d\"", r4)
	}
	if r5 != "e" {
		t.Errorf("head.Search(5) = %v, want \"e\"", r5)
	}
}

func TestUpdateSkipListNode(t *testing.T) {
	// GIVEN
	head := NewSkipList(1, "a")
	head = head.Insert(2, "b")
	head = head.Insert(3, "c")

	// WHEN
	head.Update(2, "d")

	// THEN
	if head.Search(2) != "d" {
		t.Errorf("head.Search(2) = %v, want \"d\"", head.Search(2))
	}
}

func TestUpdateSkipListHeadNode(t *testing.T) {
	// GIVEN
	head := NewSkipList(1, "a")
	head = head.Insert(12, "b")
	head = head.Insert(13, "c")
	head = head.Insert(14, "d")

	// WHEN
	head.Update(1, "k")

	// THEN
	if head.Search(1) != "k" {
		t.Errorf("head.Search(1) = %v, want \"k\"", head.Search(1))
	}
	for p := head; p != nil; p = p.down {
		if p.key == 1 && p.value != "k" {
			t.Errorf("p.value = %v, want \"k\"", p.value)
		}
	}

}

func TestInsertWithMockDicisionMaker(t *testing.T) {
	// GIVEN
	head := NewSkipList(1, "a")
	head.SetDicisionMaker(&mockDicisionMaker{false})
	head = head.Insert(2, "b")

	// GIVEN
	head.SetDicisionMaker(&mockDicisionMaker{true})
	head = head.Insert(3, "c")

	// THEN
	if head.key != 1 {
		t.Errorf("head.key = %v, want 1", head.key)
	}
	if head.right.key != 3 {
		t.Errorf("head.right.key = %v, want 3", head.right.key)
	}
	if head.down.key != 1 {
		t.Errorf("head.down.key = %v, want 1", head.down.key)
	}
	if head.down.right.key != 2 {
		t.Errorf("head.down.right.key = %v, want 2", head.down.right.key)
	}

	// GIVEN
	head.SetDicisionMaker(&mockDicisionMaker{false})
	head = head.Insert(4, "d")
	head.SetDicisionMaker(&mockDicisionMaker{true})
	head = head.Insert(5, "e")

	// THEN
	if head.key != 1 {
		t.Errorf("head.key = %v, want 1", head.key)
	}
	if head.right.key != 5 {
		t.Errorf("head.right.key = %v, want 5", head.right.key)
	}
	if head.down.key != 1 {
		t.Errorf("head.down.key = %v, want 1", head.down.key)
	}
	if head.down.right.key != 3 || head.down.right.right.key != 5 {
		t.Errorf("head.down.right.key = %v, want 3", head.down.right.key)
		t.Errorf("head.down.right.right.key = %v, want 5", head.down.right.right.key)
	}
	if head.down.down.key != 1 {
		t.Errorf("head.down.down.key = %v, want 1", head.down.down.key)
	}
	if head.down.down.right.right.right.key != 4 {
		t.Errorf("head.down.down.right.right.right.key = %v, want 4", head.down.down.right.right.right.key)
	}
}

func TestDeleteSkipListNode(t *testing.T) {
	// GIVEN
	head := NewSkipList(1, "a")
	head = head.Insert(2, "b")
	head = head.Insert(3, "c")
	head = head.Insert(4, "d")
	head = head.Insert(5, "e")

	// WHEN
	head.Delete(3)

	// THEN
	if head.Search(3) != nil {
		t.Errorf("head.Search(3) = %v, want nil", head.Search(3))
	}
	if head.down.Search(3) != nil {
		t.Errorf("head.down.Search(3) = %v, want nil", head.down.Search(3))
	}
	if head.right.Search(3) != nil {
		t.Errorf("head.right.Search(3) = %v, want nil", head.right.Search(3))
	}
	if head.right.down.Search(3) != nil {
		t.Errorf("head.right.down.Search(3) = %v, want nil", head.right.down.Search(3))
	}
}

func TestTheMostLowListContainsAllNodes(t *testing.T) {
	// GIVEN
	head := NewSkipList(1, "a")
	head.SetDicisionMaker(&mockDicisionMaker{true})
	head = head.Insert(2, "b")
	head = head.Insert(3, "c")
	head = head.Insert(4, "d")
	head = head.Insert(5, "e")

	// WHEN
	p := head
	for p.down != nil {
		p = p.down
	}
	data := make(map[int]string)
	for p != nil {
		data[p.key] = p.value.(string)
		p = p.right
	}

	// THEN
	if len(data) != 5 {
		t.Errorf("len(data) = %v, want 5", len(data))
	}
	if data[1] != "a" {
		t.Errorf("data[1] = %v, want \"a\"", data[1])
	}
	if data[2] != "b" {
		t.Errorf("data[2] = %v, want \"b\"", data[2])
	}
	if data[3] != "c" {
		t.Errorf("data[3] = %v, want \"c\"", data[3])
	}
	if data[4] != "d" {
		t.Errorf("data[4] = %v, want \"d\"", data[4])
	}
	if data[5] != "e" {
		t.Errorf("data[5] = %v, want \"e\"", data[5])
	}

}
