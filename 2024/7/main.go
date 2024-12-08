package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Equation struct {
    Result  int
    Numbers []int
}

func loadData(filepath string) []Equation {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    ret := make([]Equation, 0)
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), ":")
        result, err := strconv.Atoi(parts[0])
        if err != nil {
            log.Fatal(err)
        }

        numbers := make([]int, 0)
        for _, num := range strings.Split(strings.Trim(parts[1], " "), " ") {
            n, err := strconv.Atoi(num)
            if err != nil {
                log.Fatal(err)
            }
            numbers = append(numbers, n)
        }

        ret = append(ret, Equation{Result: result, Numbers: numbers})
    }

    return ret
}

func canCombine(eq Equation) bool {
    nbrOfOperations := len(eq.Numbers) - 1
    options := 1 << uint(nbrOfOperations)
    for i := 0; i < options; i++ {
        sum := eq.Numbers[0]
        for j := 0; j < nbrOfOperations; j++ {
            if i&(1<<uint(j)) != 0 {
                sum += eq.Numbers[j+1]
            } else {
                sum *= eq.Numbers[j+1]
            }
        }
        if sum == eq.Result {
            return true
        }
    }
    return false
}

func solve(equations []Equation) int {
    ret := 0
    for _, eq := range equations {
        if canCombine(eq) {
            ret += eq.Result
        }
    }
    return ret
}

func combineNumbers(a, b int) int {
    aStr := strconv.Itoa(a)
    bStr := strconv.Itoa(b)
    ret, err := strconv.Atoi(aStr + bStr)
    if err != nil {
        log.Fatal(err)
    }
    return ret
}

func check(equation Equation) bool {
    if len(equation.Numbers) == 1 {
        return equation.Numbers[0] == equation.Result
    }
    nbrs := []int{equation.Numbers[0] + equation.Numbers[1]};
    nbrs = append(nbrs, equation.Numbers[2:]...)
    e := Equation{equation.Result, nbrs}
    b := check(e)
    if b {
        return true
    }

    nbrs = []int{equation.Numbers[0] * equation.Numbers[1]};
    nbrs = append(nbrs, equation.Numbers[2:]...)
    e = Equation{equation.Result, nbrs}
    b = check(e)
    if b {
        return true
    }
    
    nbrs = []int{combineNumbers(equation.Numbers[0], equation.Numbers[1])};
    nbrs = append(nbrs, equation.Numbers[2:]...)
    e = Equation{equation.Result, nbrs}
    b = check(e)
    return b
}

func solve22(equations []Equation) int {
    wg := sync.WaitGroup{}
    ret := atomic.Int64{}
    for _, eq := range equations {
        wg.Add(1)
        go func (eq Equation) {
            if check(eq) {
                ret.Add(int64(eq.Result))
            }
            wg.Done()
        }(eq)
    }
    wg.Wait()
    return int(ret.Load())
}

func main() {
    equa := loadData("data.txt")
    fmt.Println(solve(equa))
    t0 := time.Now()
    fmt.Println(solve22(equa))
    fmt.Printf("Execution time: %vms\n", time.Since(t0).Milliseconds())
}
