package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pluto struct {
	pebbles []int
}

func (p *Pluto) load(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	p.pebbles = []int{}

	for fileScanner.Scan() {
		tmpStr := []string{}
		tmp := fileScanner.Text()
		if err != nil {
			panic(err)
		}
		tmpStr = strings.Split(tmp, " ")
		for _, a := range tmpStr {
			num, err := strconv.Atoi(a)
			if err != nil {
				panic(err)
			}
			p.pebbles = append(p.pebbles, num)
		}
	}
}

func (p *Pluto) blink() {
	var index = 0
	for {
		if index > len(p.pebbles)-1 {
			break
		}
		if p.pebbles[index] == 0 {
			p.pebbles[index] = 1
			index++
			continue
		}
		tmp := strconv.Itoa(p.pebbles[index])
		if len(tmp)%2 == 0 {
			a, err := strconv.Atoi(tmp[:(len(tmp) / 2)])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(tmp[(len(tmp) / 2):])
			if err != nil {
				panic(err)
			}
			p.pebbles[index] = a
			p.pebbles = slices.Insert(p.pebbles, index+1, b)
			index += 2
			continue
		}
		p.pebbles[index] = p.pebbles[index] * 2024
		index++
	}
}

func main() {
	pluto := Pluto{}
	pluto.load("day11/input_test.txt")
	fmt.Println(pluto.pebbles)
	for a := range 75 {
		pluto.blink()
		fmt.Println(a)
	}
	fmt.Println(len(pluto.pebbles))
}
