package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid := readGrid("input.txt")
	count := countXMAS(grid)
	fmt.Printf("Found %d instances of X-MAS\n", count)
}

func readGrid(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			grid = append(grid, []rune(line))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return grid
}

func countXMAS(grid [][]rune) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	// Check each position that could be the center of an X
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			// The center must be 'A'
			if grid[row][col] != 'A' {
				continue
			}

			// Check all possible combinations of MAS in X pattern
			if isValidXMAS(grid, row, col) {
				count++
			}
		}
	}
	return count
}

func isValidXMAS(grid [][]rune, centerRow, centerCol int) bool {
	// Check if we can form a valid X-MAS pattern
	// We need to check both MAS and SAM in each diagonal

	// Check upper-left to lower-right diagonal
	ul_lr := string([]rune{
		grid[centerRow-1][centerCol-1],
		grid[centerRow][centerCol],
		grid[centerRow+1][centerCol+1],
	})

	// Check upper-right to lower-left diagonal
	ur_ll := string([]rune{
		grid[centerRow-1][centerCol+1],
		grid[centerRow][centerCol],
		grid[centerRow+1][centerCol-1],
	})

	// Both diagonals must contain either MAS or SAM
	return isValidMAS(ul_lr) && isValidMAS(ur_ll)
}

func isValidMAS(s string) bool {
	// Check if the string is either MAS or SAM
	return s == "MAS" || s == "SAM"
}