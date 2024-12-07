package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	input map[int][]int
	sum   int
}

func sum(in []int) int {
	var ret = 0
	for _, a := range in {
		ret += a
	}
	return ret
}

func mult(in []int) int {
	var ret = 0
	for a, b := range in {
		if a == 0 {
			ret = b
			continue
		}
		ret = ret * b
	}
	return ret
}

func (g *Game) init(path string) {
	g.input = make(map[int][]int)
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var ints_i []int
		line := strings.Split(fileScanner.Text(), ": ")
		solution, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}

		ints := strings.Split(line[1], " ")
		for _, num := range ints {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			ints_i = append(ints_i, i)
		}
		g.input[solution] = append(g.input[solution], ints_i...)
	}
}

func solvable1(solution int, numbers []int) bool {
	// for 4 elements we would have 3 slots of computations
	var slots = len(numbers) - 1
	var possible = int(math.Pow(2, float64(slots)))
	var operations [][]string
	for a := range possible {
		var compute []string
		current := a
		for range slots {
			if current%2 == 0 {
				compute = append(compute, "+")
			} else {
				compute = append(compute, "*")
			}
			current /= 2
		}
		operations = append(operations, compute)
	}

	for _, operation := range operations {
		var sol = numbers[0]
		for b, op := range operation {
			if op == "+" {
				sol += numbers[b+1]
			} else {
				sol *= numbers[b+1]
			}
		}
		if sol == solution {
			return true
		}
	}
	return false
}

func solvable2(solution int, numbers []int) bool {
	// for 4 elements we would have 3 slots of computations
	var slots = len(numbers) - 1
	var possible = int(math.Pow(3, float64(slots)))
	var operations [][]string
	for a := range possible {
		var compute []string
		current := a
		for range slots {
			switch current % 3 {
			case 0:
				compute = append(compute, "+")
			case 1:
				compute = append(compute, "*")
			case 2:
				compute = append(compute, "|")
			}
			current /= 3
		}
		operations = append(operations, compute)
	}

	for _, operation := range operations {
		var sol = numbers[0]
		for b, op := range operation {
			switch op {
			case "+":
				sol += numbers[b+1]
			case "*":
				sol *= numbers[b+1]
			case "|":
				s1 := strconv.FormatInt(int64(sol), 10)
				s2 := strconv.FormatInt(int64(numbers[b+1]), 10)
				in, err := strconv.Atoi(s1 + s2)
				if err != nil {
					panic(err)
				}
				sol = in
			}
		}
		if sol == solution {
			return true
		}
	}
	return false
}

func main() {
	var game = Game{}
	game.init("day7/input.txt")
	var cnt = 0
	for a, b := range game.input {
		if solvable1(a, b) {
			cnt += a
		}
	}
	fmt.Println(cnt)
	cnt = 0
	for a, b := range game.input {
		if solvable2(a, b) {
			cnt += a
		}
	}
	fmt.Println(cnt)
}
