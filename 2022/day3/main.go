package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	input := strings.Split(string(data), "\r\n")

	for _, container := range input {
		middle := len(container) / 2
		first := container[:middle]
		second := container[middle:]

		for _, char := range first {
			if strings.Contains(second, string(char)) {
				fmt.Print(string(char))
			}
		}
	}
}
