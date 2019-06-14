package binarysearch

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var arr []int

func init() {
	rand.Seed(time.Now().Unix())

	arr = rand.Perm(20000)[3000:3999]
	sort.Ints(arr)
}

func TestIterativeExisting(t *testing.T) {
	expected := rand.Intn(len(arr) - 1)
	term := arr[expected]

	found := Iterative(arr, term)

	if found != expected {
		t.Fatalf("expected %d, got %d", expected, found)
	}
}

func TestIterativeNotExisting(t *testing.T) {
	found := Iterative(arr, -2)
	if found != -1 {
		t.Fatalf("expected -1, got %d", found)
	}
}

func TestRecursiveExisting(t *testing.T) {
	expected := rand.Intn(len(arr) - 1)
	term := arr[expected]

	found := Recursive(arr, term)

	if found != expected {
		t.Fatalf("expected %d, got %d", expected, found)
	}
}

func TestRecursiveNotExisting(t *testing.T) {
	found := Recursive(arr, -2)
	if found != -1 {
		t.Fatalf("expected -1, got %d", found)
	}
}

func benchmarkSearch(i int, f func(arr []int, term int) int, b *testing.B) {
	arr := rand.Perm(i * 20)[i : i*2-1]
	sort.Ints(arr)

	for n := 0; n < b.N; n++ {
		term := arr[rand.Intn(len(arr)-1)]
		f(arr, term)
	}
}

func BenchmarkIterative10(b *testing.B) {
	benchmarkSearch(10, Iterative, b)
}

func BenchmarkIterative100(b *testing.B) {
	benchmarkSearch(100, Iterative, b)
}

func BenchmarkIterative1000(b *testing.B) {
	benchmarkSearch(1000, Iterative, b)
}

func BenchmarkRecursive10(b *testing.B) {
	benchmarkSearch(10, Recursive, b)
}

func BenchmarkRecursive100(b *testing.B) {
	benchmarkSearch(100, Recursive, b)
}

func BenchmarkRecursive1000(b *testing.B) {
	benchmarkSearch(1000, Recursive, b)
}
