package main

// regex mul\(\d{1,3},\d{1,3}\)

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func executeMul(s string) int {
	var sum = 0
	// prepare regex
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}

	// use regex to find all mul instructions
	var multi = r.FindAllString(s, -1)

	// parse found mul instructions into ints and perform multiplication
	for _, mul := range multi {
		var tmp_str = strings.ReplaceAll(mul, "mul(", "")
		tmp_str = strings.ReplaceAll(tmp_str, ")", "")
		var temp = strings.Split(tmp_str, ",")
		int1, err := strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		int2, err := strconv.Atoi(temp[1])
		if err != nil {
			panic(err)
		}
		sum += int1 * int2
	}
	return sum
}

func main() {
	var sum = 0
	var file_bytes []byte
	var file_string string

	// read file into string
	file_bytes, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	file_string = string(file_bytes)

	fmt.Printf("Part 1: %d\n", executeMul(file_string))

	var do = true
	var remaining = file_string
	for {
		var tmp []string
		if do {
			tmp = strings.SplitN(remaining, "don't()", 2)
			if len(tmp) < 2 {
				if len(tmp) == 1 {
					sum += executeMul(tmp[0])
				}
				break
			}
			sum += executeMul(tmp[0])
			remaining = tmp[1]
			do = false
		} else {
			tmp = strings.SplitN(remaining, "do()", 2)
			if len(tmp) < 2 {
				if len(tmp) == 1 {
					sum += executeMul(tmp[0])
				}
				break
			}
			remaining = tmp[1]
			do = true
		}
	}
	fmt.Printf("Part 2: %d", sum)
}
