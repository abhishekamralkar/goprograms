package reverseString_test

import (
	"testing"

	"github.com/abhishekamralkar/reverseString"
)

type testCase struct {
	name     string
	input    string
	expected string
}

func TestReverseIt(t *testing.T) {
	t.Parallel()

	testcases := []testCase{
		{"test1", "abhishek", "kehsihba"},
		{"test2", "cat", "tac"},
		{"test3", "one", "eno"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverseString.ReverseIt(tc.input)
			if result != tc.expected {
				t.Errorf("ReverseIt(%s): want %s, got %s", tc.input, tc.expected, result)
			}
		})
	}

}
