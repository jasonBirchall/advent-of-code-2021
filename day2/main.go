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
