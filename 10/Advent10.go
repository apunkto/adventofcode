package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	a := readFromFile()
	a = append(a, 0)
	sort.Ints(a)
	max := a[len(a)-1]
	a = append(a, max+3)

	m := make(map[int]int)
	m[0] = 1
	for i, v := range a {
		for j := 3; j > 0; j-- {
			if i-j >= 0 && canComeFrom(v, a[i-j]) {
				m[i] += m[i-j]
			}
		}
	}
	fmt.Print(m[len(a)-1])
}

func canComeFrom(a int, b int) bool {
	return a-b < 4
}

func readFromFile() []int {
	var a []int
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var intA, _ = strconv.Atoi(scanner.Text())
		a = append(a, intA)
	}
	return a
}

func contains(arr []int, search int) bool {
	for _, a := range arr {
		if a == search {
			return true
		}
	}
	return false
}
