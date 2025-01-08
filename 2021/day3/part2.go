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
	data, err := os.ReadFile("day3.txt")
	check(err)

	//split input into 2d slice
	var in [][]string
	vals := strings.Split(string(data), "\n")
	for _, val := range vals {
		in = append(in, strings.Split(val, ""))
	}

	//transpose
	transposed := make([][]string, len(in[0])-1)
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0])-1; j++ {
			transposed[j] = append(transposed[j], in[i][j])
		}
	}

	out := in
	rule := findRule(in, transposed)
	for i, val := range rule {
		for _, d := range out {
			if d[i] != val {
				out = append(out[i:], out[i+1:]...)
			}
		}
	}
	fmt.Printf("answer: %q", rule)
}

func findRule(diagnostics [][]string, transposed [][]string) []string {
	var rule []string
	for _, val := range transposed {
		numZeros := strings.Count(strings.Join(val, ""), "0")
		numOnes := strings.Count(strings.Join(val, ""), "1")

		if numZeros > numOnes {
			// fmt.Printf("%d: %d ? %d = %d\n", i, numZeros, numOnes, 0)
			rule = append(rule, "0")
		} else {
			// fmt.Printf("%d: %d ? %d = %d\n", i, numZeros, numOnes, 1)
			rule = append(rule, "1")
		}
	}
	return rule
}
