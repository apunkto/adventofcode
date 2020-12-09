package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	run(readFromFile(), 41682220)
	log.Printf("Part 2 took %s", time.Since(start))
}

func run(a []int, invalid int) {
	for i, v := range a {
		r := []int{v}
		sum := v
		if v >= invalid {
			continue
		}

		for _, v2 := range a[i+1:] {
			if v2 >= invalid {
				continue
			}
			sum += v2
			if sum > invalid {
				break
			}
			r = append(r, v2)
			if sum == invalid {
				fmt.Println(r[0] + r[len(r)-1])
				return
			}
		}
	}
	return
}
