package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func moveUp(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[0] <= 0 {
		//grid[start[0]][start[1]] = "X"
		return false, 0, 0, grid
	}
	if grid[start[0]-1][start[1]] != "#" {
		//grid[start[0]][start[1]] = "X"
		return true, 0, 1, grid
	} else {
		//grid[start[0]][start[1]] = "X"
		return true, 1, 0, grid
	}
}

func moveRight(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[1] >= len(grid[0])-1 {
		//grid[start[0]][start[1]] = "X"
		return false, 1, 0, grid
	}
	if grid[start[0]][start[1]+1] != "#" {
		//grid[start[0]][start[1]] = "X"
		return true, 1, 1, grid
	} else {
		//grid[start[0]][start[1]] = "X"
		return true, 2, 0, grid
	}
}

func moveDown(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[0] >= len(grid)-1 {
		//grid[start[0]][start[1]] = "X"
		return false, 2, 0, grid
	}
	if grid[start[0]+1][start[1]] != "#" {
		//grid[start[0]][start[1]] = "X"
		return true, 2, 1, grid
	} else {
		//grid[start[0]][start[1]] = "X"
		return true, 3, 0, grid
	}
}

func moveLeft(grid [][]string, start []int) (movable bool, direction int, steps int, grd [][]string) {
	if start[1] <= 0 {
		//grid[start[0]][start[1]] = "X"
		return false, 3, 0, grid
	}
	if grid[start[0]][start[1]-1] != "#" {
		//grid[start[0]][start[1]] = "X"
		return true, 3, 1, grid
	} else {
		//grid[start[0]][start[1]] = "X"
		return true, 0, 0, grid
	}
}

func move(grid [][]string, start []int, direction int) (steps int) {

	var mov bool = true
	var stp int
	steps = 0
	for mov {
		switch direction {
		case 0:
			mov, direction, stp, grid = moveUp(grid, start)
			if mov && stp > 0 {
				start[0] = start[0] - 1
			}
			steps += stp
		case 1:
			mov, direction, stp, grid = moveRight(grid, start)
			if mov && stp > 0 {
				start[1] = start[1] + 1
			}
			steps += stp
		case 2:
			mov, direction, stp, grid = moveDown(grid, start)
			if mov && stp > 0 {
				start[0] = start[0] + 1
			}
			steps += stp
		case 3:
			mov, direction, stp, grid = moveLeft(grid, start)
			if mov && stp > 0 {
				start[1] = start[1] - 1
			}
			steps += stp
		default:
			fmt.Println("no way")
		}
	}
	var cnt = 0
	for _, row := range grid {
		for _, col := range row {
			if col == "X" {
				cnt++
			}
		}
	}
	return cnt
}

func checkLoop(grid [][]string, start []int, direction int) bool {

	var mov bool = true
	var stp int
	type dir struct {
		Y   int
		X   int
		Dir int
	}
	var directions []dir //[[Y,X,Dir][Y,X,Dir]] // if location gets passed in same direction twice = loop

	var tmp dir
	for mov {
		switch direction {
		case 0:
			mov, direction, stp, grid = moveUp(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if slices.Contains(directions, tmp) {
					return true
				}
				directions = append(directions, tmp)
				start[0] = start[0] - 1
			}

		case 1:
			mov, direction, stp, grid = moveRight(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if slices.Contains(directions, tmp) {
					return true
				}
				directions = append(directions, tmp)
				start[1] = start[1] + 1
			}

		case 2:
			mov, direction, stp, grid = moveDown(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if slices.Contains(directions, tmp) {
					return true
				}
				directions = append(directions, tmp)
				start[0] = start[0] + 1
			}

		case 3:
			mov, direction, stp, grid = moveLeft(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if slices.Contains(directions, tmp) {
					return true
				}
				directions = append(directions, tmp)
				start[1] = start[1] - 1
			}

		default:
			fmt.Println("no way")
		}
	}
	return false

}

func main() {
	var fileLines [][]string
	var fileLines2 [][]string
	readFile, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var start []int
	var direction int // 0 -> up | 1 -> right | 2 -> down | 3 -> left
	var line []string
	var line2 []string
	var a = 0
	for fileScanner.Scan() {
		line = strings.Split(fileScanner.Text(), "")
		line2 = strings.Split(fileScanner.Text(), "")
		fileLines = append(fileLines, line)
		fileLines2 = append(fileLines2, line2)
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

	var cnt = 0
	for a, line := range fileLines2 {
		for b, row := range line {
			if row != "." {
				continue
			}
			fileLines2[a][b] = "#"
			if checkLoop(fileLines2, start, direction) {
				cnt++
				fmt.Println(cnt)
			}
			fileLines2[a][b] = ":"
			fmt.Println(a, b)
		}
	}

	fmt.Println(fileLines2)
	fmt.Println(cnt)

}
