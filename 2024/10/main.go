package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

type Pair struct {
    X int
	Y int
}

var found map[Pair]bool
var found2 map[Pair]int

func loadData(filepath string) [][]int {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var m [][]int

    scanner := bufio.NewScanner(file)
    
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, c := range line {
			nbr, _ := strconv.Atoi(string(c))
			row = append(row, nbr)
		}
		m = append(m, row)
	}

    return m
}

func findReachable(m [][]int, x, y, prev int) {
	if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
		return
	}
	
	if m[x][y] == prev + 1 {
		findReachable(m, x-1, y, m[x][y])
		findReachable(m, x+1, y, m[x][y])
		findReachable(m, x, y-1, m[x][y])
		findReachable(m, x, y+1, m[x][y])
	} else {
		return
	}

	if m[x][y] == 9 {
		found[Pair{x, y}] = true
		return
	}
}

func findReachable2(m [][]int, x, y, prev int) {
	if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
		return
	}
	
	if m[x][y] == prev + 1 {
		findReachable2(m, x-1, y, m[x][y])
		findReachable2(m, x+1, y, m[x][y])
		findReachable2(m, x, y-1, m[x][y])
		findReachable2(m, x, y+1, m[x][y])
	} else {
		return
	}

	if m[x][y] == 9 {
		found2[Pair{x, y}]++
		return
	}
}

func solve(m [][]int) int {
	ret := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				findReachable(m, i, j, -1)
				ret += len(found)
				found = make(map[Pair]bool)
			}
		}
	}

    return ret
}

func solve2(m [][]int) int {
	ret := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				findReachable2(m, i, j, -1)
				for _, v := range found2 {
					ret += v
				}
				found2 = make(map[Pair]int)
			}
		}
	}

    return ret
}


func main() {
	found = make(map[Pair]bool)
	found2 = make(map[Pair]int)
    file := "data.txt"
    m := loadData(file)
    fmt.Println(solve(m))
    fmt.Println(solve2(m))
}
