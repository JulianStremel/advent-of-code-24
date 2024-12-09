package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Storage struct {
	input         []int
	decrompressed []int
	checksum      int
	chunks_free   []Chunk
	chunks_file   []Chunk
}

type Chunk struct {
	id    int
	index int
	size  int
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
	var virtual_index = 0
	for index, num := range s.input {
		if index%2 == 0 {
			for range num {
				s.decrompressed = append(s.decrompressed, id)
			}
			s.chunks_file = append(s.chunks_file, Chunk{id, virtual_index, num})
			virtual_index += num
			id++
		} else {
			for range num {
				s.decrompressed = append(s.decrompressed, -1)
			}
			s.chunks_free = append(s.chunks_free, Chunk{-1, virtual_index, num})
			virtual_index += num
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

func (s *Storage) compactNonFragment() {
	for a, _ := range s.chunks_file {
		chunk_file := s.chunks_file[len(s.chunks_file)-1-a]
		for b, chunk_free := range s.chunks_free {
			if chunk_file.size <= chunk_free.size {
				// switch file chunk into free space
				for c := range chunk_file.size {
					s.decrompressed[chunk_free.index+c] = s.decrompressed[chunk_file.index+c]
					s.decrompressed[chunk_file.index+c] = -1
				}
				if chunk_file.size < chunk_free.size {
					s.chunks_free = append(s.chunks_free, Chunk{chunk_free.index + chunk_file.size, chunk_free.size - chunk_file.size, -1})
				}
				slices.Delete(s.chunks_free, b, b+1)
			}
		}
	}
}

func (s *Storage) updateChecksum() {
	for a, num := range s.decrompressed {
		if num > 0 {
			s.checksum += a * num
		}
	}
}

// 2333133121414131402
// 00992111777.44.333....5555.6666.....8888..
// -> 2858

func main() {
	storage := Storage{}
	storage.read("day9/input_test.txt")
	storage.decompress()
	fmt.Println(storage.decrompressed)
	storage.compactNonFragment()
	storage.updateChecksum()
	fmt.Println(storage.decrompressed)
	fmt.Println(storage.checksum)
}
