package main

import (
	"fmt"
)

func Task2() {
	a := readFromFile()
	group := ""
	c := 0
	sum := 0
	for _, l := range a {
		if l == "" {
			sum += count2(c, group)
			group = ""
			c = 0
		} else {
			c++
			group += l
		}
	}
	sum += count2(c, group)

	fmt.Println(sum)
}

func count2(total int, s string) int {
	m := make(map[string]int)
	c := 0
	for _, g := range s {
		m[string(g)] += 1
	}
	for _, v := range m {
		if v == total {
			c++
		}
	}
	return c
}
