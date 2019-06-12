package main

import "math"

// BinarySearch is a recursive algorithm for finding x in arr
func BinarySearch(arr []int64, l, r, x int64) int64 {
	if r >= l {
		mid := int64(math.Ceil(float64(l + (r-l)/2)))

		// mid == x return mid
		if arr[mid] == x {
			return mid
		}

		// x is smaller than mid, focus on the left of arr
		if arr[mid] > x {
			return BinarySearch(arr, l, mid-l, x)
		}

		// x is larger than mid, focus on the right of arr
		return BinarySearch(arr, mid+l, r, x)
	}

	// x was not found
	return -1
}
