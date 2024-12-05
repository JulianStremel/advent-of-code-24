package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var combinations [][]int
	var checks [][]int

	readFile, err := os.Open("day5/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var tmp []string
	for fileScanner.Scan() {
		tmp = strings.Split(fileScanner.Text(), "|")
		int1, err := strconv.Atoi(tmp[0])
		if err != nil {
			panic(err)
		}
		int2, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		combinations = append(combinations, []int{int1, int2})
	}
	readFile.Close()

	readFile, err = os.Open("day5/check.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	tmp = nil
	for fileScanner.Scan() {
		var temp []int
		tmp = strings.Split(fileScanner.Text(), ",")
		for _, num := range tmp {
			number, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			temp = append(temp, number)
		}
		checks = append(checks, temp)
	}
	readFile.Close()

	// Part 1
	// pages -> list of pages
	var stop = false
	var sum = 0
	for _, pages := range checks {
		if stop {
			stop = false
		}
		// page -> single page (int)
		for b, page := range pages {
			if stop {
				break
			}
			// pair -> (int,int)
			for _, pair := range combinations {
				if page == pair[0] {
					if slices.Contains(pages[0:b], pair[1]) {
						stop = true
						fmt.Printf("%d does include %d\n", pages[0:b], pair[1])
						break
					}
				}
			}
		}
		if !stop {
			sum += pages[(len(pages) / 2)]
		}
	}
	fmt.Printf("Part 1: %d", sum)
}
