package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Password: %d\n", result)
}
