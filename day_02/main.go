package main

import (
	"fmt"
	"log"
)

func main() {

	horizontal, depth, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result Part1: \nHorizontal: %v\nDepth: %v\n", horizontal, depth)
	fmt.Printf("Horizontal * Depth: %v\n\n", horizontal*depth)

	/*	resultPart2, err := part2("./input.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Result Part2: %v \n", resultPart2)*/
}
