package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var a = readFromFile()

func main() {
	//90135
	sum := 0
	for _, v := range a {
		vv := strings.ReplaceAll(v, " ", "")
		sum += doMagic(vv)
	}
	log.Print(sum)

}

func doMagic(vv string) int {
	num := ""
	op := "+"

	p := 0
	x := 0
	y := 0
	pp := ""
	for i, s := range vv {
		ss := string(s)

		if ss == ")" {
			p--
			if p == 0 {
				x = doMath(x, doMagic(pp[1:]), op)
				pp = ""
				continue
			}
		}

		if ss == "(" {
			p++
		}

		if p > 0 {
			pp = pp + ss
		} else {

			if unicode.IsDigit(s) {
				num = num + ss
				if i == len(vv)-1 {
					y, _ = strconv.Atoi(num)
					x = doMath(x, y, op)
					continue
				}
			} else {
				if num != "" {
					y, _ = strconv.Atoi(num)
					x = doMath(x, y, op)
				}
				op = ss
				num = ""
			}
		}
	}

	return x
}

func doMath(x int, y int, s string) int {
	if s == "" {
		return y
	}
	if s == "+" {
		return x + y
	}
	return x * y
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
