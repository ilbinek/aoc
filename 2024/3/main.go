package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func loadData(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`^\([0-9]{1,3},[0-9]{1,3}\).*`)

	scanner := bufio.NewScanner(file)
	ret := make([]string, 0)
	for scanner.Scan() {
		splt := strings.Split(scanner.Text(), "mul")
		for _, s := range splt {
			if len(ret) == 228 {
				fmt.Println(s)
			}
			if (strings.Contains(s, "when(127")) {
				fmt.Println(s)
			}
			if r.MatchString(s) {
				ret = append(ret, s)
			}
		}
	}

	ret2 := make([][]int, 0)
	indx := 0
	for _, r := range ret {
		fmt.Printf("%d: %s\n", indx, r)
		indx++
		ret2 = append(ret2, make([]int, 0))
		// find first )
		i := strings.Index(r, ")")
		// remove ( and first )
		r = r[1:i]
		nbrs := strings.Split(r, ",")
		for _, nbr := range nbrs {
			n, err := strconv.Atoi(nbr)
			if err != nil {
				log.Fatal(err)
			}
			ret2[len(ret2)-1] = append(ret2[len(ret2)-1], n)
		}
	}
	
	return ret2
}

func solveOne(data [][]int) int {
	ret := 0
	for i := 0; i < len(data); i++ {
		ret += data[i][0] * data[i][1]
	}

	return ret
}

func main() {
	data := loadData("data.dat")
	fmt.Println(solveOne(data))
}
