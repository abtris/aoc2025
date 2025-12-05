package main

import (
	"bufio"
	"fmt"
	"os"
)

// countAdjacentRolls counts how many rolls of paper are adjacent to position (row, col)
func countAdjacentRolls(grid []string, row, col int) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Check all 8 adjacent positions
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // top-left, top, top-right
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // bottom-left, bottom, bottom-right
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check bounds
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

// solve counts how many rolls can be accessed (have < 4 adjacent rolls)
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
		line := scanner.Text()
		if len(line) > 0 {
			grid = append(grid, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Count accessible rolls
	accessible := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' {
				adjacentCount := countAdjacentRolls(grid, row, col)
				if adjacentCount < 4 {
					accessible++
				}
			}
		}
	}

	return accessible, nil
}

// solvePart2 iteratively removes accessible rolls until no more can be removed
func solvePart2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read the grid
	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			grid = append(grid, []byte(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	totalRemoved := 0

	// Keep removing accessible rolls until no more can be removed
	for {
		// Find all accessible rolls in current state
		var toRemove [][2]int

		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				if grid[row][col] == '@' {
					adjacentCount := countAdjacentRollsBytes(grid, row, col)
					if adjacentCount < 4 {
						toRemove = append(toRemove, [2]int{row, col})
					}
				}
			}
		}

		// If no rolls can be removed, we're done
		if len(toRemove) == 0 {
			break
		}

		// Remove all accessible rolls
		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = '.'
		}

		totalRemoved += len(toRemove)
	}

	return totalRemoved, nil
}

// countAdjacentRollsBytes is like countAdjacentRolls but works with [][]byte
func countAdjacentRollsBytes(grid [][]byte, row, col int) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Check all 8 adjacent positions
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // top-left, top, top-right
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // bottom-left, bottom, bottom-right
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check bounds
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

func main() {
	// Part 1
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Accessible rolls: %d\n", result)

	// Part 2
	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Total removed rolls: %d\n", result2)
}
