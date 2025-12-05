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

func main() {
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Fresh ingredient IDs: %d\n", result)
}
