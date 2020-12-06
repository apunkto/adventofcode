package main

import (
	"fmt"
)

func main() {
	a := readFromFile()

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if contains(a, 2020-a[i]-a[j]) {
				fmt.Println(a[i])
				fmt.Println(a[j])
				fmt.Println(2020 - a[i] - a[j])

				fmt.Println((2020 - a[i] - a[j]) * a[i] * a[j])
				break
			}
		}
	}
}
