package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var a = readFromFile()
var re = regexp.MustCompile("#")

func main() {
	mp := make(map[int][]string)
	mp[0] = a
	//mp[1] = a
	//mp[-1] = a

	for r := 1; r <= 6; r++ {
		mp2 := make(map[int][]string)
		mp = incDim(mp, r)

		for zz := -r; zz <= r; zz++ {

			for i, v := range mp[zz] {
				nc := ""
				for j, c := range v {
					cnt := count(i, j, zz, mp, "#")
					if c == '#' {
						if cnt >= 2 && cnt <= 3 {
							nc = nc + "#"
						} else {
							nc = nc + "."
						}
					} else if c == '.' {
						if cnt == 3 {
							nc = nc + "#"
						} else {
							nc = nc + "."
						}
					}
				}
				mp2[zz] = append(mp2[zz], nc)
			}
			/*if reflect.DeepEqual(a, b) {
				break
			}
			*/
		}
		// Copy from the original map to the target map
		for key, value := range mp2 {
			mp[key] = value
		}

	}
	cnt := 0
	for _, v := range mp {
		for _, vv := range v {
			for _, vvv := range vv {
				if string(vvv) == "#" {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)

}

func incDim(mp map[int][]string, dim int) map[int][]string {

	dim2 := len(mp[0])
	mp2 := make(map[int][]string)

	for zz := -dim; zz <= dim; zz++ {
		if mp[zz] == nil {
			mp2[zz] = emptyLayer(dim2 + 2)
		} else {
			mp2[zz] = append(mp2[zz], emptyLine(dim2+2))
			for _, v := range mp[zz] {
				nc := "." + v + "."
				mp2[zz] = append(mp2[zz], nc)
			}
			mp2[zz] = append(mp2[zz], emptyLine(dim2+2))
		}
	}
	return mp2
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

func count(i int, j int, z int, mp map[int][]string, c string) int {
	cnt := 0
	for zz := z - 1; zz < z+2; zz++ {
		aa := mp[zz]
		if aa != nil {
			for ii := i - 1; ii < i+2; ii++ {
				for jj := j - 1; jj < j+2; jj++ {
					if ii > -1 && ii < len(aa) && jj > -1 && jj < len(aa[0]) && !(i == ii && j == jj && z == zz) {
						if string(aa[ii][jj]) == c {
							cnt++
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
