package main

import (
	"bufio"
	"fmt"
	"os"
)

// findMaxJoltage finds the maximum joltage for a bank by trying all pairs
func findMaxJoltage(bank string) int {
	maxJoltage := 0

	// Try all pairs of batteries (i, j) where i < j
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			// Form a two-digit number from batteries at positions i and j
			digit1 := int(bank[i] - '0')
			digit2 := int(bank[j] - '0')
			joltage := digit1*10 + digit2

			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
	}

	return maxJoltage
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	totalJoltage := 0
	bankCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		bankCount++
		maxJoltage := findMaxJoltage(line)
		totalJoltage += maxJoltage
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	fmt.Printf("Processed %d banks\n", bankCount)
	return totalJoltage, nil
}

func main() {
	// First verify with test input
	testResult, err := solve("input_test")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (test): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Test output joltage: %d (expected 357)\n", testResult)

	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Total output joltage: %d\n", result)
}
