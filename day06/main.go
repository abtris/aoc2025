package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	numbers   []int
	operation string // "*" or "+"
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if len(lines) == 0 {
		return 0, nil
	}

	// Find the width of the worksheet
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Pad all lines to the same width
	for i := range lines {
		if len(lines[i]) < maxWidth {
			lines[i] += strings.Repeat(" ", maxWidth-len(lines[i]))
		}
	}

	// Extract problems column by column
	var problems []Problem

	for col := 0; col < maxWidth; col++ {
		// Check if this column is all spaces (separator)
		allSpaces := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] != ' ' {
				allSpaces = false
				break
			}
		}

		if allSpaces {
			continue
		}

		// Extract this column's content
		var columnChars []byte
		for row := 0; row < len(lines); row++ {
			columnChars = append(columnChars, lines[row][col])
		}

		// Check if this is the start of a new problem
		// A problem starts when we find a non-space character
		// We need to collect all columns that belong to this problem
		if columnChars[0] != ' ' || hasNonSpaceInColumn(columnChars) {
			// Collect all consecutive columns that are part of this problem
			problem := extractProblem(lines, col, maxWidth)
			if problem != nil {
				problems = append(problems, *problem)
				// Skip the columns we just processed
				col += getProblemWidth(lines, col, maxWidth) - 1
			}
		}
	}

	// Calculate grand total
	grandTotal := 0
	for _, problem := range problems {
		result := calculateProblem(problem)
		grandTotal += result
	}

	return grandTotal, nil
}

func hasNonSpaceInColumn(chars []byte) bool {
	for _, c := range chars {
		if c != ' ' {
			return true
		}
	}
	return false
}

func extractProblem(lines []string, startCol, maxWidth int) *Problem {
	// Find the width of this problem (until we hit a column of all spaces)
	endCol := startCol
	for endCol < maxWidth {
		allSpaces := true
		for row := 0; row < len(lines); row++ {
			if lines[row][endCol] != ' ' {
				allSpaces = false
				break
			}
		}
		if allSpaces {
			break
		}
		endCol++
	}

	// Extract the problem from startCol to endCol
	var numbers []int
	var operation string

	// Last row contains the operation
	opRow := len(lines) - 1
	opStr := strings.TrimSpace(lines[opRow][startCol:endCol])
	if opStr == "*" || opStr == "+" {
		operation = opStr
	} else {
		return nil
	}

	// Other rows contain numbers
	for row := 0; row < opRow; row++ {
		numStr := strings.TrimSpace(lines[row][startCol:endCol])
		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	if len(numbers) == 0 {
		return nil
	}

	return &Problem{
		numbers:   numbers,
		operation: operation,
	}
}

func getProblemWidth(lines []string, startCol, maxWidth int) int {
	width := 0
	for col := startCol; col < maxWidth; col++ {
		allSpaces := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] != ' ' {
				allSpaces = false
				break
			}
		}
		if allSpaces {
			break
		}
		width++
	}
	return width
}

func calculateProblem(problem Problem) int {
	if len(problem.numbers) == 0 {
		return 0
	}

	result := problem.numbers[0]
	for i := 1; i < len(problem.numbers); i++ {
		if problem.operation == "*" {
			result *= problem.numbers[i]
		} else if problem.operation == "+" {
			result += problem.numbers[i]
		}
	}

	return result
}

// solvePart2 reads problems right-to-left with each column being a digit position
func solvePart2(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if len(lines) == 0 {
		return 0, nil
	}

	// Find the width of the worksheet
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Pad all lines to the same width
	for i := range lines {
		if len(lines[i]) < maxWidth {
			lines[i] += strings.Repeat(" ", maxWidth-len(lines[i]))
		}
	}

	// Extract problems by scanning right-to-left
	var problems []Problem

	col := maxWidth - 1
	for col >= 0 {
		// Skip columns of all spaces
		allSpaces := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] != ' ' {
				allSpaces = false
				break
			}
		}

		if allSpaces {
			col--
			continue
		}

		// Found start of a problem (from the right)
		// Find the left boundary of this problem
		startCol := col
		for startCol >= 0 {
			allSpaces := true
			for row := 0; row < len(lines); row++ {
				if lines[row][startCol] != ' ' {
					allSpaces = false
					break
				}
			}
			if allSpaces {
				break
			}
			startCol--
		}
		startCol++ // Move back to first non-space column

		// Extract problem from startCol to col (inclusive), reading right-to-left
		problem := extractProblemPart2(lines, startCol, col)
		if problem != nil {
			problems = append(problems, *problem)
		}

		col = startCol - 1
	}

	// Calculate grand total
	var grandTotal int64
	for _, problem := range problems {
		result := calculateProblemPart2(problem)
		grandTotal += result
	}

	return grandTotal, nil
}

func extractProblemPart2(lines []string, startCol, endCol int) *Problem {
	opRow := len(lines) - 1

	// Get operation from the last row
	var operation string
	for col := endCol; col >= startCol; col-- {
		if lines[opRow][col] != ' ' {
			operation = string(lines[opRow][col])
			break
		}
	}

	if operation != "*" && operation != "+" {
		return nil
	}

	// Read numbers right-to-left, each column is a number
	var numbers []int64
	for col := endCol; col >= startCol; col-- {
		// Check if this column has any digits
		hasDigit := false
		for row := 0; row < opRow; row++ {
			if lines[row][col] >= '0' && lines[row][col] <= '9' {
				hasDigit = true
				break
			}
		}

		if !hasDigit {
			continue
		}

		// Read this column top-to-bottom to form a number
		var numStr string
		for row := 0; row < opRow; row++ {
			if lines[row][col] >= '0' && lines[row][col] <= '9' {
				numStr += string(lines[row][col])
			}
		}

		if numStr != "" {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	if len(numbers) == 0 {
		return nil
	}

	// Convert int64 numbers to int for Problem struct
	intNumbers := make([]int, len(numbers))
	for i, n := range numbers {
		intNumbers[i] = int(n)
	}

	return &Problem{
		numbers:   intNumbers,
		operation: operation,
	}
}

func calculateProblemPart2(problem Problem) int64 {
	if len(problem.numbers) == 0 {
		return 0
	}

	result := int64(problem.numbers[0])
	for i := 1; i < len(problem.numbers); i++ {
		if problem.operation == "*" {
			result *= int64(problem.numbers[i])
		} else if problem.operation == "+" {
			result += int64(problem.numbers[i])
		}
	}

	return result
}

func main() {
	// Part 1
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Grand total: %d\n", result)

	// Part 2
	result2, err := solvePart2("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Grand total: %d\n", result2)
}
