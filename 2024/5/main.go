package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "slices"
    "strconv"
    "strings"
)

type Rule struct {
    Before int
    After  int
}

func loadData(filepath string) ([]Rule, [][]int) {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    ret := make([]Rule, 0)
    ret2 := make([][]int, 0)
    rules := true
    for scanner.Scan() {
        if scanner.Text() == "" {
            rules = false
            continue
        }

        if rules {
            parts := strings.Split(scanner.Text(), "|")
            before, _ := strconv.Atoi(parts[0])
            after, _ := strconv.Atoi(parts[1])
            ret = append(ret, Rule{Before: before, After: after})
        } else {
            parts := strings.Split(scanner.Text(), ",")
            row := make([]int, 0)
            for _, part := range parts {
                num, _ := strconv.Atoi(part)
                row = append(row, num)
            }
            ret2 = append(ret2, row)
        }
    }
    return ret, ret2
}

func isOk(rules []Rule, row []int) bool {
    passed := make(map[int]bool)
    // Filter rules
    filteredRules := make([]Rule, 0)
    for _, rule := range rules {
        if slices.Contains(row, rule.After) && slices.Contains(row, rule.Before) {
            filteredRules = append(filteredRules, rule)
        }
    }
    for _, num := range row {
        for _, rule := range filteredRules {
            if rule.After == num {
                if _, ok := passed[rule.Before]; !ok {
                    return false
                }
            }
        }
        passed[num] = true
    }
    return true
}

func order(rules []Rule, row []int) []int {
    // Filter rules
    filteredRules := make([]Rule, 0)
    for _, rule := range rules {
        if slices.Contains(row, rule.After) && slices.Contains(row, rule.Before) {
            filteredRules = append(filteredRules, rule)
        }
    }
    // Order row
    for i := 0; i < len(row); i++ {
        for j := 0; j < len(row); j++ {
            if i == j {
                continue
            }
            for _, rule := range filteredRules {
                if row[i] == rule.After && row[j] == rule.Before {
                    row[i], row[j] = row[j], row[i]
                }
            }
        }
    }
    return row
}

func solve(rules []Rule, data [][]int) (int, int) {
    ret := 0
    ret2 := 0
    for _, row := range data {
        if isOk(rules, row) {
            // Get the middle number from row
            ret += row[(len(row) / 2)]
        } else {
            row = order(rules, row)
            ret2 += row[(len(row) / 2)]
        }
    }
    return ret, ret2
}

func solve2(rules []Rule, data [][]int) int {
    return 0
}

func main() {
    rules, data := loadData("test.txt")
    fmt.Println(solve(rules, data))
}
