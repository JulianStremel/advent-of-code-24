package main

// regex mul\(\d{1,3},\d{1,3}\)

import (
	"bufio"
	"fmt"
	"os"
)

func checkXmas(data []byte) bool {
	return string(data[:]) == "XMAS" || string(data[:]) == "SAMX"
}

func horizontal(row int, col int, array [][]byte) []byte {
	var tmp []byte = make([]byte, 0)
	if len(array) < row || row < 0 {
		return tmp
	}
	if len(array[row])-4 < col || col < 0 {
		return tmp
	}
	return array[row][col : col+4]
}

func vertical(row int, col int, array [][]byte) []byte {
	var tmp []byte = make([]byte, 0)
	if len(array)-4 < row || row < 0 {
		return tmp
	}
	if len(array[row]) < col || col < 0 {
		return tmp
	}
	tmp = append(tmp, array[row][col])
	tmp = append(tmp, array[row+1][col])
	tmp = append(tmp, array[row+2][col])
	tmp = append(tmp, array[row+3][col])
	return tmp
}

func diagLtoR(row int, col int, array [][]byte) []byte {
	var tmp []byte = make([]byte, 0)
	if len(array)-4 < row || row < 0 {
		return tmp
	}
	if len(array[row])-4 < col || col < 0 {
		return tmp
	}
	tmp = append(tmp, array[row][col])
	tmp = append(tmp, array[row+1][col+1])
	tmp = append(tmp, array[row+2][col+2])
	tmp = append(tmp, array[row+3][col+3])
	return tmp
}

func diagRtoL(row int, col int, array [][]byte) []byte {
	var tmp []byte = make([]byte, 0)
	if len(array)-4 < row || row < 0 {
		return tmp
	}
	if len(array[row])-4 < col || col < 0 {
		return tmp
	}
	tmp = append(tmp, array[row][col+3])
	tmp = append(tmp, array[row+1][col+2])
	tmp = append(tmp, array[row+2][col+1])
	tmp = append(tmp, array[row+3][col])
	return tmp
}

func main() {
	var file_bytes [][]byte
	// read file into string
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		file_bytes = append(file_bytes, fileScanner.Bytes())
	}

	var count = 0
	for row, row_bytes := range file_bytes {
		for column, _ := range row_bytes {
			if checkXmas(horizontal(row, column, file_bytes)) {
				count += 1
			}
			if checkXmas(vertical(row, column, file_bytes)) {
				count += 1
			}
			if checkXmas(diagLtoR(row, column, file_bytes)) {
				count += 1
			}
			if checkXmas(diagRtoL(row, column, file_bytes)) {
				count += 1
			}
		}

	}
	// 2276 too low
	fmt.Printf("Part 1: %d\n", count)
}
