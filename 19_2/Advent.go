package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var a = readFromFile()
var nRe = regexp.MustCompile("^[0-9]+$")
var lRe = regexp.MustCompile("^\"([a-z]+)\"$")

func main() {

	rules := make(map[string]string)
	var m []string

	rules, m = readData(rules, m)
	rex := "^" + getRex("0", rules) + "$"
	var re = regexp.MustCompile(rex)

	cnt := 0
	for _, mm := range m {
		if re.MatchString(mm) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func getRex(s string, rules map[string]string) string {
	s = strings.TrimSpace(s)
	switch s {
	case "8":
		return "(" + getRex("42", rules) + "+)"
	case "11":
		rule42 := getRex("42", rules)
		rule31 := getRex("31", rules)
		ss := "("
		for i := 0; i < 4; i++ {
			ss = ss + "("
			for j := 0; j < i+1; j++ {
				ss = ss + rule42
			}
			for j := 0; j < i+1; j++ {
				ss = ss + rule31
			}
			ss = ss + ")|"
		}
		ss = ss[:len(ss)-1] + ")"
		return ss
	}

	if nRe.MatchString(s) {
		return getRex(rules[s], rules)
	}
	ss := lRe.FindStringSubmatch(s)
	if ss != nil {
		return ss[1]
	}
	if strings.Contains(s, "|") {
		rs := "("
		sss := strings.Split(s, "|")
		for _, ssss := range sss {
			rs = rs + getRex(ssss, rules) + "|"
		}
		rs = rs[:len(rs)-1] + ")"
		return rs
	}

	rs := ""
	sss := strings.Split(s, " ")
	for _, ssss := range sss {
		rs = rs + getRex(ssss, rules)
	}
	return rs
}

func readData(rules map[string]string, m []string) (map[string]string, []string) {
	row := 0
	for _, v := range a {

		if v == "" {
			break
		}
		row++
		rule := strings.Split(v, ":")

		rules[rule[0]] = rule[1]
	}
	for _, v := range a[row+1:] {
		m = append(m, v)
	}
	return rules, m
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
