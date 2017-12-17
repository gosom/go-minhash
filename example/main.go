package main

import (
	"bufio"
	"fmt"
	"github.com/gosom/go-minhash"
	"os"
	"time"
)

func main() {
	f, _ := os.Open("plato1.txt")
	scanner := bufio.NewScanner(f)

	var tokens1 []string
	for scanner.Scan() {
		line := scanner.Text()
		tokens1 = append(tokens1, line)
	}

	f, _ = os.Open("plato2.txt")
	scanner = bufio.NewScanner(f)
	var tokens2 []string
	for scanner.Scan() {
		line := scanner.Text()
		tokens2 = append(tokens2, line)
	}
	start := time.Now()
	perms := minhash.NewPermutations(64, int64(0))
	m1 := minhash.NewMinhash(perms)
	for _, t := range tokens1 {
		m1.Update([]byte(t))
	}
	m2 := minhash.NewMinhash(perms)
	for _, t := range tokens2 {
		m2.Update([]byte(t))
	}
	similarity := m2.Jaccard(m1)

	elapsed := time.Since(start)
	fmt.Println("Similar: %f and Took %s", similarity, elapsed)
}
