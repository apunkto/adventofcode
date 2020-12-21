package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Tile struct {
	id int
	mp [10][10]string
}

var dim = 12
var a = readFromFile()
var re = regexp.MustCompile("^Tile ([0-9]+):$")

func main() {

	mp := readData()
	fp := [][]Tile{}
	ffp, _, _ := tryTiles(mp, fp, 0, -1, dim-1)
	fmt.Println(ffp[0][0].id * ffp[0][11].id * ffp[11][0].id * ffp[11][11].id)
	/*
		for _, f := range ffp {
			for j := 1; j < len(f[0].mp)-1; j++ {
				for _, ff := range f {
					for w := 1; w < len(ff.mp)-1; w++ {
						fmt.Print(ff.mp[j][w])
					}
				}
				fmt.Println()
			}

		}
	*/
}

func tryTiles(mp []Tile, fp [][]Tile, pos int, xx int, yy int) ([][]Tile, []Tile, bool) {

	if len(mp) == 0 {
		return fp, mp, true
	}
	if pos == (len(mp)) {
		return fp, mp, false
	}

	tp := copyFp(fp)
	x, y := xx, yy
	if yy == dim-1 {
		x = xx + 1
		tp = append(tp, []Tile{})
		y = 0
	} else {
		y = yy + 1
	}

	i, tiles, b := check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}
	mp[pos].mp = rotate(mp[pos].mp)
	i, tiles, b = check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}
	mp[pos].mp = rotate(mp[pos].mp)
	i, tiles, b = check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}
	mp[pos].mp = rotate(mp[pos].mp)
	i, tiles, b = check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}
	mp[pos].mp = flipH(mp[pos].mp)
	i, tiles, b = check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}
	mp[pos].mp = rotate(mp[pos].mp)
	i, tiles, b = check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}
	mp[pos].mp = rotate(mp[pos].mp)
	i, tiles, b = check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}
	mp[pos].mp = rotate(mp[pos].mp)
	i, tiles, b = check2(mp, fp, pos, x, y, tp)
	if b {
		return i, tiles, true
	}

	return tryTiles(mp, fp, pos+1, xx, yy)

}

func check2(mp []Tile, fp [][]Tile, pos int, x int, y int, tp [][]Tile) ([][]Tile, []Tile, bool) {
	if check(fp, mp[pos].mp, x, y) {
		xx, yy, b := tryTiles(minusTile(mp, mp[pos].id), appendTile(tp, x, mp[pos]), 0, x, y)
		if b {
			return xx, yy, true
		}
	}
	return nil, nil, false
}

func copyFp(fp [][]Tile) [][]Tile {
	dp := make([][]Tile, len(fp))
	for i := range fp {
		dp[i] = make([]Tile, len(fp[i]))
		copy(dp[i], fp[i])
	}
	return dp
}

func appendTile(tp [][]Tile, x int, tile Tile) [][]Tile {
	rp := copyFp(tp)
	rp[x] = append(rp[x], tile)
	return rp
}

func minusTile(mp []Tile, id int) []Tile {
	resp := []Tile{}
	for _, v := range mp {
		if v.id != id {
			resp = append(resp, v)
		}
	}
	return resp
}

func check(fp [][]Tile, mp [10][10]string, x int, y int) bool {

	if y > 0 { //check left
		if !lrMatch(fp[x][y-1].mp, mp, 9) {
			return false
		}
	}
	if x > 0 { //check top
		if !tbMatch(fp[x-1][y].mp, mp, 9) {
			return false
		}
	}
	return true
}

func lrMatch(mp [10][10]string, mp2 [10][10]string, c int) bool {
	for i := 0; i < 10; i++ {
		if mp[i][c] != mp2[i][9-c] {
			return false
		}
	}
	return true
}

func tbMatch(mp [10][10]string, mp2 [10][10]string, c int) bool {
	for i := 0; i < 10; i++ {
		if mp[c][i] != mp2[9-c][i] {
			return false
		}
	}
	return true
}

func readData() []Tile {
	mp := []Tile{}
	tileNr := 0
	i := 0
	var tile = [10][10]string{}
	for _, l := range a {
		if l == "" {
			continue
		}
		ff := re.FindStringSubmatch(l)
		if ff != nil {
			tileNr, _ = strconv.Atoi(ff[1])
			tile = [10][10]string{}
			i = 0
		} else {
			for j, s := range l {
				tile[i][j] = string(s)
			}
			i++
			if i == 10 {
				mp = append(mp, Tile{tileNr, tile})
			}
		}
	}
	return mp
}

func flipH(m [10][10]string) [10][10]string {
	r := [10][10]string{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			r[i][j] = m[i][10-1-j]
			r[i][10-1-j] = m[i][j]
		}
	}
	return r
}

func rotate(m [10][10]string) [10][10]string {
	r := [10][10]string{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			r[i][j] = m[10-j-1][i]
		}
	}
	return r
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
