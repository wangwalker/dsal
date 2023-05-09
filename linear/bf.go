package linear

import (
	"hash/fnv"
	"math"
)

// BloomFilter is a probabilistic data structure to check whether
// an element exists in a set. It may return a false positive result,
// but not a false negative.
type BloomFilter struct {
	bitmap []bool
	k      uint64 // number of hash functions
}

// NewBloomFilter creates a new Bloom Filter with given capacity and
// false positive probability.
func NewBloomFilter(capacity int, fpProbability float64) *BloomFilter {
	bf := &BloomFilter{}
	bf.bitmap = make([]bool, capacity)
	bf.k = uint64(math.Ceil(-math.Log(fpProbability) / math.Log(2)))
	return bf
}

// Add adds a new element to the Bloom Filter.
func (bf *BloomFilter) Add(data []byte) {
	for i := uint64(0); i < bf.k; i++ {
		h := fnv.New64a()
		h.Write(data)
		index := h.Sum64() % uint64(len(bf.bitmap))
		bf.bitmap[index] = true
	}
}

// Contains checks whether a given element exists in the Bloom Filter.
// It may return a false positive result, but not a false negative.
func (bf *BloomFilter) Contains(data []byte) bool {
	for i := uint64(0); i < bf.k; i++ {
		h := fnv.New64a()
		h.Write(data)
		index := h.Sum64() % uint64(len(bf.bitmap))
		if !bf.bitmap[index] {
			return false
		}
	}
	return true
}
