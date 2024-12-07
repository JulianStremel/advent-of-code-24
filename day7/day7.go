package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

type jobData struct {
	solution int
	data     []int
}

type jobResponse struct {
	solution int
	possible bool
}

func worker1(jobs <-chan jobData, results chan<- jobResponse, progressChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		possible := solvable1(job.solution, job.data)
		results <- jobResponse{solution: job.solution, possible: possible}
		progressChan <- 1 // Update progress
	}
}

func worker2(jobs <-chan jobData, results chan<- jobResponse, progressChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		possible := solvable2(job.solution, job.data)
		results <- jobResponse{solution: job.solution, possible: possible}
		progressChan <- 1 // Update progress
	}
}

func (g *Game) solvePart1() int {

	num_jobs := len(g.input)
	jobs := make(chan jobData, num_jobs)
	results := make(chan jobResponse, num_jobs)
	progressChan := make(chan int, num_jobs)
	numCores := runtime.NumCPU()

	go func() {
		for i, block := range g.input {
			jobs <- jobData{solution: i, data: block}
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go worker1(jobs, results, progressChan, &wg)
	}

	go func() {
		completed := 0
		for progress := range progressChan {
			completed += progress
			percentage := float64(completed) / float64(num_jobs) * 100
			fmt.Printf("\rColculating Part 1: %d/%d (%.2f%%)", completed, num_jobs, percentage)
		}
		fmt.Println()
	}()

	wg.Wait()
	close(results)
	close(progressChan)

	var cnt = 0
	for frame := range results {
		if frame.possible {
			cnt += frame.solution
		}
	}
	return cnt

}

func (g *Game) solvePart2() int {

	num_jobs := len(g.input)
	jobs := make(chan jobData, num_jobs)
	results := make(chan jobResponse, num_jobs)
	progressChan := make(chan int, num_jobs)
	numCores := runtime.NumCPU()

	go func() {
		for i, block := range g.input {
			jobs <- jobData{solution: i, data: block}
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go worker2(jobs, results, progressChan, &wg)
	}

	go func() {
		completed := 0
		for progress := range progressChan {
			completed += progress
			percentage := float64(completed) / float64(num_jobs) * 100
			fmt.Printf("\rColculating Part 2: %d/%d (%.2f%%)", completed, num_jobs, percentage)
		}
		fmt.Println()
	}()

	wg.Wait()
	close(results)
	close(progressChan)

	var cnt = 0
	for frame := range results {
		if frame.possible {
			cnt += frame.solution
		}
	}
	return cnt

}

func main() {
	var game = Game{}
	game.init("day7/input.txt")
	part1 := game.solvePart1()
	part2 := game.solvePart2()
	fmt.Printf("Part 1\n: %d", part1)
	fmt.Printf("Part 2\n: %d", part2)
}
