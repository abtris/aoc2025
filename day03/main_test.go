package main

import "testing"

func TestFindMaxJoltage(t *testing.T) {
	tests := []struct {
		bank     string
		expected int
	}{
		{"987654321111111", 98}, // 9 and 8
		{"811111111111119", 89}, // 8 and 9
		{"234234234234278", 78}, // 7 and 8
		{"818181911112111", 92}, // 9 and 2
		{"12345", 45},           // 4 and 5
	}

	for _, tt := range tests {
		result := findMaxJoltage(tt.bank)
		if result != tt.expected {
			t.Errorf("findMaxJoltage(%s) = %d, expected %d", tt.bank, result, tt.expected)
		}
	}
}

func TestSolveWithTestInput(t *testing.T) {
	result, err := solve("input_test")
	if err != nil {
		t.Fatalf("Error solving: %v", err)
	}

	expected := 357 // 98 + 89 + 78 + 92
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
