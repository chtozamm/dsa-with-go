package selection_sort

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Run("sort slice of integers", func(t *testing.T) {
		nums := []int{9, 2, 11, 3, 4, 8, 7}

		want := []int{2, 3, 4, 7, 8, 9, 11}
		got := SelectionSort(nums)

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("shouldn't modify original slice", func(t *testing.T) {
		vals := []int{9, 2, 11, 3, 4, 8, 7}

		_ = SelectionSort(vals)

		initialVals := []int{9, 2, 11, 3, 4, 8, 7}

		if !slices.Equal(vals, initialVals) {
			t.Errorf("Original slice shouldn't be modified: got %v, want %v", vals, initialVals)
		}
	})
}
