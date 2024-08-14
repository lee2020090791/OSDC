package main

import (
	"math/rand"
	"testing"
)

var (
	N    = 10000
	seed = int64(10)
	diff = 2
)

func BenchmarkLinearSearch(b *testing.B) {
	rand.Seed(seed)
	var set LinearSlice
	for i := 0; i < N; i++ {
		set.Add(i)
	}
	for i := 0; i < b.N; i++ {
		set.Contains(rand.Int() % (N * diff))
	}
}
func BenchmarkBinarySearch(b *testing.B) {
	rand.Seed(seed)
	var set SortedInts
	for i := 0; i < N; i++ {
		set.Add(i)
	}
	for i := 0; i < b.N; i++ {
		set.Contains(rand.Int() % (N * diff))
	}
}
func BenchmarkFindKeyInMap(b *testing.B) {
	rand.Seed(seed)
	var set IntHashSet = make(map[int]struct{}, 0)
	for i := 0; i < N; i++ {
		set.Add(i)
	}
	for i := 0; i < b.N; i++ {
		set.Contains(rand.Int() % (N * diff))
	}
}
