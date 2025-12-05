package main

import "testing"

func TestSolveWithTestInput(t *testing.T) {
	result, err := solve("input_test")
	if err != nil {
		t.Fatalf("Error solving: %v", err)
	}

	expected := 13
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2WithTestInput(t *testing.T) {
	result, err := solvePart2("input_test")
	if err != nil {
		t.Fatalf("Error solving part 2: %v", err)
	}

	expected := 43 // 13 + 12 + 7 + 5 + 2 + 1 + 1 + 1 + 1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
