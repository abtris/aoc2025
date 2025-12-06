package main

import "testing"

func TestCalculateProblem(t *testing.T) {
	tests := []struct {
		problem  Problem
		expected int
	}{
		{Problem{[]int{123, 45, 6}, "*"}, 33210},
		{Problem{[]int{328, 64, 98}, "+"}, 490},
		{Problem{[]int{51, 387, 215}, "*"}, 4243455},
		{Problem{[]int{64, 23, 314}, "+"}, 401},
	}

	for _, tt := range tests {
		result := calculateProblem(tt.problem)
		if result != tt.expected {
			t.Errorf("calculateProblem(%v) = %d, expected %d", tt.problem, result, tt.expected)
		}
	}
}

func TestSolveWithTestInput(t *testing.T) {
	result, err := solve("input_test")
	if err != nil {
		t.Fatalf("Error solving: %v", err)
	}

	expected := 4277556 // 33210 + 490 + 4243455 + 401
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
