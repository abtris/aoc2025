package main

import "testing"

func TestIsFresh(t *testing.T) {
	ranges := []Range{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}

	tests := []struct {
		id       int
		expected bool
	}{
		{1, false},  // spoiled
		{5, true},   // in range 3-5
		{8, false},  // spoiled
		{11, true},  // in range 10-14
		{17, true},  // in range 16-20 and 12-18
		{32, false}, // spoiled
	}

	for _, tt := range tests {
		result := isFresh(tt.id, ranges)
		if result != tt.expected {
			t.Errorf("isFresh(%d) = %v, expected %v", tt.id, result, tt.expected)
		}
	}
}

func TestSolveWithTestInput(t *testing.T) {
	result, err := solve("input_test")
	if err != nil {
		t.Fatalf("Error solving: %v", err)
	}

	expected := 3 // IDs 5, 11, 17 are fresh
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
