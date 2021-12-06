package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day6.txt
var input string

func main() {
	fmt.Printf("Part A: %d\n", partA(80))
	fmt.Printf("Part B: %d\n", partB(256))
}

func partA(iterations int) int {
	s := strings.Split(input, ",")
	for i := 0; i < iterations; i++ {
		for k, v := range s {
			c, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}

			if c == 0 {
				s[k] = "6"
				s = append(s, "8")
				continue
			} else {
				s[k] = strconv.Itoa(c - 1)
			}
		}
	}
	return len(s)
}

func partB(iterations int) int {
	s := strings.Split(input, ",")

	res := make(map[int]int)
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}

		res[i] += 1
	}

	for i := 0; i < iterations; i++ {
		created := make(map[int]int)
		for i := 0; i <= 8; i++ {

			k := i
			v, ok := res[k]
			if !ok {
				continue
			}

			if k == 0 {
				delete(res, k)
				created[6] += v
				created[8] += v
				continue
			}
			res[k] -= v
			res[k-1] += v
		}

		for k, v := range created {
			res[k] += v
		}
	}

	var sum int

	for _, v := range res {
		sum += v
	}

	return sum
}
