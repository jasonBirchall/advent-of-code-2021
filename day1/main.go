package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	depths, err := readDepths(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The number of times the depth measurement increases is: ", measureSingleDepth(depths)) // task 1

	fmt.Println("Considering the sums of a three-measurement sliding window, the number of times the depth measurement increases is: ", measureMultipleDepths(depths)) // task 2
}

func measureSingleDepth(depths []int) int {
	var count int
	for i := 1; i < len(depths); i++ {

		if depths[i] > depths[i-1] {
			count++
		}
	}

	return count
}

func measureMultipleDepths(depths []int) int {
	var count int

	for i := 0; i+3 < len(depths); i++ {
		a := depths[i] + depths[i+1] + depths[i+2]
		b := depths[i+1] + depths[i+2] + depths[i+3]
		if b > a {
			count++
		}
	}
	return count
}

func readDepths(r io.Reader) ([]int, error) {
	var depths []int
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		dep := strings.TrimSpace(scanner.Text())
		if dep == "" {
			continue
		}

		n, err := strconv.Atoi(dep)
		if err != nil {
			log.Fatalln(err)
		}

		depths = append(depths, n)
	}

	return depths, scanner.Err()
}
