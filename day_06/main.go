package main

import (
	"fmt"
	"log"
)

func main() {

	result, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part1: %v\n", result)

}
