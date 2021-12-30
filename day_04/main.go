package main

import (
	"fmt"
	"log"
)

func main() {

	sumUnmarked, finalNumber, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result Part1: \nBoard-Sum: %v\nFinal Number: %v\n", sumUnmarked, finalNumber)
	fmt.Printf("Board-Sum * Final-Number: %v\n\n", sumUnmarked*finalNumber)

}
