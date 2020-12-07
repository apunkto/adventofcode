package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main2() {

	a := readFromFile()
	c := canContain(a, make(map[string]bool), "shiny gold")
	fmt.Printf("MaAnswerx: %d\n", c)

}

func canContain(a []string, m map[string]bool, s string) int {
	re := regexp.MustCompile("^(?P<Bag>[a-z]* [a-z]*) (bags)")
	c := 0
	for _, l := range a {
		container := strings.Split(l, "contain")
		if strings.Contains(container[1], s) {
			f := re.FindStringSubmatch(container[0])
			if m[f[1]] != true {
				m[f[1]] = true
				c += 1 + canContain(a, m, f[1])
			}
		}
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
