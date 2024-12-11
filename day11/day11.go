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
	sum     int
	cache   map[int]map[int]int
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

func (p *Pluto) blink2(num int) {
	p.cache = make(map[int]map[int]int)
	for b := range num {
		p.cache[b] = make(map[int]int)
	}

	for _, a := range p.pebbles {
		//fmt.Println(a)
		p.sum += p.blinkRecurs(num, a)
	}
}

func (p *Pluto) blinkRecurs(depth, num int) (result int) {

	if _, ok := p.cache[depth]; !ok {
		p.cache[depth] = make(map[int]int)
	}
	if val, ok := p.cache[depth][num]; ok {
		return val
	}
	if depth <= 0 {
		return 1
	}
	if num == 0 {
		tmp := p.blinkRecurs(depth-1, 1)
		p.cache[depth][num] = tmp
		return tmp
	}
	str := strconv.Itoa(num)
	if len(str)%2 == 0 {
		a, err := strconv.Atoi(str[:(len(str) / 2)])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(str[(len(str) / 2):])
		if err != nil {
			panic(err)
		}
		c, d := p.blinkRecurs(depth-1, a), p.blinkRecurs(depth-1, b)
		//p.cache[depth-1][a] = c
		//p.cache[depth-1][b] = d
		return c + d
	}
	tmp := p.blinkRecurs(depth-1, num*2024)
	p.cache[depth][num] = tmp
	return tmp
}

func main() {
	pluto := Pluto{}
	pluto.load("day11/input.txt")
	//start := time.Now()
	for range 25 {
		pluto.blink()
	}
	//duration := time.Since(start)
	//fmt.Println("Recurse 1 took: ", duration.Milliseconds(), " ms")
	fmt.Println("Part 1:", len(pluto.pebbles))

	pluto = Pluto{}
	pluto.load("day11/input.txt")
	//start = time.Now()
	pluto.blink2(75)
	//duration = time.Since(start)
	//fmt.Println("Recurse 1 took: ", duration.Milliseconds(), " ms")
	fmt.Println("Part 2:", pluto.sum)

}
