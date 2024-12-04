package main

import (
	"bufio"
	"fmt"
	"os"
)

func isInBounds(row, col, rows, cols int) bool {
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return false
	}
	return true
}

func matchWord(grid []string, startRow, startCol, dirRow, dirCol int) bool {
	for i := 0; i < 4; i++ {
		r := startRow + i*dirRow
		c := startCol + i*dirCol

		// If out of bounds or character does not match, return false
		if !isInBounds(r, c, len(grid), len(grid[0])) || grid[r][c] != "XMAS"[i] {
			return false
		}
	}
	return true
}

func countXmas(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [8][2]int{
		{0, 1},   // Horizontal right
		{0, -1},  // Horizontal left
		{1, 0},   // Vertical down
		{-1, 0},  // Vertical up
		{1, 1},   // Diagonal top-left to bottom-right
		{-1, -1}, // Diagonal bottom-right to top-left
		{1, -1},  // Diagonal top-right to bottom-left
		{-1, 1},  // Diagonal bottom-left to top-right
	}

	for row := range rows {
		for col := range cols {
			for _, dir := range directions {
				if matchWord(grid, row, col, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func countXMasPattern(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for row := range rows - 2 {
		for col := range cols - 2 { // c war 1 ist jetzt 0
			diag1 := string([]byte{grid[row][col], grid[row+1][col+1], grid[row+2][col+2]})
			diag2 := string([]byte{grid[row][col+2], grid[row+1][col+1], grid[row+2][col]})

			if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
				count++
			}
		}
	}

	return count
}

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	partOneResult := countXmas(grid)
	fmt.Printf("Part 1: %d\n", partOneResult)

	partTwoResult := countXMasPattern(grid)
	fmt.Printf("Part 2: %d\n", partTwoResult)
}
