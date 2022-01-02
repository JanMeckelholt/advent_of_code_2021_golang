package main

import (
	"fmt"
	"log"
)

func main() {

	position1, minFuel1, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part1:\nMinimal fuel: %v\nat Position: %v\n", minFuel1, position1)

	position2, minFuel2, err2 := part2("./input.txt")
	if err2 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part2:\nMinimal fuel: %v\nat Position: %v\n", minFuel2, position2)
}
