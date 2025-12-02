package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isInvalidID checks if a number is made of a sequence repeated twice
func isInvalidID(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	// Must have even length to be repeated twice
	if length%2 != 0 {
		return false
	}

	// Check if first half equals second half
	mid := length / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf == secondHalf
}

// isInvalidIDPart2 checks if a number is made of a sequence repeated at least twice
func isInvalidIDPart2(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	// Try all possible pattern lengths from 1 to length/2
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		// Check if the string length is divisible by pattern length
		if length%patternLen != 0 {
			continue
		}

		// Check if the entire string is made of this pattern repeated
		pattern := s[:patternLen]
		valid := true

		for i := patternLen; i < length; i += patternLen {
			if s[i:i+patternLen] != pattern {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}

func solve(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, fmt.Errorf("empty file")
	}

	line := scanner.Text()
	ranges := strings.Split(line, ",")

	var sum int64

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid range: %s", r)
		}

		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid start: %s", parts[0])
		}

		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid end: %s", parts[1])
		}

		// Check each number in the range
		for i := start; i <= end; i++ {
			if isInvalidID(int(i)) {
				sum += i
			}
		}
	}

	return sum, nil
}

func solvePart2(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, fmt.Errorf("empty file")
	}

	line := scanner.Text()
	ranges := strings.Split(line, ",")

	var sum int64

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid range: %s", r)
		}

		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid start: %s", parts[0])
		}

		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid end: %s", parts[1])
		}

		// Check each number in the range
		for i := start; i <= end; i++ {
			if isInvalidIDPart2(int(i)) {
				sum += i
			}
		}
	}

	return sum, nil
}

func main() {
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Sum of invalid IDs: %d\n", result)

	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Sum of invalid IDs: %d\n", result2)
}
