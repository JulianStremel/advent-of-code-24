package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countInt(slice []T, x T) int {
	var count = 0
	for _, a := range slice {
		if a == x {
			count += 1
		}
	}
	return count
}

func main() {
	var distance = 0

	var similarity = 0

	var list1 []int
	var list2 []int

	var fileLines []string

	readFile, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		var tmp = strings.Split(line, "   ")
		int1, err1 := strconv.ParseInt(tmp[0], 0, 0)
		if err1 != nil {
			fmt.Println(err1)
		}
		int2, err2 := strconv.ParseInt(tmp[1], 0, 0)
		if err2 != nil {
			fmt.Println(err2)
		}
		list1 = append(list1, int(int1))
		list2 = append(list2, int(int2))
	}

	// Part 1
	slices.Sort(list1)
	slices.Sort(list2)
	if len(list1) != len(list2) {
		fmt.Println("Lists are not equal in length")
	}

	for index, _ := range list1 {
		distance += absInt(list1[index] - list2[index])
	}

	fmt.Printf("Part1 (distance): %d \n", distance)

	// Part 2
	for _, number := range list1 {
		similarity += number * countInt(list2, number)
	}

	fmt.Printf("Part2 (similarity): %d \n", similarity)

}
