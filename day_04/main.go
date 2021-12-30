package main

import (
	"fmt"
	"log"
)

func main() {

	sumUnmarked1, finalNumber1, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Result Part1: \nBoard-Sum: %v\nFinal Number: %v\n", sumUnmarked1, finalNumber1)
	fmt.Printf("Board-Sum * Final-Number: %v\n\n", sumUnmarked1*finalNumber1)

	sumUnmarked2, finalNumber2, err2 := part2("./input.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Result Part2: \nBoard-Sum: %v\nFinal Number: %v\n", sumUnmarked2, finalNumber2)
	fmt.Printf("Board-Sum * Final-Number: %v\n\n", sumUnmarked2*finalNumber2)

}
