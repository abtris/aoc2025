package main

import "testing"

func TestIsInvalidID(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{11, true},         // "11" = "1" + "1"
		{22, true},         // "22" = "2" + "2"
		{55, true},         // "55" = "5" + "5"
		{99, true},         // "99" = "9" + "9"
		{6464, true},       // "6464" = "64" + "64"
		{123123, true},     // "123123" = "123" + "123"
		{1010, true},       // "1010" = "10" + "10"
		{222222, true},     // "222222" = "222" + "222"
		{446446, true},     // "446446" = "446" + "446"
		{1188511885, true}, // "1188511885" = "11885" + "11885"
		{38593859, true},   // "38593859" = "3859" + "3859"
		{101, false},       // not repeated
		{12, false},        // not repeated
		{100, false},       // odd length
		{1698522, false},   // not repeated
	}

	for _, tt := range tests {
		result := isInvalidID(tt.id)
		if result != tt.expected {
			t.Errorf("isInvalidID(%d) = %v, expected %v", tt.id, result, tt.expected)
		}
	}
}

func TestSolveWithTestInput(t *testing.T) {
	result, err := solve("input_test")
	if err != nil {
		t.Fatalf("Error solving: %v", err)
	}

	expected := int64(1227775554)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestIsInvalidIDPart2(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{11, true},         // "11" = "1" twice
		{22, true},         // "22" = "2" twice
		{99, true},         // "99" = "9" twice
		{111, true},        // "111" = "1" three times
		{999, true},        // "999" = "9" three times
		{1010, true},       // "1010" = "10" twice
		{222222, true},     // "222222" = "222" twice
		{446446, true},     // "446446" = "446" twice
		{1188511885, true}, // "1188511885" = "11885" twice
		{38593859, true},   // "38593859" = "3859" twice
		{565656, true},     // "565656" = "56" three times
		{824824824, true},  // "824824824" = "824" three times
		{2121212121, true}, // "2121212121" = "21" five times
		{12341234, true},   // "12341234" = "1234" twice
		{123123123, true},  // "123123123" = "123" three times
		{1212121212, true}, // "1212121212" = "12" five times
		{1111111, true},    // "1111111" = "1" seven times
		{101, false},       // not repeated
		{12, false},        // not repeated
		{100, false},       // not repeated
		{1698522, false},   // not repeated
	}

	for _, tt := range tests {
		result := isInvalidIDPart2(tt.id)
		if result != tt.expected {
			t.Errorf("isInvalidIDPart2(%d) = %v, expected %v", tt.id, result, tt.expected)
		}
	}
}

func TestSolvePart2WithTestInput(t *testing.T) {
	result, err := solvePart2("input_test")
	if err != nil {
		t.Fatalf("Error solving part 2: %v", err)
	}

	expected := int64(4174379265)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
