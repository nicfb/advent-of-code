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

	//rock > scissors > paper > rock
	myScore := 0
	for _, line := range input {
		turn := strings.Split(line, " ")
		them := turn[0]
		me := turn[1]
		myScore += play(them, me)
	}
	fmt.Println(myScore)
}

func play(them string, outcome string) int {
	// return part1(them, outcome)
	return part2(them, outcome)
}

func part1(them string, me string) int {
	win := 6
	draw := 3

	rock := 1
	paper := 2
	scissors := 3

	if them == "A" {
		if me == "X" {
			return draw + rock
		} else if me == "Y" {
			return win + paper
		} else if me == "Z" {
			return scissors
		}
	} else if them == "B" {
		if me == "X" {
			return rock
		} else if me == "Y" {
			return draw + paper
		} else if me == "Z" {
			return win + scissors
		}
	} else {
		if me == "X" {
			return win + rock
		} else if me == "Y" {
			return paper
		} else if me == "Z" {
			return draw + scissors
		}
	}
	return 0
}

func part2(them string, outcome string) int {
	win := 6
	draw := 3

	rock := 1
	paper := 2
	scissors := 3

	if outcome == "X" {
		//lose
		if them == "A" {
			return scissors
		} else if them == "B" {
			return rock
		} else if them == "C" {
			return paper
		}
	} else if outcome == "Y" {
		//draw
		if them == "A" {
			return draw + rock
		} else if them == "B" {
			return draw + paper
		} else if them == "C" {
			return draw + scissors
		}
	} else if outcome == "Z" {
		//win
		if them == "A" {
			return win + paper
		} else if them == "B" {
			return win + scissors
		} else if them == "C" {
			return win + rock
		}
	}
	return 0
}
