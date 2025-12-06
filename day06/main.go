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

func main() {
	result, err := solve("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Grand total: %d\n", result)
}
