package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "slices"
    "strconv"
    "strings"
)

func loadData(path string) [][]int {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    ret := make([][]int, 2)
    ret[0] = make([]int, 0)
    ret[1] = make([]int, 0)

    // read every line and split it by space
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(line)
        n1, err := strconv.Atoi(words[0])
        if err != nil {
            log.Fatal(err)
        }

        n2, err := strconv.Atoi(words[1])
        if err != nil {
            log.Fatal(err)
        }

        ret[0] = append(ret[0], n1)
        ret[1] = append(ret[1], n2)
    }

    return ret
}

func calc(data [][]int) float64 {
    ret := 0.0
    for i := 0; i < len(data[0]); i++ {
        ret += math.Abs(float64(data[0][i] - data[1][i]))
    }

    return ret
}

func calc2(data [][]int) int {
    ret := 0
    for i := 0; i < len(data[0]); i++ {
        ret += data[0][i] * count(data[1], data[0][i])
    }
    return ret
}

func count(data []int, desired int) int {
    ret := 0
    for i := 0; i < len(data); i++ {
        if data[i] == desired {
            ret++
        }
    }

    return ret
}

func main() {
    data := loadData("data.dat")
    slices.Sort(data[0])
    slices.Sort(data[1])

    fmt.Println(int(calc(data)))
    fmt.Println(calc2(data))
}
