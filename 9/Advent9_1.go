package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main2() {
	a := readFromFile()
	pre := 25

	for i, value := range a {
		if i >= pre {
			if noSum(value, a[i-pre:i]) {
				fmt.Println(value)
				break
			}
		}
	}
}

func noSum(v int, b []int) bool {
	for _, x := range b {
		if 2*x != v && contains(b, v-x) {
			return false
		}
	}
	return true

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
