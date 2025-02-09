package pratices

import "testing"

//Use table-driven tests for readability.

// Test function (must start with "Test")
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{10, -5, 5},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

// Function to be tested
func Add(a, b int) int {
	return a + b
}
