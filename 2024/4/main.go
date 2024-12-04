package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var word string = "MAS"

type direction int

const (
	up direction = iota
	upRight
	right
	downRight
	down
	downLeft
	left
	upLeft
)

func loadData(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([]string, 0)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func continueSearch(data []string, x int, y int, dir direction) bool {
	for letter := range word {
		switch dir {
		case up:
			y--
		case upRight:
			x++
			y--
		case right:
			x++
		case downRight:
			x++
			y++
		case down:
			y++
		case downLeft:
			x--
			y++
		case left:
			x--
		case upLeft:
			x--
			y--
		}

		if x < 0 || x >= len(data[0]) || y < 0 || y >= len(data) {
			return false
		}

		if data[y][x] != word[letter] {
			return false
		}
	}

	return true
}

func solve(data []string) int {
	ret := 0
	// find every X and start searching from there
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == 'X' {
				for dir := up; dir <= upLeft; dir++ {
					if continueSearch(data, x, y, dir) {
						ret++
					}
				}
			}
		}
	}

	return ret
}

func check2(data []string, x int, y int) bool {
	if y < 1 || y >= len(data)-1 || x < 1 || x >= len(data[0])-1 {
		return false
	}

	if 	(data[y-1][x-1] == 'M' && data[y+1][x+1] == 'S' ||
		data[y-1][x-1] == 'S' && data[y+1][x+1] == 'M') &&
		(data[y-1][x-1] == 'S' && data[y+1][x+1] == 'M' ||
		data[y-1][x-1] == 'M' && data[y+1][x+1] == 'S') {
			return true
		}

	return false
}

func solve2(data []string) int {
	ret := 0
	// find every X and start searching from there
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == 'A' {
				if check2(data, x, y) {
					ret++
				}
			}
		}
	}

	return ret
}

func main() {
	data := loadData("test.txt")
	fmt.Println(solve(data))
	fmt.Println(solve2(data))
}
