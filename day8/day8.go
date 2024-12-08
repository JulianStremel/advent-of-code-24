package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	y int
	x int
}

func (p Position) isContained(l []Position) bool {
	for _, l := range l {
		if l == p {
			return true
		}
	}
	return false
}

func antinodes1(a, b Position, max_y, max_x int) []Position { // c...a...b...d
	tmp := []Position{}
	c := Position{a.y - (b.y - a.y), a.x - (b.x - a.x)}
	d := Position{b.y - (a.y - b.y), b.x - (a.x - b.x)}
	if !(c.x > max_x || c.x < 0 || c.y > max_y || c.y < 0) {
		tmp = append(tmp, c)
	}
	if !(d.x > max_x || d.x < 0 || d.y > max_y || d.y < 0) {
		tmp = append(tmp, d)
	}
	return tmp
}

func antinodes2(a, b Position, max_y, max_x int) []Position { // c...a...b...d
	tmp := []Position{}
	diff_x := b.x - a.x
	diff_y := b.y - a.y
	cnt := 0
	for !(a.x-(cnt*diff_x) > max_x || a.x-(cnt*diff_x) < 0 || a.y-(cnt*diff_y) > max_y || a.y-(cnt*diff_y) < 0) {
		tmp = append(tmp, Position{a.x - (cnt * diff_x), a.y - (cnt * diff_y)})
		cnt++
	}
	cnt = -1
	for !(a.x-(cnt*diff_x) > max_x || a.x-(cnt*diff_x) < 0 || a.y-(cnt*diff_y) > max_y || a.y-(cnt*diff_y) < 0) {
		tmp = append(tmp, Position{a.x - (cnt * diff_x), a.y - (cnt * diff_y)})
		cnt--
	}
	return tmp
}

type Game struct {
	input     [][]string
	nodes     map[string][]Position
	antinodes []Position
	sum       int
}

func (g *Game) init(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var a = 0
	g.nodes = make(map[string][]Position)
	for fileScanner.Scan() {
		var line = strings.Split(fileScanner.Text(), "")
		g.input = append(g.input, line)
		for b, char := range line {
			if char != "." {
				g.nodes[char] = append(g.nodes[char], Position{a, b})
			}
		}
		a++
	}
}

func (g *Game) insertAntinode(nodes map[string][]Position) {
	for _, positions := range nodes {
		for _, position := range positions {
			if !position.isContained(g.antinodes) {
				g.antinodes = append(g.antinodes, position)
				g.sum += 1
			}
		}
	}
}

func (g *Game) computeAntinodes1() {
	g.antinodes = []Position{}
	g.sum = 0
	tmp := make(map[string][]Position)
	for signal, positions := range g.nodes {
		for a := range (len(positions) / 2) + 1 {
			for b := range (len(positions) / 2) + 1 {
				if positions[a] == positions[len(positions)-1-b] {
					break
				}
				tmp[signal] = append(tmp[signal], antinodes1(positions[a], positions[len(positions)-1-b], len(g.input)-1, len(g.input)-1)...)
			}
		}
	}
	g.insertAntinode(tmp)
}

func (g *Game) computeAntinodes2() {
	g.antinodes = []Position{}
	g.sum = 0
	tmp := make(map[string][]Position)
	for signal, positions := range g.nodes {
		for a := range (len(positions) / 2) + 1 {
			for b := range (len(positions) / 2) + 1 {
				if positions[a] == positions[len(positions)-1-b] {
					break
				}
				tmp[signal] = append(tmp[signal], antinodes2(positions[a], positions[len(positions)-1-b], len(g.input)-1, len(g.input)-1)...)
			}
		}
	}
	g.insertAntinode(tmp)
}

func main() {
	game := Game{}
	game.init("day8/input.txt")
	game.computeAntinodes1()
	fmt.Printf("Part 1: %d\n", game.sum)

	game = Game{}
	game.init("day8/input.txt")
	game.computeAntinodes2()
	fmt.Printf("Part 2: %d\n", game.sum)

}
