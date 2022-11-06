package main

import (
	"fmt"
	"log"
)

func main() {

	numberOfFlashes, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nNumber of Flashes: %v\n", numberOfFlashes)

	// scorePart2, err := part2("./input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Score Part 2: %v\n", scorePart2)

}
