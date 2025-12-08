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
