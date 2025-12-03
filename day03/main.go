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

// findMaxJoltagePart2 finds the maximum 12-digit joltage by selecting 12 batteries
// Strategy: greedily select the largest digits while maintaining order
func findMaxJoltagePart2(bank string) int64 {
	n := len(bank)
	k := 12 // number of batteries to select

	if n < k {
		return 0
	}

	// Greedy approach: for each position in the result, pick the largest digit
	// from the remaining candidates that still leaves enough digits for the rest
	result := make([]byte, 0, k)
	startPos := 0

	for i := 0; i < k; i++ {
		// We need to select k-i more digits
		// We can look at positions from startPos to n-(k-i)
		maxDigit := byte('0')
		maxPos := startPos

		endPos := n - (k - i) + 1
		for j := startPos; j < endPos; j++ {
			if bank[j] > maxDigit {
				maxDigit = bank[j]
				maxPos = j
			}
		}

		result = append(result, maxDigit)
		startPos = maxPos + 1
	}

	// Convert result to int64
	var joltage int64 = 0
	for _, digit := range result {
		joltage = joltage*10 + int64(digit-'0')
	}

	return joltage
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

func solvePart2(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var totalJoltage int64 = 0
	bankCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		bankCount++
		maxJoltage := findMaxJoltagePart2(line)
		totalJoltage += maxJoltage
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	fmt.Printf("Processed %d banks (Part 2)\n", bankCount)
	return totalJoltage, nil
}

func main() {
	// Part 1
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Total output joltage: %d\n", result)

	// Part 2
	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Total output joltage: %d\n", result2)
}
