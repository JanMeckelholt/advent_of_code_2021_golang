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

	interationZero, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of Itertions: %v\n", interationZero)

}
