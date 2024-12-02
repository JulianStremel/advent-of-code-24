package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func convertStringToIntSlice(a []string) []int {
	var tmp []int
	for _, number := range a {
		num, err := strconv.ParseInt(number, 0, 0)
		if err != nil {
			panic("we do not need Errorhandling around here ;)")
		}
		tmp = append(tmp, int(num))
	}
	return tmp
}

func testForSafe(a []int) bool {
	if a[0] > a[1] {
		for index, _ := range a {
			if index > 0 {
				if a[index-1]-a[index] < 1 || a[index-1]-a[index] > 3 {
					return false
				}
			}
		}
		return true
	}
	if a[0] < a[1] {
		for index, _ := range a {
			if index > 0 {
				if a[index]-a[index-1] < 1 || a[index]-a[index-1] > 3 {
					return false
				}
			}
		}
		return true
	}
	return false
}

func increasing(a []int) bool {
	for x := range a {
		if x == 0 {
			continue
		}
		if a[x]-a[x-1] > 3 || a[x]-a[x-1] < 1 {
			return false
		}
	}
	return true
}

func decreasing(a []int) bool {
	for x := range a {
		if x == 0 {
			continue
		}
		if a[x-1]-a[x] > 3 || a[x-1]-a[x] < 1 {
			return false
		}
	}
	return true
}

func testForSafe2(a []int) bool {
	if increasing(a) || decreasing(a) {
		return true
	}
	for x := range a {
		var tmp = slices.Clone(a)
		tmp = slices.Delete(tmp, x, x+1)
		if increasing(tmp) || decreasing(tmp) {
			return true
		}
	}
	return false
}

func main() {
	var safe_reports int = 0
	var safe_reports2 int = 0

	var fileLines []string

	readFile, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	// Part 1
	for _, line := range fileLines {
		var tmp = strings.Split(line, " ")
		if testForSafe(convertStringToIntSlice(tmp)) {
			safe_reports += 1
		}
	}

	fmt.Printf("Part1: %d \n", safe_reports)

	// Part 2
	for _, line := range fileLines {
		var tmp1 = strings.Split(line, " ")
		if testForSafe2(convertStringToIntSlice(tmp1)) {
			safe_reports2 += 1
		}
	}

	fmt.Printf("Part2: %d \n", safe_reports2)
}
