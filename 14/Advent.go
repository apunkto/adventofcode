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

var re = regexp.MustCompile("([0-9]+).* = ([0-9]+)")
var me = regexp.MustCompile("mask = (.*)")

func main() {

	a := readFromFile()
	mask := ""
	mp := make(map[int]int64)

	for _, l := range a {
		if strings.Contains(l, "mask") {
			mask = me.FindStringSubmatch(l)[1]
			continue
		}
		f := re.FindStringSubmatch(l)
		i, _ := strconv.Atoi(f[1])
		n, _ := strconv.Atoi(f[2])

		d := ToBits(int64(n))
		dd := Apply(mask, d)
		ddd, _ := strconv.ParseInt(dd, 2, 64)
		mp[i] = ddd
	}

	var sum int64 = 0
	for _, x := range mp {
		sum += x
	}
	fmt.Println(sum)

}

func Apply(mask string, d string) string {
	for i, c := range mask {
		if string(c) != "X" {
			d = d[:i] + string(c) + d[i+1:]
		}
	}
	return d
}

func ToBits(l int64) string {
	return fmt.Sprintf("%036s", strconv.FormatInt(l, 2))
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
