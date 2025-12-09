package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read all red tile positions
	var tiles []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		tiles = append(tiles, Point{x: x, y: y})
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Find the largest rectangle by checking all pairs
	maxArea := 0
	n := len(tiles)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// Calculate area of rectangle with tiles[i] and tiles[j] as opposite corners
			// Add 1 to include both endpoints
			width := abs(tiles[i].x-tiles[j].x) + 1
			height := abs(tiles[i].y-tiles[j].y) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func main() {
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Largest rectangle area: %d\n", result)
}
