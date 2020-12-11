package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
)

var re = regexp.MustCompile("#")

func main2() {
	a := readFromFile()
	b := make([]string, len(a))
	for {
		for i, v := range a {
			nc := ""
			for j, c := range v {
				if c == 'L' && count(i, j, a) == 0 {
					nc = nc + "#"
				} else if c == '#' && count(i, j, a) >= 4 {
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
		fmt.Println(v)
	}

	fmt.Println(c)

}

func count(i int, j int, a []string) int {
	s := ""
	for ii := i - 1; ii < i+2; ii++ {
		for jj := j - 1; jj < j+2; jj++ {
			if ii > -1 && ii < len(a) && jj > -1 && jj < len(a[0]) && !(i == ii && j == jj) {
				s = s + string(a[ii][jj])
			}
		}
	}

	matches := re.FindAllStringIndex(s, -1)
	return len(matches)
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
