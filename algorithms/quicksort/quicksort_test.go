package quicksort

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {

	cases := []struct {
		Name     string
		Nums     []int
		Expected []int
	}{
		{
			"positive integers",
			[]int{5, 0, 3, 17, 1, 6},
			[]int{0, 1, 3, 5, 6, 17},
		},
		{
			"negative values",
			[]int{-1, -2, 0, -14, 7},
			[]int{-14, -2, -1, 0, 7},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := QuickSort(test.Nums)
			if !slices.Equal(got, test.Expected) {
				t.Errorf("got %d, expected %d", got, test.Expected)
			}
		})
	}
}
