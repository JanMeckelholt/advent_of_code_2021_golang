package main

import (
	"fmt"
	"log"
)

func main() {

	scorePart1, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nScore Part 1: %v\n", scorePart1)

	scorePart2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Score Part 2: %v\n", scorePart2)

}
