package linear

import "testing"

// TestCreateBloomFilter tests the creation of a Bloom Filter.
func TestCreateBloomFilter(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)
	if bf == nil {
		t.Error("failed to create a Bloom Filter")
	}
}

// TestAddAndContains tests adding elements to a Bloom Filter.
func TestAddAndContains(t *testing.T) {
	// GIVEN
	bf := NewBloomFilter(1000, 0.01)

	// WHEN
	bf.Add([]byte("Hello"))
	bf.Add([]byte("World"))

	// THEN
	if !bf.Contains([]byte("Hello")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("World")) {
		t.Error("failed to add element to Bloom Filter")
	}
}

// TestMoreAddAndContains tests adding more elements to a Bloom Filter.
func TestMoreAddAndContains(t *testing.T) {
	// GIVEN
	bf := NewBloomFilter(1000, 0.01)

	// WHEN
	bf.Add([]byte("Hello"))
	bf.Add([]byte("World"))
	bf.Add([]byte("Golang"))
	bf.Add([]byte("is"))
	bf.Add([]byte("awesome"))
	bf.Add([]byte("and"))
	bf.Add([]byte("fast"))
	bf.Add([]byte("and2"))
	bf.Add([]byte("easy"))
	bf.Add([]byte("to"))
	bf.Add([]byte("learn"))
	bf.Add([]byte("and3"))
	bf.Add([]byte("concurrency"))
	bf.Add([]byte("is"))
	bf.Add([]byte("built"))
	bf.Add([]byte("in"))
	bf.Add([]byte("and4"))
	bf.Add([]byte("it"))
	bf.Add([]byte("is5"))
	bf.Add([]byte("fun"))
	bf.Add([]byte("to"))
	bf.Add([]byte("use"))
	bf.Add([]byte("and6"))
	bf.Add([]byte("it"))
	bf.Add([]byte("is7"))
	bf.Add([]byte("is8"))
	bf.Add([]byte("simple"))
	bf.Add([]byte("to"))
	bf.Add([]byte("compile"))
	bf.Add([]byte("and9"))
	bf.Add([]byte("run"))
	bf.Add([]byte("and10"))
	bf.Add([]byte("it"))
	bf.Add([]byte("is11"))
	bf.Add([]byte("open"))
	bf.Add([]byte("source"))

	// THEN
	if !bf.Contains([]byte("Hello")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("World")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("Golang")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("is")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("awesome")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("and")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("fast")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("open")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if !bf.Contains([]byte("source")) {
		t.Error("failed to add element to Bloom Filter")
	}
	if bf.Contains([]byte("Golang2")) {
		t.Error("false positive result")
	}
	if bf.Contains([]byte("walk")) {
		t.Error("false positive result")
	}
}

// TestFalsePositive tests the false positive probability of a Bloom Filter.
func TestFalsePositive(t *testing.T) {
	// GIVEN
	bf := NewBloomFilter(1000, 0.01)

	// WHEN
	bf.Add([]byte("Hello"))
	bf.Add([]byte("World"))

	// THEN
	if bf.Contains([]byte("Golang")) {
		t.Error("false positive result")
	}
}
