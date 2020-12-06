package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Task1() {
	a := readFromFile()
	group := ""
	sum := 0
	for _, l := range a {
		if l == "" {
			sum += count(group)
			group = ""
		} else {
			group += l
		}
	}
	sum += count(group)

	fmt.Println(sum)
}

func count(s string) int {
	m := make(map[string]bool)
	for _, g := range s {
		m[string(g)] = true
	}

	return len(m)
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
