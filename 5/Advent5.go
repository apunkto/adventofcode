package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	a := readFromFile()

	var ids []int
	for _, l := range a {
		ids = append(ids, (decode(l[:7])*8 + decode(l[7:])))
	}
	sort.Ints(ids)
	fmt.Printf("Max: %d\n", ids[len(ids)-1])

	for i, id := range ids {
		if ids[i+1] != id+1 {
			fmt.Printf("Free seat: %d\n", id+1)
			return
		}
	}
}

func decode(l string) int {
	mid := (math.Exp2(float64(len(l))) - 1) / 2

	for i := 0; i < len(l); i++ {
		m := math.Pow(2, float64(len(l)-2-i))
		switch l[i] {
		case 'F', 'L':
			mid = mid - m
		case 'B', 'R':
			mid = mid + m
		}
	}
	return int(mid)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return y
	}
	return x
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
