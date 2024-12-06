package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

func loadData(filename string) [][]int {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    ret := make([][]int, 0)

    scanner := bufio.NewScanner(file)
    cnt := 0
    for scanner.Scan() {
        ret = append(ret, make([]int, 0))
        line := scanner.Text()
        nbrs := strings.Fields(line)
        for _, nbr := range nbrs {
            n, err := strconv.Atoi(nbr)
            if err != nil {
                log.Fatal(err)
            }
            ret[cnt] = append(ret[cnt], n)
        }
        cnt++
    }

    return ret
}

func solveOne(data [][]int) int {
    ret := 0
    for i := 0; i < len(data); i++ {
        if isSafe(data[i]) {
            ret++
        }
    }

    return ret
}

func isSafe(data []int) bool {
    ascending := data[0] < data[1]
    for j := 0; j < len(data)-1; j++ {
        // Check if the current and next element have max distance of three
        if math.Abs(float64(data[j]-data[j+1])) > 3 {
            return false
        }
        if ascending && data[j] >= data[j+1] {
            return false
        } else if !ascending && data[j] <= data[j+1] {
            return false
        }
    }

    return true
}

func solveTwo(data [][]int) int {
    ret := 0
    for i := 0; i < len(data); i++ {
        if isSafe(data[i]) {
            ret++
        } else {
            // Try to run it by omiting one element
            for j := 0; j < len(data[i]); j++ {
                tmp := make([]int, 0)
                for k := 0; k < len(data[i]); k++ {
                    if k != j {
                        tmp = append(tmp, data[i][k])
                    }
                }
                if isSafe(tmp) {
                    ret++
                    break
                }
            }
        }
    }

    return ret
}

func main() {
    data := loadData("data.dat")
    fmt.Println(solveOne(data))
    fmt.Println(solveTwo(data))
}
