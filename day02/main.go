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

func main() {
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Sum of invalid IDs: %d\n", result)
}
