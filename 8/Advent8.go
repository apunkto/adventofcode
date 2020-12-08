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

func main() {

	a := readFromFile()

	for i, l := range a {
		if strings.Contains(l, "nop") {
			l = strings.Replace(l, "nop", "jmp", 1)
		} else if strings.Contains(l, "jmp") {
			l = strings.Replace(l, "jmp", "nop", 1)
		} else {
			continue
		}

		a2 := make([]string, len(a))
		copy(a2, a)
		a2[i] = l
		m := make(map[int]bool)
		c := countContains(a2, 0, 0, m)
		if m[len(a)-1] == true {
			fmt.Printf("MaAnswerx: %d\n", c)
		}
	}

}

func countContains(a []string, c int, i int, m map[int]bool) int {
	if i == len(a) {
		return c
	}

	re := regexp.MustCompile("([a-z]{3})\\s\\+?(-?[0-9]*$)")
	f := re.FindStringSubmatch(a[i])
	if m[i] == true {
		return c
	}
	m[i] = true
	switch f[1] {
	case "nop":
		return countContains(a, c, i+1, m)
	case "acc":
		num, _ := strconv.Atoi(f[2])
		return countContains(a, c+num, i+1, m)
	case "jmp":
		num, _ := strconv.Atoi(f[2])
		return countContains(a, c, i+num, m)
	}

	return c
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
