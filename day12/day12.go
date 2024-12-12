package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Garden struct {
	tiles   map[string][]Point
	max_y   int
	max_x   int
	plots   []Plot
	visited []Point
}

type Plot struct {
	flower    string
	perimeter int
	tiles     []Point
}

type Point struct {
	y int
	x int
}

func (g *Garden) load(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	g.tiles = map[string][]Point{}

	var y = 0
	for fileScanner.Scan() {
		tmp := strings.Split(fileScanner.Text(), "")
		for x, thing := range tmp {
			if _, ok := g.tiles[thing]; !ok {
				g.tiles[thing] = []Point{}
			}
			g.tiles[thing] = append(g.tiles[thing], Point{y: y, x: x})
		}
		y++
		g.max_x = len(tmp) - 1
	}
	g.max_y = y
}

func (g *Garden) recursiveSearch(parent Point, flower string) (perimeter int) {
	// append calling node to visited
	g.visited = append(g.visited, parent)

	// check if outer bounds for y
	if parent.y <= 0 {
		perimeter++

	} else {
		//fmt.Println("walking up")
		tmp := Point{y: parent.y - 1, x: parent.x}

		if !slices.Contains(g.visited, tmp) {
			if slices.Contains(g.tiles[flower], tmp) {
				perimeter += g.recursiveSearch(tmp, flower)
			} else {
				perimeter++
			}
		}
	}
	if parent.y >= g.max_y {
		perimeter++

	} else {
		//fmt.Println("walking down")
		tmp := Point{y: parent.y + 1, x: parent.x}
		if !slices.Contains(g.visited, tmp) {
			if slices.Contains(g.tiles[flower], tmp) {
				perimeter += g.recursiveSearch(tmp, flower)
			} else {
				perimeter++
			}
		}
	}

	if parent.x <= 0 {
		perimeter++

	} else {
		//fmt.Println("walking left")
		tmp := Point{y: parent.y, x: parent.x - 1}

		if !slices.Contains(g.visited, tmp) {
			if slices.Contains(g.tiles[flower], tmp) {
				perimeter += g.recursiveSearch(tmp, flower)
			} else {
				perimeter++
			}
		}
	}
	if parent.x >= g.max_x {
		perimeter++

	} else {
		//fmt.Println("walking right")
		tmp := Point{y: parent.y, x: parent.x + 1}
		if !slices.Contains(g.visited, tmp) {
			if slices.Contains(g.tiles[flower], tmp) {
				perimeter += g.recursiveSearch(tmp, flower)
			} else {
				perimeter++
			}
		}
	}

	return perimeter
}

func (g *Garden) findPlots() {
	g.plots = []Plot{}
	sum := 0
	for flower, tiles := range g.tiles {
		for range tiles {
			if len(g.tiles[flower]) == 0 {
				break
			}
			tile := g.tiles[flower][0]
			g.visited = []Point{}
			tmp := g.recursiveSearch(tile, flower)
			plt := Plot{flower: flower, perimeter: tmp, tiles: g.visited}
			g.plots = append(g.plots, plt)
			fmt.Println("start", tile, flower, len(g.visited), tmp)
			sum += len(g.visited) * tmp
			for _, a := range g.visited {
				ind := slices.Index(g.tiles[flower], a)
				g.tiles[flower] = slices.Delete(g.tiles[flower], ind, ind+1)
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	garden := Garden{}
	garden.load("day12/input.txt")
	garden.findPlots()
}
