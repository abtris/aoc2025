package main

import "testing"

func TestSolveWithTestInput(t *testing.T) {
	result, err := solve("input_test")
	if err != nil {
		t.Fatalf("Error solving: %v", err)
	}

	expected := 3
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
