package pratices

import "testing"

//Use table-driven tests for readability.

func BenchmarkAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{10, -5, 5},
	}

	for _, tt := range tests {
		result := add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func add(a, b int) int {
	return a + b
}
