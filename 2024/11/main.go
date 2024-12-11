package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var precalc = make([]int, 0)

func loadData(filepath string) []int {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var m []int

    scanner := bufio.NewScanner(file)
    
	scanner.Scan()
	for _, c := range strings.Split(scanner.Text(), " ") {
		nbr, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}

		m = append(m, nbr)
	}
	
	
    return m
}

func hasEvenDigits(n int) (int, int, bool) {
	nbr := strconv.Itoa(n)
	if len(nbr) % 2 == 0 {
		nbr1, _ := strconv.Atoi(nbr[:len(nbr)/2])
		nbr2, _ := strconv.Atoi(nbr[len(nbr)/2:])
		return nbr1, nbr2, true
	}
	return 0, 0, false
}

func solve(m []int, n int) int {
	mm := make(map[int]int)
	for _, v := range m {
		mm[v]++
	}

	for i := 0; i < n; i++ {
		tmp := make(map[int]int)
		for v := range mm {
			// rule 1
			if v == 0 {
				tmp[1] += mm[0]
				delete(mm, 0)
				continue
			}

			// rule 2
			if dgt1, dgt2, ok := hasEvenDigits(v); ok {
				tmp[dgt1] += mm[v];
				tmp[dgt2] += mm[v];
				delete(mm, v)
				continue
			}

			// rule 3
			tmp[v * 2024] += mm[v]
			delete(mm, v)
		}
		mm = tmp
	}

	ret := 0
	for _, v := range mm {
		ret += v
	}

	return ret
}

func main() {
    m := loadData("data.txt")
	t0 := time.Now()
	fmt.Println(solve(m, 25))
	fmt.Printf("Execution time 25 iter: %v\n", time.Since(t0))
	t0 = time.Now()
	fmt.Println(solve(m, 75))
	fmt.Printf("Execution time 75 iter: %v\n", time.Since(t0))
}
