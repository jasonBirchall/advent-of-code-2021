package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// read the input file line by line and compare count the number of times a depth
// measurement increases from the previous measurement.
func main() {
	var inc int
	var depth int

	// read the input file line by line
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// convert the string to an integer
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}

		// if the depth is greater than the previous depth, increment the count
		if depth == 0 {
			depth = i
		}

		if i > depth {
			inc++
			depth = i
		} else {
			depth = i
		}
	}

	fmt.Println(inc)
}
