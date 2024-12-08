package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
    X int
    Y int
}

func loadData(filepath string) [][]rune {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    ret := make([][]rune, 0)
    for scanner.Scan() {
        line := scanner.Text()
        ret = append(ret, []rune(line))        
    }

    return ret
}

func solve(m [][]rune) int {
    anti := make([][]bool, len(m))
    for i := range m {
        anti[i] = make([]bool, len(m[i]))
    }

    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[i]); j++ {
            // Check if it is empty
            if m[i][j] == '.' {
                continue
            }

            antena := m[i][j];

            // Look for the same antena in the map
            for k := 0; k < len(m); k++ {
                for l := 0; l < len(m[k]); l++ {
                    if m[k][l] == antena {
                        if k == i && l == j {
                            continue
                        }
                        dX := k - i
                        dY := l - j
                        // double the distance
                        dX *= 2
                        dY *= 2

                        // Set it as anti
                        if i + dX >= 0 && i + dX < len(m) && j + dY >= 0 && j + dY < len(m[i]) {
                            anti[i + dX][j + dY] = true
                        }
                    }
                }
            }    
        }
    }

    ret := 0
    for i := range anti {
        for j := range anti[i] {
            if anti[i][j] {
                ret++
                fmt.Printf("X")
            } else {
                fmt.Printf(".")
            }
        }
        fmt.Println()
    }
    return ret
}

func findAntenas(m [][]rune, t rune, x, y, dx, dy int) []Pos {
    ret := make([]Pos, 0)
    i := x + dx
    j := y + dy
    for {
        if (i >= 0 && i < len(m) && j >= 0 && j < len(m[i])) {
            if m[i][j] == t {
                ret = append(ret, Pos{i, j})
            }
            i += dx 
            j += dy 
        } else {
            break
        }
    }
    return ret
}

func solve2(m [][]rune) int {
    anti := make([][]bool, len(m))
    for i := range m {
        anti[i] = make([]bool, len(m[i]))
    }

    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[i]); j++ {
            // Check if it is empty
            if m[i][j] == '.' {
                continue
            }

            antena := m[i][j];

            // Look for the same antena in the map
            for k := 0; k < len(m); k++ {
                for l := 0; l < len(m[k]); l++ {
                    if m[k][l] == antena {
                        if k == i && l == j {
                            continue
                        }

                        // Set both as anti
                        anti[i][j] = true
                        anti[k][l] = true

                        dX := k - i
                        dY := l - j

                        // Set it as anti in one direction
                        for {
                            if i + dX >= 0 && i + dX < len(m) && j + dY >= 0 && j + dY < len(m[i]) {
                                anti[i + dX][j + dY] = true
                                dX += k - i
                                dY += l - j
                            } else {
                                break
                            }
                        }

                        // Set it as anti in the other direction
                        dX = i - k
                        dY = j - l
                        for {
                            if k + dX >= 0 && k + dX < len(m) && l + dY >= 0 && l + dY < len(m[i]) {
                                anti[k + dX][l + dY] = true
                                dX += i - k
                                dY += j - l
                            } else {
                                break
                            }
                        }
                        
                    }
                }
            }    
        }
    }

    ret := 0
    for i := range anti {
        for j := range anti[i] {
            if anti[i][j] {
                ret++
                fmt.Printf("X")
            } else {
                fmt.Printf(".")
            }
        }
        fmt.Println()
    }
    return ret
}

func main() {
    m := loadData("data.txt")
    fmt.Println(solve(m))
    fmt.Println(solve2(m))
}
