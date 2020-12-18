package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var a = readFromFile()

func main() {
	mp := make(map[int]map[int][]string)
	mp[0] = make(map[int][]string)
	mp[0][0] = a

	for r := 1; r <= 6; r++ {
		mp2 := make(map[int]map[int][]string)
		mp = incDim(mp, r)
		for zzz := -r; zzz <= r; zzz++ {
			mp2[zzz] = make(map[int][]string)

			for zz, vv := range mp[zzz] {
				mp2[zzz][zz] = []string{}
				for i, v := range vv {
					nc := ""
					for j, c := range v {
						cnt := count(i, j, zz, zzz, mp, "#")
						if c == '#' {
							if cnt >= 2 && cnt <= 3 {
								nc = nc + "#"
							} else {
								nc = nc + "."
							}
						} else if c == '.' {
							if cnt == 3 {
								cnt = count(i, j, zz, zzz, mp, "#")
								nc = nc + "#"
							} else {
								nc = nc + "."
							}
						}
					}
					mp2[zzz][zz] = append(mp2[zzz][zz], nc)
				}
			}

		}
		for key, value := range mp2 {
			for k, v := range value {
				mp[key][k] = v
			}

		}
	}
	cnt := 0
	for _, v := range mp {
		for _, vv := range v {
			for _, vvv := range vv {
				for _, vvvv := range vvv {
					if string(vvvv) == "#" {
						cnt++
					}
				}
			}
		}
	}
	fmt.Println(cnt)

}

func incDim(mp map[int]map[int][]string, dim int) map[int]map[int][]string {

	dim2 := len(mp[0][0])
	mp2 := make(map[int]map[int][]string)
	for zzz := -dim; zzz <= dim; zzz++ {
		if mp[zzz] == nil {
			mp2[zzz] = empty4Layer(dim, dim2+2)
		} else {
			for zz := -dim; zz <= dim; zz++ {
				if mp2[zzz] == nil {
					mp2[zzz] = make(map[int][]string)
				}
				if mp[zzz][zz] == nil {
					mp2[zzz][zz] = emptyLayer(dim2 + 2)
				} else {
					mp2[zzz][zz] = append(mp2[zzz][zz], emptyLine(dim2+2))
					for _, v := range mp[zzz][zz] {
						nc := "." + v + "."
						mp2[zzz][zz] = append(mp2[zzz][zz], nc)
					}
					mp2[zzz][zz] = append(mp2[zzz][zz], emptyLine(dim2+2))
				}
			}
		}
	}
	return mp2
}

func empty4Layer(cc, c int) map[int][]string {
	var ans = make(map[int][]string)
	for i := -cc; i <= cc; i++ {
		ans[i] = emptyLayer(c)
	}
	return ans
}

func emptyLine(c int) string {
	nc := ""
	for j := 0; j < c; j++ {
		nc = nc + "."
	}
	return nc
}

func emptyLayer(c int) []string {
	var ans []string
	for i := 0; i < c; i++ {
		ans = append(ans, emptyLine(c))
	}
	return ans
}

func count(i int, j int, z int, y int, mp map[int]map[int][]string, c string) int {
	cnt := 0
	for yy := y - 1; yy < y+2; yy++ {
		bb := mp[yy]
		if bb != nil {
			for zz := z - 1; zz < z+2; zz++ {
				aa := bb[zz]
				if aa != nil {
					for ii := i - 1; ii < i+2; ii++ {
						for jj := j - 1; jj < j+2; jj++ {
							if ii > -1 && ii < len(aa) && jj > -1 && jj < len(aa[0]) && !(i == ii && j == jj && z == zz && y == yy) {
								if string(aa[ii][jj]) == c {
									cnt++
								}
							}
						}
					}
				}
			}
		}
	}

	return cnt
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
