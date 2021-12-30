package main

import (
	"fmt"
	"log"
)

func main() {

	resultPart1, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result Part1: %v \n", resultPart1)

	resultPart2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result Part2: %v \n", resultPart2)
}
