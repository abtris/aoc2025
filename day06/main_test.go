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

func TestCalculateProblemPart2(t *testing.T) {
	tests := []struct {
		problem  Problem
		expected int64
	}{
		{Problem{[]int{4, 431, 623}, "+"}, 1058},
		{Problem{[]int{175, 581, 32}, "*"}, 3253600},
		{Problem{[]int{8, 248, 369}, "+"}, 625},
		{Problem{[]int{356, 24, 1}, "*"}, 8544},
	}

	for _, tt := range tests {
		result := calculateProblemPart2(tt.problem)
		if result != tt.expected {
			t.Errorf("calculateProblemPart2(%v) = %d, expected %d", tt.problem, result, tt.expected)
		}
	}
}

func TestSolvePart2WithTestInput(t *testing.T) {
	result, err := solvePart2("input_test")
	if err != nil {
		t.Fatalf("Error solving part 2: %v", err)
	}

	expected := int64(3263827) // 1058 + 3253600 + 625 + 8544
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
