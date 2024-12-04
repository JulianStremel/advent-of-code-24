package main

// regex mul\(\d{1,3},\d{1,3}\)

import (
	"bufio"
	"fmt"
	"os"
)

func checkXmas(data []byte) bool {
	var str = string(data)
	return str == "XMAS" || str == "SAMX"
}

func horizontal(row int, col int, array [][]byte) []byte {
	if row < 0 || row >= len(array) {
		return nil
	}
	if col < 0 || col+3 >= len(array[row]) {
		return nil
	}
	return array[row][col : col+4]
}

func vertical(row int, col int, array [][]byte) []byte {
	if col < 0 || col >= len(array[row]) {
		return nil
	}
	if row < 0 || row+3 >= len(array) {
		return nil
	}
	return []byte{
		array[row][col],
		array[row+1][col],
		array[row+2][col],
		array[row+3][col],
	}
}

func diagLtoR(row int, col int, array [][]byte) []byte {
	if row < 0 || row+3 >= len(array) {
		return nil
	}
	if col < 0 || col+3 >= len(array[row]) {
		return nil
	}
	return []byte{
		array[row][col],
		array[row+1][col+1],
		array[row+2][col+2],
		array[row+3][col+3],
	}
}

func diagRtoL(row int, col int, array [][]byte) []byte {
	if row < 0 || row+3 >= len(array) {
		return nil
	}
	if col < 0 || col+3 >= len(array[row]) {
		return nil
	}
	return []byte{
		array[row][col+3],
		array[row+1][col+2],
		array[row+2][col+1],
		array[row+3][col],
	}
}

func xMas(row int, col int, array [][]byte) bool {
	if len(array)-3 < row || row < 0 {
		return false
	}
	if len(array[row])-3 < col || col < 0 {
		return false
	}

	var tmp [][]byte = make([][]byte, 2)
	tmp[0] = []byte{
		array[row][col],
		array[row+1][col+1],
		array[row+2][col+2]}

	tmp[1] = []byte{
		array[row+2][col],
		array[row+1][col+1],
		array[row][col+2]}

	var str = string(tmp[0])
	if !(str == "MAS" || str == "SAM") {
		return false
	}
	str = string(tmp[1])
	if !(str == "MAS" || str == "SAM") {
		return false
	}

	return true
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

	var count1 = 0
	var count2 = 0
	for row := range file_bytes {
		for column := range file_bytes[row] {
			if checkXmas(horizontal(row, column, file_bytes)) {
				count1 += 1
			}
			if checkXmas(vertical(row, column, file_bytes)) {
				count1 += 1
			}
			if checkXmas(diagLtoR(row, column, file_bytes)) {
				count1 += 1
			}
			if checkXmas(diagRtoL(row, column, file_bytes)) {
				count1 += 1
			}
			// Part 2
			if xMas(row, column, file_bytes) {
				count2 += 1
			}
		}

	}

	fmt.Printf("Part 1: %d\n", count1)

	fmt.Printf("Part 2: %d\n", count2)

}
