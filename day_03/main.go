package main

import (
	"fmt"
	"log"
)

func main() {

	gamma1, ypsilon1, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part1: \nGamma: %v\nYpsilon: %v\n", gamma1, ypsilon1)
	fmt.Printf("Gamma * Ypsilon: %v\n\n", gamma1*ypsilon1)
}
