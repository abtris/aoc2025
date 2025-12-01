package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part 1: Count only when dial ends at 0 after a rotation
func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	position := 50 // Starting position
	count := 0     // Count of times dial points at 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// Parse direction and distance
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, fmt.Errorf("invalid rotation: %s", line)
		}

		// Apply rotation
		if direction == 'L' {
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		} else if direction == 'R' {
			position = (position + distance) % 100
		}

		// Check if dial points at 0
		if position == 0 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

// Part 2: Count every time dial passes through 0 during rotation
func solvePart2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	position := 50 // Starting position
	count := 0     // Count of times dial points at 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// Parse direction and distance
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, fmt.Errorf("invalid rotation: %s", line)
		}

		// Count how many times we pass through 0 during rotation
		if direction == 'L' {
			// Moving left (decreasing)
			for i := 1; i <= distance; i++ {
				pos := (position - i) % 100
				if pos < 0 {
					pos += 100
				}
				if pos == 0 {
					count++
				}
			}
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		} else if direction == 'R' {
			// Moving right (increasing)
			for i := 1; i <= distance; i++ {
				pos := (position + i) % 100
				if pos == 0 {
					count++
				}
			}
			position = (position + distance) % 100
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func main() {
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 Password: %d\n", result)

	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 Password: %d\n", result2)
}
