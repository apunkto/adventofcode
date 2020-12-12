package main

import (
	"log"
	"math"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	a := readFromFile()
	x := 0
	y := 0
	wx := 10
	wy := 1
	for _, l := range a {
		f := re.FindStringSubmatch(l)
		c, _ := strconv.Atoi(f[2])
		switch f[1] {
		case "F":
			x = x + wx*c
			y = y + wy*c
		case "N":
			wx, wy = move(wx, wy, 3, c)
		case "S":
			wx, wy = move(wx, wy, 1, c)
		case "E":
			wx, wy = move(wx, wy, 0, c)
		case "W":
			wx, wy = move(wx, wy, 2, c)
		case "R":
			cc := c / 90
			for i := 0; i < cc; i++ {
				wx, wy = wy, -wx
			}
		case "L":
			cc := c / 90
			for i := 0; i < cc; i++ {
				wx, wy = -wy, wx
			}
		}
	}

	log.Println(math.Abs(float64((y))) + math.Abs(float64(x)))

	log.Printf("Part 2 took %s", time.Since(start))

}
