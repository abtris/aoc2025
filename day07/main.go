package main

import (
	"bufio"
	"fmt"
	"os"
)

type Beam struct {
	row int
	col int
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read the grid
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if len(grid) == 0 {
		return 0, nil
	}

	// Find the starting position 'S'
	startCol := -1
	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == 'S' {
			startCol = col
			break
		}
	}

	if startCol == -1 {
		return 0, fmt.Errorf("no starting position found")
	}

	// Simulate the beam splitting
	splitCount := 0

	// Queue of active beams (row, col)
	beams := []Beam{{row: 1, col: startCol}} // Start from row 1 (below S)

	// Track which splitters have been hit to avoid counting the same split twice
	hitSplitters := make(map[Beam]bool)

	// Track which beam starting positions we've already processed
	processedBeams := make(map[Beam]bool)

	for len(beams) > 0 {
		// Process all beams at current level
		var nextBeams []Beam

		for _, beam := range beams {
			// Skip if we've already processed a beam starting from this position
			if processedBeams[beam] {
				continue
			}
			processedBeams[beam] = true

			// Move beam downward until it hits a splitter or exits
			currentRow := beam.row
			currentCol := beam.col

			// Move down until we hit a splitter or exit
			for currentRow < len(grid) {
				if currentCol < 0 || currentCol >= len(grid[currentRow]) {
					break
				}

				if grid[currentRow][currentCol] == '^' {
					// Check if we've already hit this splitter
					splitterKey := Beam{row: currentRow, col: currentCol}
					if !hitSplitters[splitterKey] {
						// Hit a splitter for the first time - count it
						splitCount++
						hitSplitters[splitterKey] = true
					}

					// Create two new beams (left and right)
					// Left beam starts at row+1, col-1
					if currentCol-1 >= 0 {
						leftBeam := Beam{row: currentRow + 1, col: currentCol - 1}
						if !processedBeams[leftBeam] {
							nextBeams = append(nextBeams, leftBeam)
						}
					}

					// Right beam starts at row+1, col+1
					if currentCol+1 < len(grid[currentRow]) {
						rightBeam := Beam{row: currentRow + 1, col: currentCol + 1}
						if !processedBeams[rightBeam] {
							nextBeams = append(nextBeams, rightBeam)
						}
					}

					break
				}

				currentRow++
			}
		}

		beams = nextBeams
	}

	return splitCount, nil
}

// solvePart2 counts the number of different timelines (paths) a particle can take
func solvePart2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read the grid
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if len(grid) == 0 {
		return 0, nil
	}

	// Find the starting position 'S'
	startCol := -1
	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == 'S' {
			startCol = col
			break
		}
	}

	if startCol == -1 {
		return 0, fmt.Errorf("no starting position found")
	}

	// Count all possible paths using DFS with memoization
	// Each path represents a timeline
	memo := make(map[Beam]int)
	timelineCount := countPathsMemo(grid, 1, startCol, memo)

	return timelineCount, nil
}

// countPathsMemo recursively counts all possible paths from a given position with memoization
func countPathsMemo(grid []string, row, col int, memo map[Beam]int) int {
	// Check memo
	key := Beam{row: row, col: col}
	if count, found := memo[key]; found {
		return count
	}

	// Move down until we hit a splitter or exit
	currentRow := row
	for currentRow < len(grid) {
		if col < 0 || col >= len(grid[currentRow]) {
			// Exited the grid - this is one complete path/timeline
			return 1
		}

		if grid[currentRow][col] == '^' {
			// Hit a splitter - particle takes both paths
			leftPaths := 0
			rightPaths := 0

			// Left path
			if col-1 >= 0 {
				leftPaths = countPathsMemo(grid, currentRow+1, col-1, memo)
			}

			// Right path
			if col+1 < len(grid[currentRow]) {
				rightPaths = countPathsMemo(grid, currentRow+1, col+1, memo)
			}

			// Total timelines is sum of both paths
			result := leftPaths + rightPaths
			memo[key] = result
			return result
		}

		currentRow++
	}

	// Exited the bottom of the grid - this is one complete path/timeline
	memo[key] = 1
	return 1
}

func main() {
	// Part 1
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Beam splits: %d\n", result)

	// Part 2
	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Timelines: %d\n", result2)
}
