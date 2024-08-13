package factorial_test

import (
	"fmt"
	"testing"

	"github.com/abhishekamralkar/factorial"
)

func TestFac(t *testing.T) {
	t.Parallel()

	type testCase struct {
		n    int
		want int
	}

	testcases := []testCase{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, tc := range testcases {
		got := factorial.Fac(tc.n)

		if got != tc.want {
			fmt.Printf("Want %d, got %d", tc.want, got)
		}
	}
}
