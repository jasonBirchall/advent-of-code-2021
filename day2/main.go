package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day2.txt
var input string

func main() {
	fmt.Printf("The final depth for part one is: %v\n", part1())

	fmt.Printf("The final depth for part two is: %v\n", part2())
}

func part1() int {
	horizontal := 0
	depth := 0

	moves := strings.Split(input, "\n")

	for _, move := range moves {
		sections := strings.Split(move, " ")
		value, err := strconv.Atoi(sections[1])
		if err != nil {
			log.Fatalln(err)
		}

		switch sections[0] {
		case "up":
			depth -= value
		case "down":
			depth += value
		case "forward":
			horizontal += value
		}
	}

	return horizontal * depth
}

func part2() int {
	horizontal := 0
	depth := 0
	aim := 0

	moves := strings.Split(input, "\n")

	for _, move := range moves {
		sections := strings.Split(move, " ")
		value, err := strconv.Atoi(sections[1])
		if err != nil {
			log.Fatalln(err)
		}

		switch sections[0] {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			horizontal += value
			depth += value * aim
		}
	}

	return horizontal * depth

}
