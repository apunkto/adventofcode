package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var a = readFromFile()

func main() {
	dim := len(a)

	var mp = [][]string{}
	for i, s := range a {
		mp = append(mp, []string{})
		for _, ss := range s {
			mp[i] = append(mp[i], string(ss))
		}
	}
	monsterInc(dim, mp)
	mp = rotate(mp)
	monsterInc(dim, mp)
	mp = rotate(mp)
	monsterInc(dim, mp)
	mp = rotate(mp)
	monsterInc(dim, mp)
	mp = flipH(mp)
	monsterInc(dim, mp)
	mp = rotate(mp)
	monsterInc(dim, mp)
	mp = rotate(mp)
	monsterInc(dim, mp)
	mp = rotate(mp)
	monsterInc(dim, mp)

	c := 0
	for _, m := range mp {
		for _, mm := range m {
			if mm == "#" {
				c++
			}
		}
	}
	fmt.Println(c)
}

func monsterInc(dim int, mp [][]string) {
	for i := 0; i < dim-3; i++ {
		for j := 1; j < dim-20; j++ {
			if mp[i][j] == "#" && isTail(mp, i, j) {
				identifyMonster(mp, i, j)
			}
		}
	}
}

func identifyMonster(a [][]string, i int, j int) {
	a[i][j] = "O"
	a[i][j+5] = "O"
	a[i][j+6] = "O"
	a[i][j+11] = "O"
	a[i][j+12] = "O"
	a[i][j+17] = "O"
	a[i][j+18] = "O"
	a[i][j+19] = "O"
	a[i+1][j+1] = "O"
	a[i+1][j+4] = "O"
	a[i+1][j+7] = "O"
	a[i+1][j+10] = "O"
	a[i+1][j+13] = "O"
	a[i+1][j+16] = "O"
	a[i-1][j+18] = "O"
}

func isTail(a [][]string, i int, j int) bool {
	return a[i][j] == "#" && a[i][j+5] == "#" && a[i][j+6] == "#" &&
		a[i][j+11] == "#" && a[i][j+12] == "#" &&
		a[i][j+17] == "#" && a[i][j+18] == "#" &&
		a[i][j+19] == "#" && a[i+1][j+1] == "#" && a[i+1][j+4] == "#" &&
		a[i+1][j+7] == "#" && a[i+1][j+10] == "#" && a[i+1][j+13] == "#" &&
		a[i+1][j+16] == "#" && a[i-1][j+18] == "#"
}

func flipH(m [][]string) [][]string {
	r := [][]string{}
	for i := 0; i < len(m); i++ {
		r = append(r, []string{})
		for j := 0; j < len(m); j++ {
			r[i] = append(r[i], "")
		}

		for j := 0; j < len(m)/2; j++ {
			r[i][j] = m[i][len(m)-1-j]
			r[i][len(m)-1-j] = m[i][j]
		}
	}
	return r
}

func rotate(m [][]string) [][]string {
	r := [][]string{}
	for i := 0; i < len(m); i++ {
		r = append(r, []string{})
		for j := 0; j < len(m); j++ {
			r[i] = append(r[i], m[len(m)-j-1][i])
		}
	}
	return r
}

func printMap(s [][]string) {
	for _, b := range s {
		fmt.Println(b)
	}
	fmt.Println()
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
