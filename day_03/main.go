package main

import (
	"fmt"
	"log"
)

func main() {

	gamma, ypsilon, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result Part1: \nGamma: %v\nYpsilon: %v\n", gamma, ypsilon)
	fmt.Printf("Gamma * Ypsilon: %v\n\n", gamma*ypsilon)

	oxygen, co2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result Part2: \nOxygen: %v\nCO2: %v\n", oxygen, co2)
	fmt.Printf("Oxygen * CO2: %v\n\n", oxygen*co2)
}
