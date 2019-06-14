package binarysearch

// Iterative is an iterative algorithm for finding term in arr
// return is the index where term is located, or -1 (not found)
func Iterative(arr []int, term int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < term {
			l = mid + 1
		} else if arr[mid] > term {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// Recursive is the wrapper function for the recursive function below
func Recursive(arr []int, term int) int {
	return recursive(arr, 0, len(arr)-1, term)
}

// recursive is a recursive algorithm for finding term in arr
// return is the index where term is located, or -1 (not found)
func recursive(arr []int, l, r, search int) int {
	if r >= l {
		mid := l + (r-l)/2
		if arr[mid] == search {
			return mid
		}
		if arr[mid] > search {
			return recursive(arr, l, mid-1, search)
		}
		return recursive(arr, mid+1, r, search)
	}
	return -1
}
