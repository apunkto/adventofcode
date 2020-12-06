package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var rules = []*regexp.Regexp{
	regexp.MustCompile("(\\s|^)byr:(19[2-9][0-9]|200[0-2])(\\s|$)"),
	regexp.MustCompile("(\\s|^)iyr:(201[0-9]|2020)(\\s|$)"),
	regexp.MustCompile("(\\s|^)eyr:(202[0-9]|2030)(\\s|$)"),
	regexp.MustCompile("(\\s|^)hgt:((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)(\\s|$)"),
	regexp.MustCompile("(\\s|^)hcl:(#[A-Fa-f0-9]{6})(\\s|$)"),
	regexp.MustCompile("(\\s|^)ecl:(amb|blu|brn|gry|grn|hzl|oth)(\\s|$)"),
	regexp.MustCompile("(\\s|^)pid:([0-9]{9})(\\s|$)"),
}

func main() {
	a := readFromFile()
	r := 0
	passport := ""
	for _, l := range a {
		if l == "" {
			if validatePassport(passport) {
				r++
			}
			passport = ""
		} else {
			passport = passport + " " + l
		}
	}
	if validatePassport(passport) {
		r++
	}
	fmt.Println(r)
}

func validatePassport(s string) bool {
	for _, rule := range rules {
		if !rule.MatchString(s) {
			return false
		}
	}
	return true
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
