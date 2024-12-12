package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Garden struct {
	tiles map[string][]Point
	max_y int
	max_x int
	plots []Plot
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

func (g *Garden) recursiveSearch(parent, child Point, flower string) (next []Point, perimeter int) {
	if parent.y <= 0 {
		perimeter++
	} else {
		nxt, per := g.recursiveSearch(child, Point{x: child.x, y: child.y}, flower)

	}
	if parent.y >= g.max_y {
		perimeter++
	}

}

func (g *Garden) findPlots() {
	g.plots = []Plot{}
	for flower, tiles := range g.tiles {
		for _, tile := range tiles {

		}
		fmt.Println(flower, tile)
	}
}

func main() {
	garden := Garden{}
	garden.load("day12/input_test.txt")
	garden.findPlots()
	panic("test")
	fmt.Println(garden.tiles)
}
