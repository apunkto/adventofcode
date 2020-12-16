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

var re = regexp.MustCompile("(.*): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
var a = readFromFile()

func main() {
	m := make(map[string][4]int)

	row := 0
	for _, l := range a {
		row++
		if l == "" {
			break
		}
		evaluateRuleMap(l, m)
	}

	ticket := strings.Split(a[row+1], ",")

	var validtickets [][]int
	for _, l := range a[row+4:] {
		valid := true
		s := strings.Split(l, ",")
		var tl []int
		for _, ss := range s {
			x, _ := strconv.Atoi(ss)
			if !isValidByRules(x, m) {
				valid = false
				break
			}
			tl = append(tl, x)
		}
		if valid {
			validtickets = append(validtickets, tl)
		}
	}

	c := len(validtickets[0])
	m2 := make(map[string][]int)
	m3 := make(map[string]int)

	check(m, c, validtickets, m2)
	for len(m2) != len(m3) {
		for k, v := range m2 {
			if len(v) == 1 {
				m3[k] = v[0]
				removeDuplicates(k, v[0], m2)
			}
		}
	}
	ans := 1
	for k, v := range m3 {
		if strings.Contains(k, "departure") {
			tv, _ := strconv.Atoi(ticket[v])
			ans = ans * tv
		}
	}
	fmt.Println(ans)
}

func evaluateRuleMap(l string, m map[string][4]int) {
	f := re.FindStringSubmatch(l)
	x1, _ := strconv.Atoi(f[2])
	x2, _ := strconv.Atoi(f[3])
	x3, _ := strconv.Atoi(f[4])
	x4, _ := strconv.Atoi(f[5])
	m[f[1]] = [4]int{x1, x2, x3, x4}
}

func removeDuplicates(kk string, i int, m2 map[string][]int) {
	for k, v := range m2 {
		if k != kk {
			for j, r := range v {
				if r == i {
					m2[k] = remove(v, j)
					return
				}
			}
		}
	}
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func check(m map[string][4]int, c int, validtickets [][]int, m2 map[string][]int) {
	for k, mm := range m {
		for i := 0; i < c; i++ {
			allValid := true
			for _, l := range validtickets {
				x := l[i]
				if !isValid(x, mm) {
					allValid = false
					break
				}
			}
			if allValid {
				m2[k] = append(m2[k], i)
			}
		}
	}
}

func isValidByRules(x int, m map[string][4]int) bool {
	for _, l := range m {
		if isValid(x, l) {
			return true
		}
	}
	return false
}

func isValid(x int, l [4]int) bool {
	return (x >= l[0] && x <= l[1]) || (x >= l[2] && x <= l[3])
}

func readFromFile() []string {
	var a []string
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	return a
}
