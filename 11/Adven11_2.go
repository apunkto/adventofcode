package main

import (
	"fmt"
	"reflect"
)

var a = readFromFile()

func main() {

	b := make([]string, len(a))
	for {
		for i, v := range a {
			nc := ""

			for j, c := range v {
				if c == 'L' && seeCount(i, j) == 0 {
					nc = nc + "#"
				} else if c == '#' && seeCount(i, j) >= 5 {
					nc = nc + "L"
				} else {
					nc = nc + string(c)
				}
			}
			b[i] = nc
		}
		if reflect.DeepEqual(a, b) {
			break
		}
		copy(a, b)
	}
	c := 0
	for _, v := range a {
		matches := re.FindAllStringIndex(v, -1)
		c += len(matches)
	}

	fmt.Println(c)

}

func seeCount(i int, j int) int {
	c := 0
	for v := -1; v < 2; v++ {
		for h := -1; h < 2; h++ {
			if v == 0 && h == 0 {
				continue
			}
			if seesOccupied(i, j, h, v) {
				c++
			}
		}
	}
	return c
}

func seesOccupied(i int, j int, hor int, ver int) bool {
	y := j - hor
	x := i - ver

	for ; y >= 0 && y < len(a[0]) && x >= 0 && x < len(a); y, x = y-hor, x-ver {
		s := string(a[x][y])
		switch s {
		case "#":
			return true
		case "L":
			return false
		}
	}
	return false
}
