package calculator

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},    // Test Case 1: 2 + 3 = 5
		{-5, 10, 5},  // Test Case 2: -5 + 10 = 5
		{0, 0, 0},    // Test Case 3: 0 + 0 = 0
		{100, -100, 0}, // Test Case 4: 100 + (-100) = 0
		// TODO: Add another test
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d, expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
