package quicksort

func QuickSort(nums []int) (sortedNums []int) {

	// Base case
	if len(nums) < 2 {
		return nums
	}

	// Pivot could be set to any value in the given slice.
	mid := len(nums) / 2
	pivot := nums[mid]

	var left []int
	var right []int

	for _, num := range nums {
		if num == pivot {
			continue
		}
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	sortedNums = append(sortedNums, QuickSort(left)...)
	sortedNums = append(sortedNums, pivot)
	sortedNums = append(sortedNums, QuickSort(right)...)

	return
}
