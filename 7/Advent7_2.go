package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	a := readFromFile()
	c := countContains(a, "shiny gold")
	fmt.Printf("MaAnswerx: %d\n", c)
}

func countContains(a []string, s string) int {
	re := regexp.MustCompile("(?P<Count>[0-9]) (?P<Bag>[a-z]* [a-z]*) (bag)")
	c := 0

	for _, l := range a {
		container := strings.Split(l, "contain")
		if strings.Contains(container[0], s) {
			f := re.FindAllStringSubmatch(container[1], -1)
			for _, ff := range f {
				ffs, _ := strconv.Atoi(ff[1])
				c += ffs + ffs*countContains(a, ff[2])
			}
			break
		}
	}
	return c
}
