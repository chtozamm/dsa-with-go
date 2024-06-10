package binary_search

// BinarySearch searches for needle in a sorted slice of ints and returns its index.
// Returns -1 if needle is not present in haystack. The slice must be sorted in ascending order.
func BinarySearch(haystack []int, needle int) int {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := (low + high) / 2
		if haystack[mid] == needle {
			return mid
		}
		if haystack[mid] < needle {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
