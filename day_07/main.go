package main

import (
	"fmt"
	"log"
)

func main() {

	position, minFuel, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part1:\nMinimal fuel: %v\nat Position: %v\n", minFuel, position)
}
