package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed day7.txt
var input string

func main() {
	crabs := stringsToInts(strings.Split(input, ","))
	log.Println("Part 1:", part1(input, crabs))
	log.Println("Part 2:", partB(input, crabs))
}

func part1(input string, crabs []int) int {
	min := 9999999999

	ch := make(chan int, len(crabs))

	for i := 2; i <= len(crabs); i++ {
		go func(i int) {
			ch <- align(crabs, i)
		}(i)
	}

	for i := 2; i <= len(crabs); i++ {
		c := <-ch

		if c < min {
			min = c
		}
	}

	close(ch)

	return min
}

func partB(input string, crabs []int) int {
	min := 99999999999

	for i := 2; i <= len(crabs); i++ {
		c := alignTo(crabs, i)
		if c < min {
			min = c
		}
	}

	return min
}

func align(crabs []int, position int) int {
	c := 0

	for _, crab := range crabs {
		c += abs(position - crab)
	}

	return c
}

func alignTo(crabs []int, position int) int {
	c := 0

	for _, crab := range crabs {
		moves := abs(position - crab)
		c += int(float64(1+moves) / 2.0 * float64(moves))
	}

	return c
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func stringsToInts(s []string) []int {
	ints := []int{}
	for _, str := range s {
		mes, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}

		ints = append(ints, mes)
	}
	return ints
}
