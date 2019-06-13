package main

import "testing"

func TestIterativeBinarySearchExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}

	for idx, val := range arr {
		if tmp := iterativeBinarySearch(arr, 0, 5, val); tmp != idx {
			t.Errorf("expected %d, got %d", idx, tmp)
		}
	}
}

func TestIterativeBinarySearchNotExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3}

	if tmp := iterativeBinarySearch(arr, 0, 3, 5); tmp != -1 {
		t.Errorf("expected -1, got %d", tmp)
	}
}

func TestRecursiveBinarySearchExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}

	for idx, val := range arr {
		if tmp := recursiveBinarySearch(arr, 0, 5, val); tmp != idx {
			t.Errorf("expected %d, got %d", idx, tmp)
		}
	}
}

func TestRecursiveBinarySearchNotExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3}

	if tmp := recursiveBinarySearch(arr, 0, 3, 5); tmp != -1 {
		t.Errorf("expected -1, got %d", tmp)
	}
}

func BenchmarkIterativeBinarySearch(b *testing.B) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	for n := 0; n < b.N; n++ {
		iterativeBinarySearch(arr, 0, 6, 1)
	}
}

func BenchmarkRecursiveBinarySearch(b *testing.B) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	for n := 0; n < b.N; n++ {
		recursiveBinarySearch(arr, 0, 6, 1)
	}
}
