package binarysearch

import "testing"

func TestIterativeExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}

	for idx, val := range arr {
		if tmp := Iterative(arr, val); tmp != idx {
			t.Errorf("expected %d, got %d", idx, tmp)
		}
	}
}

func TestIterativeNotExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3}

	if tmp := Iterative(arr, 5); tmp != -1 {
		t.Errorf("expected -1, got %d", tmp)
	}
}

func TestRecursiveExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}

	for idx, val := range arr {
		if tmp := Recursive(arr, val); tmp != idx {
			t.Errorf("expected %d, got %d", idx, tmp)
		}
	}
}

func TestRecursiveNotExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3}

	if tmp := Recursive(arr, 5); tmp != -1 {
		t.Errorf("expected -1, got %d", tmp)
	}
}

func BenchmarkIterative(b *testing.B) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	for n := 0; n < b.N; n++ {
		Iterative(arr, 1)
	}
}

func BenchmarkRecursive(b *testing.B) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	for n := 0; n < b.N; n++ {
		Recursive(arr, 1)
	}
}
