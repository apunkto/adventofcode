package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	strings "strings"
)

func main() {
	a := readFromFile()
	i := 0
	for _, value := range a {
		split := strings.Split(value, " ")
		bounds := strings.Split(split[0], "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])
		letter := strings.ReplaceAll(split[1], ":", "")
		password := split[2]
		count := strings.Count(password, letter)
		if count >= lower && count <= upper {
			i++
		}
	}
	fmt.Println(i)

	j := 0
	for _, value := range a {
		split := strings.Split(value, " ")
		bounds := strings.Split(split[0], "-")
		pos1, _ := strconv.Atoi(bounds[0])
		pos2, _ := strconv.Atoi(bounds[1])
		letter := strings.ReplaceAll(split[1], ":", "")
		password := []rune(split[2])
		fmt.Println(value + " " + string(password[pos1-1]) + " " + string(password[pos2-1]))
		pos1Letter := string(password[pos1-1])
		pos2Letter := string(password[pos2-1])

		if (pos1Letter == letter || pos2Letter == letter) && (pos1Letter != pos2Letter) {
			j++
		}
	}
	fmt.Println(j)

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
