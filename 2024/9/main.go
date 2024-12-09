package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

type AOCFile struct {
    Nbr   int
    Count int
    Start int
}

func loadData(filepath string) []int {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var m []int

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    cur := 0
    free := false
    for _, c := range scanner.Text() {
        nbr, err := strconv.Atoi(string(c))
        if err != nil {
            log.Fatal(err)
        }
        if free {
            free = false
            for i := 0; i < nbr; i++ {
                m = append(m, -1)
            }
        } else {
            free = true
            for i := 0; i < nbr; i++ {
                m = append(m, cur)
            }
            cur++
        }
    }

    return m
}

func solve(m []int) int {
    prevLeft := 0
    for {
        // most left .
        left := 0
        for i := 0; i < len(m); i++ {
            if m[i] == -1 {
                left = i
                break
            }
        }

        // right most not .
        right := 0
        for i := len(m) - 1; i >= 0; i-- {
            if m[i] != -1 {
                right = i
                break
            }
        }

        if left >= right {
            break
        }

        if left == prevLeft {
            break
        }

        // swap
        m[left], m[right] = m[right], m[left]
        prevLeft = left
    }

    ret := 0

    for i := 0; i < len(m); i++ {
        if m[i] == -1 {
            break
        }

        ret += i * m[i]
    }

    return ret
}

func solve2(m []int) int {
    pairs := []AOCFile{}
    // Parse into pairs
    pairNbr := 0
    count := 0
    for i := 0; i < len(m); i++ {
        if pairNbr == m[i] {
            count++
        } else {
            pairs = append(pairs, AOCFile{pairNbr, count, i - count})
            count = 1
            pairNbr = m[i]
        }
    }
    pairs = append(pairs, AOCFile{pairNbr, count, len(m) - count})

    for i := len(pairs) - 1; i >= 0; i-- {
        if pairs[i].Nbr != -1 {
            // Go from left and find empty where it fits
            emptyCount := 0
            for j := 0; j < len(m); j++ {
                if j > pairs[i].Start {
                    break
                }
                if m[j] == -1 {
                    emptyCount++
                    // check if it fits
                    if emptyCount == pairs[i].Count {
                        for k := j - pairs[i].Count + 1; k <= j; k++ {
                            m[k] = pairs[i].Nbr
                        }
                        for k := 0; k < pairs[i].Count; k++ {
                            m[k+pairs[i].Start] = -1
                        }
                        break
                    }
                } else {
                    emptyCount = 0
                }
            }
        }
    }

    ret := 0

    for i := 0; i < len(m); i++ {
        if m[i] == -1 {
            continue
        }

        ret += i * m[i]
    }
    return ret
}

func main() {
    file := "data.txt"
    m := loadData(file)
    fmt.Println(solve(m))
    m = loadData(file)
    fmt.Println(solve2(m))
}
