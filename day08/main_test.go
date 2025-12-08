package main

import "testing"

func TestSolveWithTestInput(t *testing.T) {
	// After making 10 connections, the answer should be 40
	result, err := solve("input_test", 10)
	if err != nil {
		t.Fatalf("Error solving: %v", err)
	}

	expected := 40
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2WithTestInput(t *testing.T) {
	// The last connection is between 216,146,977 and 117,168,530
	// 216 * 117 = 25272
	result, err := solvePart2("input_test")
	if err != nil {
		t.Fatalf("Error solving part 2: %v", err)
	}

	expected := 25272
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
