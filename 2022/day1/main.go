package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

	var max []int
	var values []int
	for _, x := range input {
		if x == "" {
			total := 0
			for _, val := range values {
				total += val
			}

			if len(max) < 3 {
				max = append(max, total)
			} else {
				for j, m := range max {
					if total > m {
						max[j] = total
					}
					break
				}
			}

			sort.Ints(max)

			total = 0
			values = values[:0]
			fmt.Println(max)
			continue
		}

		val, err := strconv.Atoi(x)
		check(err)
		values = append(values, val)
	}

	fmt.Println(max)
	ans := 0
	for _, x := range max {
		ans += x
	}
	fmt.Println(ans)
}
