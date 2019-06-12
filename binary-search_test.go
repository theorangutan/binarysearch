package main

import "testing"

func TestIterativeBinarySearchExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}

	for idx, val := range arr {
		if tmp := iterativeBinarySearch(arr, val); tmp != idx {
			t.Errorf("expected %d, got %d", idx, tmp)
		}
	}
}

func TestIterativeBinarySearchNotExisting(t *testing.T) {
	arr := []int{0, 1, 2, 3}

	if tmp := iterativeBinarySearch(arr, 5); tmp != -1 {
		t.Errorf("expected -1, got %d", tmp)
	}
}

func BenchmarkIterativeBinarySearch(b *testing.B) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	for n := 0; n < b.N; n++ {
		iterativeBinarySearch(arr, 1)
	}
}
