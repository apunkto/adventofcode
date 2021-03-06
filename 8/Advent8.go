package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var re = regexp.MustCompile("([a-z]{3})\\s\\+?(-?[0-9]*$)")

func main() {
	start := time.Now()

	a := readFromFile()

	for i, l := range a {
		ll := l
		if strings.Contains(l, "nop") {
			l = strings.Replace(l, "nop", "jmp", 1)
		} else if strings.Contains(l, "jmp") {
			l = strings.Replace(l, "jmp", "nop", 1)
		} else {
			continue
		}

		a[i] = l
		m := make(map[int]bool)
		c := run(a, 0, 0, m)
		if m[len(a)-1] {
			fmt.Printf("Part 2 asnwer: %d\n", c)
			break
		}
		a[i] = ll
	}
	elapsed := time.Since(start)
	log.Printf("Part 2 took %s", elapsed)

}

func run(a []string, c int, i int, m map[int]bool) int {
	if i == len(a) || m[i] {
		return c
	}

	f := re.FindStringSubmatch(a[i])
	m[i] = true
	switch f[1] {
	case "nop":
		return run(a, c, i+1, m)
	case "acc":
		num, _ := strconv.Atoi(f[2])
		return run(a, c+num, i+1, m)
	case "jmp":
		num, _ := strconv.Atoi(f[2])
		return run(a, c, i+num, m)
	}

	panic("Dont come here!")
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
