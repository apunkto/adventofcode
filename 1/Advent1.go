package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	a := readFromFile()

	for _, value := range a {
		if contains(a, 2020-value) {
			fmt.Println((2020 - value) * value)
			break
		}
	}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if contains(a, 2020-a[i]-a[j]) {
				fmt.Println((2020 - a[i] - a[j]) * a[i] * a[j])
				return
			}
		}
	}
}

func readFromFile() []int {
	var a []int
	f, err := os.Open("1/input.txt")

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
