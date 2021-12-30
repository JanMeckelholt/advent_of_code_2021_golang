package main

import (
	"fmt"
	"log"
)

func main() {

	horizontal1, depth1, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part1: \nHorizontal: %v\nDepth: %v\n", horizontal1, depth1)
	fmt.Printf("Horizontal * Depth: %v\n\n", horizontal1*depth1)

	horizontal2, depth2, err2 := part2("./input.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Result Part1: \nHorizontal: %v\nDepth: %v\n", horizontal2, depth2)
	fmt.Printf("Horizontal * Depth: %v\n\n", horizontal2*depth2)
}
