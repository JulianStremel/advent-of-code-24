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

func move(grid [][]string, start []int, direction int) (steps int, path [][]int) {
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
	var ret [][]int
	for a, row := range grd {
		for b, col := range row {
			if col == "X" {
				cnt++
				ret = append(ret, []int{a, b})
			}
		}
	}
	return cnt, ret
}

func moveUp2(grid [][]string, start []int) (movable bool, direction int, step int) {
	if start[0] <= 0 {
		return false, 0, 0
	}
	if grid[start[0]-1][start[1]] != "#" {
		return true, 0, 1
	} else {
		return true, 1, 0
	}
}

func moveRight2(grid [][]string, start []int) (movable bool, direction int, step int) {
	if start[1] >= len(grid[0])-1 {
		return false, 1, 0
	}
	if grid[start[0]][start[1]+1] != "#" {
		return true, 1, 1
	} else {
		return true, 2, 0
	}
}

func moveDown2(grid [][]string, start []int) (movable bool, direction int, step int) {
	if start[0] >= len(grid)-1 {
		return false, 2, 0
	}
	if grid[start[0]+1][start[1]] != "#" {
		return true, 2, 1
	} else {
		return true, 3, 0
	}
}

func moveLeft2(grid [][]string, start []int) (movable bool, direction int, step int) {
	if start[1] <= 0 {
		return false, 3, 0
	}
	if grid[start[0]][start[1]-1] != "#" {
		return true, 3, 1
	} else {
		return true, 0, 0
	}
}

type dir struct {
	y   int
	x   int
	dir int
}

func (d dir) compare(c dir) bool {
	if c.x == d.x && c.y == d.y && c.dir == d.dir {
		return true
	}
	return false
}

func (d dir) isCointained(c []dir) bool {
	for _, di := range c {
		if d.compare(di) {
			return true
		}
	}
	return false
}

func move2(grid [][]string, start []int, direction int) bool {
	var mov bool = true
	var dirs []dir
	var tmp dir
	var stp int
	for mov {
		switch direction {
		case 0:
			mov, direction, stp = moveUp2(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if tmp.isCointained(dirs) {
					return true
				}
				start[0] = start[0] - 1
				dirs = append(dirs, tmp)
			}
		case 1:
			mov, direction, stp = moveRight2(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if tmp.isCointained(dirs) {
					return true
				}
				start[1] = start[1] + 1
				dirs = append(dirs, tmp)
			}
		case 2:
			mov, direction, stp = moveDown2(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if tmp.isCointained(dirs) {
					return true
				}
				start[0] = start[0] + 1
				dirs = append(dirs, tmp)
			}
		case 3:
			mov, direction, stp = moveLeft2(grid, start)
			if mov && stp > 0 {
				tmp = dir{start[0], start[1], direction}
				if tmp.isCointained(dirs) {
					return true
				}
				start[1] = start[1] - 1
				dirs = append(dirs, tmp)
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
	var start2 []int
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
	start2 = append(start2, start...)
	stp, path := move(fileLines, start, direction)

	var cnt = 0
	for a, coord := range path {
		fileLines2[coord[0]][coord[1]] = "#"
		if move2(fileLines2, []int{start2[0], start2[1]}, 0) {
			cnt++
		}
		fileLines2[coord[0]][coord[1]] = "."
		fmt.Printf("tested %d of %d possbile positions (%f%%)\n", a, len(path), float32(a)/float32(len(path))*100.0)
	}

	fmt.Printf("Part 1: %d", stp)
	fmt.Printf("Part 2: %d", cnt)
}
