package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"
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

func solve(maze [][]bool, guard Guard) map[Pos]bool {
	m := make(map[Pos]bool)

	for {
		if guard.X < 0 || guard.X >= len(maze[0]) || guard.Y < 0 || guard.Y >= len(maze) {
			return m
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

func solve2Multi(maze [][]bool, guard Guard) int {
	var ret atomic.Int64
	wg := sync.WaitGroup{}
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			wg.Add(1)
			mazeCopy := copyMaze(maze)
			guardCopy := guard
			if i != guardCopy.Y || j != guardCopy.X {
				mazeCopy[i][j] = false
			}
			go func(mazeCopy [][]bool, guardCopy Guard) {
				m := make(map[Pos]bool)
				for {
					if guardCopy.X < 0 || guardCopy.X >= len(mazeCopy[0]) || guardCopy.Y < 0 || guardCopy.Y >= len(mazeCopy) {
						break
					}
					m[Pos(guardCopy)] = true
					switch guardCopy.Dir {
					case up:
						if check(mazeCopy, guardCopy.X, guardCopy.Y-1) {
							guardCopy.Y--
						} else {
							guardCopy.Dir = turnRight(guardCopy.Dir)
						}
					case right:
						if check(mazeCopy, guardCopy.X+1, guardCopy.Y) {
							guardCopy.X++
						} else {
							guardCopy.Dir = turnRight(guardCopy.Dir)
						}
					case down:
						if check(mazeCopy, guardCopy.X, guardCopy.Y+1) {
							guardCopy.Y++
						} else {
							guardCopy.Dir = turnRight(guardCopy.Dir)
						}
					case left:
						if check(mazeCopy, guardCopy.X-1, guardCopy.Y) {
							guardCopy.X--
						} else {
							guardCopy.Dir = turnRight(guardCopy.Dir)
						}
					}
					if m[Pos(guardCopy)] {
						ret.Add(1)
						break
					}
				}
			}(mazeCopy, guardCopy)
		}
	}
	return int(ret.Load())
}

func solve3(maze [][]bool, guard Guard, path map[Pos]bool) int {
	mazeCopy := copyMaze(maze)
	guardCopy := guard
	ret := 0
	delete(path, Pos(guard))
	for k, _ := range path {
		m := make(map[Pos]bool)
		maze = copyMaze(mazeCopy)
		guard = guardCopy
		maze[k.Y][k.X] = false

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
	return ret
}

func solve3Multi(maze [][]bool, guard Guard, path map[Pos]bool) int {
	var ret atomic.Int64
	wg := sync.WaitGroup{}
	delete(path, Pos(guard))
	for k := range path {
		copyMaze := copyMaze(maze)
		guardCopy := guard
		copyMaze[k.Y][k.X] = false
		wg.Add(1)
		go func(maze [][]bool, guard Guard) {
			m := make(map[Pos]bool)
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
					ret.Add(1)
					break
				}
			}
			wg.Done()
		}(copyMaze, guardCopy)
	}
	wg.Wait()
	return int(ret.Load())
}

func main() {
	maze, guard := loadData("data.txt")
	path := solve(maze, guard)
	fmt.Println(len(path))
	t0 := time.Now()
	fmt.Println(solve2(maze, guard))
	fmt.Printf("Bruteforce: %fs\n", time.Since(t0).Seconds())
	//t0 = time.Now()
	//fmt.Println(solve2Multi(maze, guard))
	//fmt.Printf("Bruteforce Multi: %fs\n", time.Since(t0).Seconds())
	t0 = time.Now()
	fmt.Println(solve3(maze, guard, path))
	fmt.Printf("Smarter: %fs\n", time.Since(t0).Seconds())
	t0 = time.Now()
	fmt.Println(solve3Multi(maze, guard, path))
	fmt.Printf("Smarter Multi: %fs\n", time.Since(t0).Seconds())
}
