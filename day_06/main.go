package main

import (
	"fmt"
	"log"
)

func main() {

	result1, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part1: %v\n", result1)

	result2, err2 := part2("./input.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Result Part2: %v\n", result2)

}
