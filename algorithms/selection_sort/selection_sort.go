package main

// SelectionSort makes a copy of the provided slice and returns it's sorted version.
func SelectionSort(nums []int) []int {
	var sortedNums = make([]int, len(nums))
	copy(sortedNums, nums)

	for i := 0; i < len(sortedNums); i++ {
		for j := i + 1; j < len(sortedNums); j++ {
			if sortedNums[j] < sortedNums[i] {
				temp := sortedNums[j]
				sortedNums[j] = sortedNums[i]
				sortedNums[i] = temp
			}
		}
	}

	return sortedNums
}
