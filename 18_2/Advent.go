package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var a = readFromFile()
var re = regexp.MustCompile("\\(([^\\(\\)]+)\\)")
var re2 = regexp.MustCompile("([0-9]+)(\\+)([0-9]+)")
var re3 = regexp.MustCompile("([0-9]+)(\\*)([0-9]+)")

func main() {
	sum := 0
	for _, v := range a {
		vv := strings.ReplaceAll(v, " ", "")
		sum += doMagic(vv)
	}
	fmt.Println(sum)
}

func doMagic(vv string) int {
	ss := re.FindStringSubmatch(vv)
	if ss != nil {
		sss := strconv.Itoa(doMagic(ss[1]))
		return doMagic(strings.Replace(vv, ss[0], sss, 1))
	}

	if !strings.Contains(vv, "*") && !strings.Contains(vv, "+") {
		x, _ := strconv.Atoi(vv)
		return x
	}

	ssPlus := re2.FindStringSubmatch(vv)
	if ssPlus != nil {
		return doMagic(strings.Replace(vv, ssPlus[0], calc(ssPlus), 1))
	}

	ssM := re3.FindStringSubmatch(vv)
	if ssM != nil {
		return doMagic(strings.Replace(vv, ssM[0], calc(ssM), 1))
	}
	panic("dont come here!")
}

func calc(ssPlus []string) string {
	x, _ := strconv.Atoi(ssPlus[1])
	y, _ := strconv.Atoi(ssPlus[3])
	op := ssPlus[2]
	return strconv.Itoa(doMath(x, y, op))
}

func doMath(x int, y int, s string) int {
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
