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

// isGreenOrRed checks if a point is red or green
func isGreenOrRed(p Point, tiles []Point, greenTiles map[Point]bool) bool {
	// Check if it's a red tile
	for _, tile := range tiles {
		if tile.x == p.x && tile.y == p.y {
			return true
		}
	}
	// Check if it's a green tile
	return greenTiles[p]
}

// buildGreenTiles creates a map of green tiles on the edges only
// We don't precompute interior points due to large coordinate space
func buildGreenTiles(tiles []Point) map[Point]bool {
	greenTiles := make(map[Point]bool)
	n := len(tiles)

	// Add green tiles on edges between consecutive red tiles
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		p1 := tiles[i]
		p2 := tiles[next]

		// Add all tiles on the line between p1 and p2
		if p1.x == p2.x {
			// Vertical line
			minY := p1.y
			maxY := p2.y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			for y := minY; y <= maxY; y++ {
				greenTiles[Point{x: p1.x, y: y}] = true
			}
		} else if p1.y == p2.y {
			// Horizontal line
			minX := p1.x
			maxX := p2.x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			for x := minX; x <= maxX; x++ {
				greenTiles[Point{x: x, y: p1.y}] = true
			}
		}
	}

	// Note: We don't precompute interior points because the coordinate space is too large
	// Interior points will be checked using isInsidePolygon() on demand

	return greenTiles
}

// isInsidePolygon uses ray casting algorithm
func isInsidePolygon(p Point, polygon []Point) bool {
	n := len(polygon)
	inside := false

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		xi, yi := polygon[i].x, polygon[i].y
		xj, yj := polygon[j].x, polygon[j].y

		if ((yi > p.y) != (yj > p.y)) &&
			(p.x < (xj-xi)*(p.y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
	}

	return inside
}

// isRectangleValid checks if a rectangle only contains red or green tiles
// by sampling points along its perimeter and checking a few interior points
func isRectangleValid(minX, maxX, minY, maxY int, polygon []Point, greenTiles map[Point]bool, redTiles map[Point]bool) bool {
	// Check corners
	corners := []Point{
		{x: minX, y: minY},
		{x: maxX, y: minY},
		{x: minX, y: maxY},
		{x: maxX, y: maxY},
	}

	for _, p := range corners {
		if !redTiles[p] && !greenTiles[p] && !isInsidePolygon(p, polygon) {
			return false
		}
	}

	// Sample points along the edges (every 100 units to keep it fast)
	step := 100

	// Top and bottom edges
	for x := minX; x <= maxX; x += step {
		for _, y := range []int{minY, maxY} {
			p := Point{x: x, y: y}
			if !redTiles[p] && !greenTiles[p] && !isInsidePolygon(p, polygon) {
				return false
			}
		}
	}

	// Left and right edges
	for y := minY; y <= maxY; y += step {
		for _, x := range []int{minX, maxX} {
			p := Point{x: x, y: y}
			if !redTiles[p] && !greenTiles[p] && !isInsidePolygon(p, polygon) {
				return false
			}
		}
	}

	// Sample a few interior points
	midX := (minX + maxX) / 2
	midY := (minY + maxY) / 2
	interiorPoints := []Point{
		{x: midX, y: midY},
		{x: (minX + midX) / 2, y: (minY + midY) / 2},
		{x: (maxX + midX) / 2, y: (maxY + midY) / 2},
	}

	for _, p := range interiorPoints {
		if !redTiles[p] && !greenTiles[p] && !isInsidePolygon(p, polygon) {
			return false
		}
	}

	return true
}

func solvePart2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read all red tile positions (in order)
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

	// Build green tiles map
	greenTiles := buildGreenTiles(tiles)

	// Create a set of red tiles for faster lookup
	redTiles := make(map[Point]bool)
	for _, t := range tiles {
		redTiles[t] = true
	}

	// Find the largest rectangle that only contains red or green tiles
	maxArea := 0
	n := len(tiles)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// Check if rectangle from tiles[i] to tiles[j] only contains red/green
			rectMinX := tiles[i].x
			rectMaxX := tiles[j].x
			if rectMinX > rectMaxX {
				rectMinX, rectMaxX = rectMaxX, rectMinX
			}

			rectMinY := tiles[i].y
			rectMaxY := tiles[j].y
			if rectMinY > rectMaxY {
				rectMinY, rectMaxY = rectMaxY, rectMinY
			}

			width := rectMaxX - rectMinX + 1
			height := rectMaxY - rectMinY + 1
			area := width * height

			// Skip if this can't beat the current max
			if area <= maxArea {
				continue
			}

			// Check if rectangle is valid
			if isRectangleValid(rectMinX, rectMaxX, rectMinY, rectMaxY, tiles, greenTiles, redTiles) {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func main() {
	// Part 1
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Largest rectangle area: %d\n", result)

	// Part 2
	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Largest rectangle area (red/green only): %d\n", result2)
}
