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
