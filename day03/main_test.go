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

func TestFindMaxJoltagePart2(t *testing.T) {
	tests := []struct {
		bank     string
		expected int64
	}{
		{"987654321111111", 987654321111}, // everything except some 1s at the end
		{"811111111111119", 811111111119}, // everything except some 1s
		{"234234234234278", 434234234278}, // skip 2,3,2 near start
		{"818181911112111", 888911112111}, // skip some 1s near front
	}

	for _, tt := range tests {
		result := findMaxJoltagePart2(tt.bank)
		if result != tt.expected {
			t.Errorf("findMaxJoltagePart2(%s) = %d, expected %d", tt.bank, result, tt.expected)
		}
	}
}

func TestSolvePart2WithTestInput(t *testing.T) {
	result, err := solvePart2("input_test")
	if err != nil {
		t.Fatalf("Error solving part 2: %v", err)
	}

	expected := int64(3121910778619) // 987654321111 + 811111111119 + 434234234278 + 888911112111
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
