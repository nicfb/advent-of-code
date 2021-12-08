package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("day1.txt")
	check(err)

	vals := strings.Split(string(data), "\n")

	prev := make([]string, 3)
	prevTotal := 0
	totalIncreased := 0
	for i, current := range vals {
		//append to list
		prev = append(prev, current)
		//remove first element
		prev = prev[1:]

		//wait until we have the first 3 values
		//before comparing
		if i <= 2 {
			continue
		}

		//add em all up
		currTotal := 0
		for _, str := range prev {
			if str == "" {
				continue
			}
			val, err := strconv.Atoi(str)
			check(err)

			currTotal += val
		}

		if currTotal > prevTotal {
			totalIncreased++
		}

		prevTotal = currTotal
	}
	fmt.Printf("increased %d times\n", totalIncreased)
}
