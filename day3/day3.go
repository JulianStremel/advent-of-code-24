package main

// regex mul\(\d{1,3},\d{1,3}\)

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var sum = 0
	var file_bytes []byte
	var file_string string

	file_bytes, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	file_string = string(file_bytes)

	//fmt.Println(file_string)

	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}
	var multi = r.FindAllString(file_string, -1)
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
	fmt.Println(sum)

}
