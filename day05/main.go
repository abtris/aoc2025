package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

// isFresh checks if an ID falls within any of the fresh ranges
func isFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var ranges []Range
	var availableIDs []int
	parsingRanges := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Empty line separates ranges from available IDs
		if len(line) == 0 {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			// Parse range like "3-5"
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				start, err1 := strconv.Atoi(parts[0])
				end, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil {
					ranges = append(ranges, Range{start: start, end: end})
				}
			}
		} else {
			// Parse available ID
			id, err := strconv.Atoi(line)
			if err == nil {
				availableIDs = append(availableIDs, id)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Count how many available IDs are fresh
	freshCount := 0
	for _, id := range availableIDs {
		if isFresh(id, ranges) {
			freshCount++
		}
	}

	return freshCount, nil
}

// solvePart2 counts total unique IDs covered by all ranges
func solvePart2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var ranges []Range

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Empty line means we're done with ranges
		if len(line) == 0 {
			break
		}

		// Parse range like "3-5"
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				ranges = append(ranges, Range{start: start, end: end})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Merge overlapping ranges and count total IDs
	return countTotalIDs(ranges), nil
}

// countTotalIDs merges overlapping ranges and counts total unique IDs
func countTotalIDs(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start position
	// Using a simple bubble sort since we need to sort
	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[j].start < ranges[i].start {
				ranges[i], ranges[j] = ranges[j], ranges[i]
			}
		}
	}

	// Merge overlapping ranges
	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		last := &merged[len(merged)-1]

		// Check if current range overlaps or is adjacent to the last merged range
		if current.start <= last.end+1 {
			// Merge: extend the last range if needed
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			// No overlap, add as new range
			merged = append(merged, current)
		}
	}

	// Count total IDs in merged ranges
	total := 0
	for _, r := range merged {
		total += r.end - r.start + 1
	}

	return total
}

func main() {
	// Part 1
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Fresh ingredient IDs: %d\n", result)

	// Part 2
	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Total fresh IDs in ranges: %d\n", result2)
}
