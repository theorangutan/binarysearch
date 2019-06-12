package main

// binarySearch is a recursive algorithm for finding x in arr
func iterativeBinarySearch(arr []int, search int) int {
	for l, r := 0, len(arr)-1; l <= r; {
		mid := (l + r) / 2
		if arr[mid] < search {
			l = mid + 1
		} else if arr[mid] > search {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func recursiveBinarySearch(arr []int, search int) int {
	return 0
}
