package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		Name     string
		Haystack []int
		Needle   int
		Expected int
	}{
		{
			"positive integers",
			[]int{1, 2, 3, 4, 5, 6},
			5,
			4,
		},
		{
			"with negative integers",
			[]int{-5, -2, 0, 4, 7},
			-2,
			1,
		},
		{
			"returns -1 if value not in slice",
			[]int{1, 2, 3, 4, 5, 6},
			7,
			-1,
		},
		{
			"needle is the first value",
			[]int{1, 2, 3, 4, 5, 6},
			1,
			0,
		},
		{
			"needle is the last value",
			[]int{1, 2, 3, 4, 5, 6},
			6,
			5,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := BinarySearch(test.Haystack, test.Needle)
			if got != test.Expected {
				t.Errorf("got %d, expected %d", got, test.Expected)
			}
		})
	}
}
