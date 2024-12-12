package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	input [][]int
	paths []PathTile
	sum   int
}

type PathTile struct {
	height   int
	x        int
	y        int
	previous *PathTile
	next     []*PathTile
	ends     []*PathTile
}

func (m *Map) load(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	m.input = [][]int{}
	a := 0
	for fileScanner.Scan() {
		line := []int{}
		for b, num := range strings.Split(fileScanner.Text(), "") {
			tmp, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			if tmp == 0 {
				m.paths = append(m.paths, PathTile{height: 0, y: a, x: b, previous: nil})
			}
			line = append(line, tmp)
		}
		m.input = append(m.input, line)
		a++
	}
}

func filterUnique(tiles []*PathTile) []*PathTile {
	seen := make(map[[2]int]bool)
	var result []*PathTile

	for _, tile := range tiles {
		if tile == nil {
			continue
		}
		key := [2]int{tile.x, tile.y}
		if !seen[key] {
			seen[key] = true
			result = append(result, tile)
		}
	}

	return result
}

func (m *Map) findPath1() {
	for _, path := range m.paths {
		b, c := m.findNext(path)
		if b {
			path.ends = append(path.ends, c...)
			path.ends = filterUnique(path.ends)
			m.sum += len(path.ends)
		}
	}
}

func (m *Map) findPath2() {
	for _, path := range m.paths {
		b, c := m.findNext(path)
		if b {
			path.ends = append(path.ends, c...)
			//path.ends = filterUnique(path.ends)
			m.sum += len(path.ends)
		}
	}
}

func (m *Map) findNext(p PathTile) (trailEnd bool, end []*PathTile) {
	if p.height == 9 {
		return true, []*PathTile{&p}
	}
	cnt := 0
	if p.y > 0 && m.input[p.y-1][p.x] == p.height+1 {
		tmp := &PathTile{height: p.height + 1, y: p.y - 1, x: p.x, previous: &p}
		a, b := m.findNext(*tmp)
		if a {
			p.next = append(p.next, tmp)
			end = append(end, b...)
			cnt++
		}
	}
	if p.y < len(m.input)-1 && m.input[p.y+1][p.x] == p.height+1 {
		tmp := &PathTile{height: p.height + 1, y: p.y + 1, x: p.x, previous: &p}
		a, b := m.findNext(*tmp)
		if a {
			p.next = append(p.next, tmp)
			end = append(end, b...)
			cnt++
		}
	}
	if p.x > 0 && m.input[p.y][p.x-1] == p.height+1 {
		tmp := &PathTile{height: p.height + 1, y: p.y, x: p.x - 1, previous: &p}
		a, b := m.findNext(*tmp)
		if a {
			p.next = append(p.next, tmp)
			end = append(end, b...)
			cnt++
		}
	}
	if p.x < len(m.input)-1 && m.input[p.y][p.x+1] == p.height+1 {
		tmp := &PathTile{height: p.height + 1, y: p.y, x: p.x + 1, previous: &p}
		a, b := m.findNext(*tmp)
		if a {
			p.next = append(p.next, tmp)
			end = append(end, b...)
			cnt++
		}
	}
	if cnt > 0 {
		return true, end
	}
	return false, []*PathTile{}
}

func main() {
	mp := Map{}
	mp.load("day10/input.txt")
	mp.findPath1()
	fmt.Println("Part 1: ", mp.sum)

	mp = Map{}
	mp.load("day10/input.txt")
	mp.findPath2()
	fmt.Println("Part 2: ", mp.sum)
}
