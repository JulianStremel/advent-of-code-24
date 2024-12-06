package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func moveUp(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[0] <= 0 {
		grid[start[0]][start[1]] = "X"
		return false, 0, 0, grid
	}
	if grid[start[0]-1][start[1]] != "#" {
		grid[start[0]][start[1]] = "X"
		return true, 0, 1, grid
	} else {
		grid[start[0]][start[1]] = "X"
		return true, 1, 0, grid
	}
}

func moveRight(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[1] >= len(grid[0])-1 {
		grid[start[0]][start[1]] = "X"
		return false, 1, 0, grid
	}
	if grid[start[0]][start[1]+1] != "#" {
		grid[start[0]][start[1]] = "X"
		return true, 1, 1, grid
	} else {
		grid[start[0]][start[1]] = "X"
		return true, 2, 0, grid
	}
}

func moveDown(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[0] >= len(grid)-1 {
		grid[start[0]][start[1]] = "X"
		return false, 2, 0, grid
	}
	if grid[start[0]+1][start[1]] != "#" {
		grid[start[0]][start[1]] = "X"
		return true, 2, 1, grid
	} else {
		grid[start[0]][start[1]] = "X"
		return true, 3, 0, grid
	}
}

func moveLeft(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[1] <= 0 {
		grid[start[0]][start[1]] = "X"
		return false, 3, 0, grid
	}
	if grid[start[0]][start[1]-1] != "#" {
		grid[start[0]][start[1]] = "X"
		return true, 3, 1, grid
	} else {
		grid[start[0]][start[1]] = "X"
		return true, 0, 0, grid
	}
}

func move(grid [][]string, start []int, direction int) (steps int) {

	var mov bool = true
	var stp int
	var dir = direction
	var strt = start
	var grd = grid
	steps = 0
	for mov {
		switch dir {
		case 0:
			mov, dir, stp, grd = moveUp(grd, strt)
			if mov && stp > 0 {
				strt[0] = strt[0] - 1
			}
			steps += stp
		case 1:
			mov, dir, stp, grd = moveRight(grd, strt)
			if mov && stp > 0 {
				strt[1] = strt[1] + 1
			}
			steps += stp
		case 2:
			mov, dir, stp, grd = moveDown(grd, strt)
			if mov && stp > 0 {
				strt[0] = strt[0] + 1
			}
			steps += stp
		case 3:
			mov, dir, stp, grd = moveLeft(grd, strt)
			if mov && stp > 0 {
				strt[1] = strt[1] - 1
			}
			steps += stp
		default:
			fmt.Println("no way")
		}
	}
	var cnt = 0
	for _, row := range grd {
		for _, col := range row {
			if col == "X" {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	var fileLines [][]string

	readFile, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var start []int
	var direction int // 0 -> up | 1 -> right | 2 -> down | 3 -> left
	var line []string
	var a = 0
	for fileScanner.Scan() {
		line = strings.Split(fileScanner.Text(), "")
		fileLines = append(fileLines, line)
		for b, letter := range line {
			switch letter {
			case "^":
				start = append(start, a, b)
				direction = 0
			case ">":
				start = append(start, a, b)
				direction = 1
			case "v":
				start = append(start, a, b)
				direction = 2
			case "<":
				start = append(start, a, b)
				direction = 3
			}
		}
		a++
	}
	readFile.Close()

	fmt.Println(move(fileLines, start, direction))

	fmt.Println(fileLines)

}
