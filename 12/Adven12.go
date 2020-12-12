package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile("([A-Z]{1})([0-9]*$)")

var dirs = map[int][2]int{
	0: [2]int{1, 0},  // E
	1: [2]int{0, -1}, //S
	2: [2]int{-1, 0}, //W
	3: [2]int{0, 1},  //N
}

func main2() {
	a := readFromFile()
	x := 0
	y := 0
	dir := 0
	for _, l := range a {
		f := re.FindStringSubmatch(l)
		c, _ := strconv.Atoi(f[2])
		switch f[1] {
		case "F":
			x, y = move(x, y, dir, c)
		case "N":
			x, y = move(x, y, 3, c)
		case "S":
			x, y = move(x, y, 1, c)
		case "E":
			x, y = move(x, y, 0, c)
		case "W":
			x, y = move(x, y, 2, c)
		case "R":
			dir = (dir + (c / 90)) % 4
		case "L":
			dir = dir - (c / 90)
			if dir < 0 {
				dir = dir + 4
			}
		}
	}

	log.Println(math.Abs(float64((y))) + math.Abs(float64(x)))

}

func move(x int, y int, dir int, c int) (int, int) {
	x = x + dirs[dir][0]*c
	y = y + dirs[dir][1]*c
	return x, y
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
