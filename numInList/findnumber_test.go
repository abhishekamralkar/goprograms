package findnumber_test

import (
	"testing"

	findnumber "github.com/abhishekamralkar/numInList"
)

func TestFindNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		search   int
		expected bool
	}{
		{"Number exists in slice", []int{1, 2, 3, 4, 5}, 3, true},
		{"Number does not exist in slice", []int{1, 2, 3, 4, 5}, 6, false},
		{"Empty slice", []int{}, 1, false},
		{"Single element slice - match", []int{1}, 1, true},
		{"Single element slice - no match", []int{2}, 1, false},
		{"Multiple identical elements", []int{3, 3, 3, 3}, 3, true},
		{"Large slice", make([]int, 100000), 99999, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findnumber.FindNumber(tt.input, tt.search)
			if result != tt.expected {
				t.Errorf("FindNumber(%v, %d) = %v; want %v", tt.input, tt.search, result, tt.expected)
			}
		})
	}
}
