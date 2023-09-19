package helpers

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		inputArray []int
		target     int
		expected   int
	}{
		{inputArray: []int{1, 2, 3, 4, 5}, target: 10, expected: -1},
		{inputArray: []int{10, 11, 12, 13, 15}, target: 5, expected: -1},
		{inputArray: []int{1, 2, 3, 4, 5}, target: 3, expected: 2},
		{inputArray: []int{1, 2, 3, 4, 5}, target: 4, expected: 3},
		{inputArray: []int{1, 2, 3, 4, 5}, target: 5, expected: 4},
		{inputArray: []int{1, 2, 4, 7, 8, 12, 15, 19, 24, 50, 69, 80, 100}, target: 100, expected: 12},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing array %v, target: %d, expected: %d\n", test.inputArray, test.target, test.expected), func(t *testing.T) {

			result := BinarySearch(test.inputArray, test.target)

			if result != test.expected {
				t.Errorf("unexpected result for non recursive version -> got: %d, expected%d\n", result, test.expected)
			}

			result = BinarySearch(test.inputArray, test.target)

			if result != test.expected {
				t.Errorf("unexpected result for recursive version -> got: %d, expected%d\n", result, test.expected)
			}
		})
	}
}
