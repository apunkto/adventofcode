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
var maskRe = regexp.MustCompile("mask = (.*)")

func main() {

	a := readFromFile()
	mask := ""
	mp := make(map[int64]int64)

	for _, l := range a {
		if strings.Contains(l, "mask") {
			mask = maskRe.FindStringSubmatch(l)[1]
			continue
		}

		f := re.FindStringSubmatch(l)
		i, _ := strconv.Atoi(f[1])
		n, _ := strconv.Atoi(f[2])

		d := ToBits(int64(i))
		dd := Apply(mask, d)

		ch := make(map[int64]bool)
		FloatBits(string(dd), ch)

		for chh, _ := range ch {
			mp[chh] = int64(n)
		}

	}

	var sum int64 = 0
	for _, x := range mp {
		sum += x
	}
	fmt.Println(sum)
}

func FloatBits(s string, ch map[int64]bool) {
	if !strings.Contains(s, "X") {
		ddd, _ := strconv.ParseInt(s, 2, 64)
		ch[ddd] = true
		return
	}
	FloatBits(strings.Replace(s, "X", "0", 1), ch)
	FloatBits(strings.Replace(s, "X", "1", 1), ch)
}

func Apply(mask string, d string) string {

	for i, c := range mask {
		if c != '0' {
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
