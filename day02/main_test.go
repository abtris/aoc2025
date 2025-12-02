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
