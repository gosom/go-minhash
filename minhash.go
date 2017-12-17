package minhash

import (
	"hash/fnv"
	"math"
	"math/rand"
)

const (
	infinity      uint64 = math.MaxUint64
	mersennePrime        = uint64((1 << 61) - 1)
)

func random(min uint64, max uint64) uint64 {
	return uint64(rand.Int63n(int64(max-min+1))) + min
}

type permutation struct {
	a uint64
	b uint64
}

type permutations struct {
	size   int
	seed   int64
	values []permutation
}

type minhash struct {
	permutations *permutations
	hashvalues   []uint64
}

func NewPermutations(size int, seed int64) *permutations {
	p := &permutations{}
	p.size = size
	p.seed = seed
	p.values = make([]permutation, size)
	rand.Seed(seed)
	for i := range p.values {
		p.values[i] = permutation{random(uint64(1), mersennePrime),
			random(uint64(0), mersennePrime)}
	}
	return p
}

func NewMinhash(permutations *permutations) *minhash {
	m := &minhash{}
	m.permutations = permutations
	m.initHashvalues()
	return m
}

func (m *minhash) Hashvalues() []uint64 {
	return m.hashvalues
}

func (m *minhash) Update(b []byte) {
	hasher := fnv.New32()
	hasher.Write(b)
	val := uint64(hasher.Sum32())
	for i, hv := range m.hashvalues {
		hi := (m.permutations.values[i].a*val + m.permutations.values[i].b) % mersennePrime
		if hi > 0 && hi < hv {
			m.hashvalues[i] = hi
		}
	}

}

func (m *minhash) Jaccard(other *minhash) float64 {
	if m.permutations.size != other.permutations.size {
		panic("Size mismatch.")
	}
	common := 0
	for i := range m.hashvalues {
		if m.hashvalues[i] == other.hashvalues[i] {
			common++
		}
	}
	return float64(common) / float64(m.permutations.size)
}

func (m *minhash) initHashvalues() {
	m.hashvalues = make([]uint64, m.permutations.size)
	for i := range m.hashvalues {
		m.hashvalues[i] = infinity
	}
}
