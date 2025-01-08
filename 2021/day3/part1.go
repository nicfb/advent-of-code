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
	data, err := os.ReadFile("day3.txt")
	check(err)

	//split input into 2d slice
	var in [][]string
	vals := strings.Split(string(data), "\n")
	for _, val := range vals {
		in = append(in, strings.Split(val, ""))
	}

	//transpose
	out := make([][]string, len(in[0])-1)
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0])-1; j++ {
			out[j] = append(out[j], in[i][j])
		}
	}

	//get most common value for each row
	gamma := ""
	for _, val := range out {
		numZeros := strings.Count(strings.Join(val, ""), "0")
		numOnes := strings.Count(strings.Join(val, ""), "1")
		if numZeros > numOnes {
			gamma += "0"
		} else {
			gamma += "1"
		}
	}

	gammaDecimal, err := strconv.ParseUint(gamma, 2, 64)
	check(err)

	//epsilon rate is inverse of gamma rate
	epsilon := 0xFFF ^ gammaDecimal

	power := gammaDecimal * epsilon
	fmt.Println(power)
}
