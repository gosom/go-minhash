package minhash

import (
	"testing"
)

func TestNewPermutations(t *testing.T) {
	size := 64
	seed := int64(0)
	p := NewPermutations(size, seed)
	if p.size != size {
		t.Fatal("Size expected %d but got %d", size, p.size)
	}
	if p.seed != seed {
		t.Fatal("Seed expected %d but got %d", seed, p.seed)
	}
	for _, value := range p.values {
		if value.a < 1 || value.a > mersennePrime {
			t.Fatal("Random %d out of bounds", value.a)
		}
		if value.b < 0 || value.b > mersennePrime {
			t.Fatal("Random %d out of bounds", value.b)
		}
	}
}

func TestNewMinhash(t *testing.T) {

	perms := NewPermutations(64, int64(0))
	m := NewMinhash(perms)
	if len(m.hashvalues) != perms.size {
		t.Fatal("Hashvalues expected size %d but got %d", 64, len(m.hashvalues))
	}
	if len(m.permutations.values) != perms.size {
		t.Fatal("Permutations expected size %d but got %d", 64, len(m.permutations.values))
	}

	for _, value := range m.hashvalues {
		if value != infinity {
			t.Fatal("Expected infinity but got %d", value)
		}
	}

}

func TestUpdate(t *testing.T) {
	perms := NewPermutations(64, int64(0))
	m := NewMinhash(perms)

	s := "gosom"
	m.Update([]byte(s))
}

func TestJaccardSame(t *testing.T) {
	perms := NewPermutations(64, int64(0))
	m1 := NewMinhash(perms)
	m2 := NewMinhash(perms)
	s1 := []string{"gosom", "Giorgos", "Komninos", "g+test@example.com"}
	for _, s := range s1 {
		m1.Update([]byte(s))
	}
	s2 := []string{"gosom", "Giorgos", "Komninos", "g+test@example.com"}
	for _, s := range s2 {
		m2.Update([]byte(s))
	}

	ans := m1.Jaccard(m2)
	if ans != 1 {
		t.Fatal("We should get similarity of 1")
	}
}

func TestJaccardDifferent(t *testing.T) {
	perms := NewPermutations(64, int64(0))
	m1 := NewMinhash(perms)
	m2 := NewMinhash(perms)
	s1 := []string{"a", "b", "c", "d"}
	for _, s := range s1 {
		m1.Update([]byte(s))
	}
	s2 := []string{"e", "f", "g", "f"}
	for _, s := range s2 {
		m2.Update([]byte(s))
	}

	ans := m1.Jaccard(m2)
	if ans != 0 {
		t.Fatal("We should get similarity of 0 but got %f", ans)
	}
}

func TestJaccardHalfEqual(t *testing.T) {
	perms := NewPermutations(60, int64(0))
	m1 := NewMinhash(perms)
	m2 := NewMinhash(perms)
	s1 := []string{"a", "b", "c", "d"}
	for _, s := range s1 {
		m1.Update([]byte(s))
	}
	s2 := []string{"e", "f", "a", "b"}
	for _, s := range s2 {
		m2.Update([]byte(s))
	}

	ans := m1.Jaccard(m2)

	if ans <= 0.3 {
		t.Fatal("We should get similarity of at least 0.3 but got %f", ans)
	}

}

func Testrandom(t *testing.T) {
	value := random(uint64(0), uint64(10))
	if value > 10 {
		t.Fatal("Expected less that 10")
	}
	value = random(uint64(0), uint64(0))
	if value != 0 {
		t.Fatal("Expected zero")
	}
	value = random(uint64(100), uint64(100))
	if value != 0 {
		t.Fatal("Expected hundred")
	}
}

func BenchmarkNew(b *testing.B) {
	perms := NewPermutations(64, int64(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewMinhash(perms)
	}
}

func BenchmarkUpdate(b *testing.B) {
	s1 := []string{"gosom", "Giorgos", "Komninos", "g+test@gkomninos.com"}
	perms := NewPermutations(64, int64(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m1 := NewMinhash(perms)
		for _, s := range s1 {
			m1.Update([]byte(s))
		}
	}
}
