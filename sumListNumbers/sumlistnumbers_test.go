package sumlistnumbers_test

import (
	"testing"

	sumlistnumbers "github.com/abhishekamralkar/sumListNumbers"
)

type testCase struct {
	name  string
	input []int
	total int
}

func TestSumList(t *testing.T) {
	t.Parallel()

	testcases := []testCase{
		{"Total1", []int{1, 2, 3, 4, 5, 6}, 21},
		{"Total2", []int{1}, 1},
		{"Total3", []int{}, 0},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := sumlistnumbers.SumList(tc.input)
			if result != tc.total {
				t.Errorf("SumList(%v) = %d; want %d", tc.input, result, tc.total)
			}
		})
	}

}
