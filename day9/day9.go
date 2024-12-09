package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Storage struct {
	input         []int
	decrompressed []int
	checksum      int
}

func (s *Storage) read(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	s.input = []int{}
	for fileScanner.Scan() {
		tmp, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			panic(err)
		}
		s.input = append(s.input, tmp)
	}
}

func (s *Storage) decompress() {
	var id = 0
	for index, num := range s.input {
		if index%2 == 0 {
			for range num {
				s.decrompressed = append(s.decrompressed, id)
			}
			id++
		} else {
			for range num {
				s.decrompressed = append(s.decrompressed, -1)
			}
		}
	}
}

func (s *Storage) compact() {
	var tail = len(s.decrompressed) - 1
	for a, num := range s.decrompressed {
		if num < 0 && tail > a {
			for s.decrompressed[tail] < 0 {
				tail--
			}
			if tail < a {
				continue
			}
			s.decrompressed[a], s.decrompressed[tail] = s.decrompressed[tail], s.decrompressed[a]
			tail--
		}
		if tail < a {
			break
		}
	}
	s.decrompressed = s.decrompressed[:tail+1]
}

func (s *Storage) updateChecksum() {
	for a, num := range s.decrompressed {
		if num > 0 {
			s.checksum += a * num
		}
	}
}

// 6519155394575 too high

func main() {
	storage := Storage{}
	storage.read("day9/input.txt")
	storage.decompress()
	storage.compact()
	storage.updateChecksum()
	fmt.Println(storage.checksum)
}
