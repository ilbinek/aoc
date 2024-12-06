package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type Pos struct {
	X   int
	Y   int
	Dir direction
}

type Guard struct {
	X   int
	Y   int
	Dir direction
}

func loadData(filepath string) ([][]bool, Guard) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([][]bool, 0)
	ret2 := Guard{0, 0, right}
	for scanner.Scan() {
		row := make([]bool, 0)
		for _, c := range scanner.Text() {
			if c == '#' {
				row = append(row, false)
			} else if c == '^' {
				ret2.X = len(row)
				ret2.Y = len(ret)
				ret2.Dir = up
				row = append(row, true)
			} else {
				row = append(row, true)
			}
		}
		ret = append(ret, row)
	}

	return ret, ret2
}

func turnRight(dir direction) direction {
	return (dir + 1) % 4
}

func check(maze [][]bool, x int, y int) bool {
	if x < 0 || x >= len(maze[0]) || y < 0 || y >= len(maze) {
		return true
	}
	return maze[y][x]
}

func solve(maze [][]bool, guard Guard) int {
	m := make(map[Pos]bool)

	for {
		if guard.X < 0 || guard.X >= len(maze[0]) || guard.Y < 0 || guard.Y >= len(maze) {
			return len(m)
		}
		m[Pos{guard.X, guard.Y, up}] = true
		switch guard.Dir {
		case up:
			if check(maze, guard.X, guard.Y-1) {
				guard.Y--
			} else {
				guard.Dir = turnRight(guard.Dir)
			}
		case right:
			if check(maze, guard.X+1, guard.Y) {
				guard.X++
			} else {
				guard.Dir = turnRight(guard.Dir)
			}
		case down:
			if check(maze, guard.X, guard.Y+1) {
				guard.Y++
			} else {
				guard.Dir = turnRight(guard.Dir)
			}
		case left:
			if check(maze, guard.X-1, guard.Y) {
				guard.X--
			} else {
				guard.Dir = turnRight(guard.Dir)
			}
		}
	}
}

func copyMaze(maze [][]bool) [][]bool {
	ret := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		ret[i] = make([]bool, len(maze[i]))
		copy(ret[i], maze[i])
	}
	return ret
}

func solve2(maze [][]bool, guard Guard) int {
	mazeCopy := copyMaze(maze)
	guardCopy := guard
	ret := 0
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {

			m := make(map[Pos]bool)
			maze = copyMaze(mazeCopy)
			guard = guardCopy
			if maze[i][j] && (i != guard.Y || j != guard.X) {
				maze[i][j] = false
			}
			for {
				if guard.X < 0 || guard.X >= len(maze[0]) || guard.Y < 0 || guard.Y >= len(maze) {
					break
				}
				m[Pos(guard)] = true
				switch guard.Dir {
				case up:
					if check(maze, guard.X, guard.Y-1) {
						guard.Y--
					} else {
						guard.Dir = turnRight(guard.Dir)
					}
				case right:
					if check(maze, guard.X+1, guard.Y) {
						guard.X++
					} else {
						guard.Dir = turnRight(guard.Dir)
					}
				case down:
					if check(maze, guard.X, guard.Y+1) {
						guard.Y++
					} else {
						guard.Dir = turnRight(guard.Dir)
					}
				case left:
					if check(maze, guard.X-1, guard.Y) {
						guard.X--
					} else {
						guard.Dir = turnRight(guard.Dir)
					}
				}
				if m[Pos(guard)] {
					ret++
					break
				}
			}
		}
	}
	return ret
}

func main() {
	maze, guard := loadData("data.txt")

	fmt.Println(solve(maze, guard))
	fmt.Println(solve2(maze, guard))
}
