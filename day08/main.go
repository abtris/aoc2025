package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

type Edge struct {
	i, j     int
	distance float64
}

// UnionFind data structure
type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already in same set
	}

	// Union by size
	if uf.size[rootX] < uf.size[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}

	return true
}

func (uf *UnionFind) GetComponentSizes() []int {
	// Map from root to size
	sizeMap := make(map[int]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		sizeMap[root] = uf.size[root]
	}

	// Convert to slice
	var sizes []int
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}

	return sizes
}

func distance(p1, p2 Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	dz := float64(p1.z - p2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func solve(filename string, numConnections int) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read all points
	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, Point{x: x, y: y, z: z})
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	n := len(points)

	// Create all edges with distances
	var edges []Edge
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{i: i, j: j, distance: dist})
		}
	}

	// Sort edges by distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	// Use Union-Find to connect the closest pairs
	uf := NewUnionFind(n)

	// Try the numConnections shortest edges
	for i := 0; i < numConnections && i < len(edges); i++ {
		uf.Union(edges[i].i, edges[i].j)
	}

	// Get component sizes
	sizes := uf.GetComponentSizes()

	// Sort sizes in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	// Multiply the three largest
	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= sizes[i]
	}

	return result, nil
}

func solvePart2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read all points
	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, Point{x: x, y: y, z: z})
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	n := len(points)

	// Create all edges with distances
	var edges []Edge
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{i: i, j: j, distance: dist})
		}
	}

	// Sort edges by distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	// Use Union-Find to connect until all are in one circuit
	uf := NewUnionFind(n)
	numComponents := n

	for _, edge := range edges {
		if uf.Union(edge.i, edge.j) {
			numComponents--

			// Check if all are now in one circuit
			if numComponents == 1 {
				// This is the last connection needed
				// Multiply the X coordinates
				result := points[edge.i].x * points[edge.j].x
				return result, nil
			}
		}
	}

	return 0, fmt.Errorf("could not connect all junction boxes")
}

func main() {
	// Part 1
	result, err := solve("input", 1000)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Product of three largest circuits: %d\n", result)

	// Part 2
	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Product of X coordinates: %d\n", result2)
}
