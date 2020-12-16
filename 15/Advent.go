package main

import (
	"fmt"
	"log"
	"time"
)

//19,0,5,1,10,13
var puzzle = []int{19, 0, 5, 1, 10, 13}

func main() {
	start := time.Now()

	mp := CreateMap()

	x := 0
	for i := len(mp) + 1; i <= 30000000; i++ {
		y := mp[x][1]
		if mp[x][0] == 0 {
			x = 0
		} else {
			x = y - mp[x][0]
		}
		mp[x] = [2]int{mp[x][1], i}
	}

	fmt.Println(x)
	log.Printf("Puzzle took %s", time.Since(start))

}

func CreateMap() map[int][2]int {
	mp := make(map[int][2]int)
	for i, p := range puzzle {
		mp[p] = [2]int{0, i + 1}
	}
	return mp
}
