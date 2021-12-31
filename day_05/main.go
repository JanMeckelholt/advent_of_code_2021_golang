package main

import (
	"fmt"
	"log"
)

func main() {

	numberOfDuplicates, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result Part1: \nNumber of Duplicates: %v", numberOfDuplicates)

}
