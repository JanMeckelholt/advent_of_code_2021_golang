package main

import (
	"fmt"
	"log"
)

func main() {

	score, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nScore: %v\n", score)

}
