package main

// iterativeBinarySearch is an iterative algorithm for finding x in arr
// return is the index where x is located, or -1 (not found)
func iterativeBinarySearch(arr []int, l, r, search int) int {
	for l <= r {
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

// recursiveBinarySearch is a recursive algorithm for finding x in arr
// return is the index where x is located, or -1 (not found)
func recursiveBinarySearch(arr []int, l, r, search int) int {
	if r >= l {
		mid := l + (r-l)/2
		if arr[mid] == search {
			return mid
		}
		if arr[mid] > search {
			return recursiveBinarySearch(arr, l, mid-1, search)
		}
		return recursiveBinarySearch(arr, mid+1, r, search)
	}
	return -1
}
