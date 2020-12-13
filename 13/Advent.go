package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	a := readFromFile()
	b := strings.Split(a[1], ",")
	c := 0
	k := 0

	for i, l := range b {
		if i == len(b)-1 {
			break
		}
		x, _ := strconv.Atoi(l)
		if i == 0 {
			c = x
			k = x
		}
		if b[i+1] == "x" {
			continue
		}
		x2, _ := strconv.Atoi(b[i+1])
		for y := c; ; y = y + k {
			if (y+i+1)%x2 == 0 {
				k = k * x2
				c = y
				break
			}
		}
	}
	fmt.Println(c)

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
