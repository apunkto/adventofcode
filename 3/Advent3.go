package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	a := readFromFile()
	fmt.Println(Trees(a, 1, 1) * Trees(a, 3, 1) * Trees(a, 5, 1) * Trees(a, 7, 1) * Trees(a, 1, 2))

}

func Trees(a []string, right int, down int) int {
	trees := 0
	for i, l := range a {
		if i%down == 0 && l[i/down*right%len(l)] == '#' {
			trees++
		}
	}
	return trees
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
