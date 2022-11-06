package main

import (
	"fmt"
	"log"
)

func main() {

	numberOfPathes, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nNumber of Pathes Part 1: %v\n", numberOfPathes)

	numberOfPathes2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nNumber of Pathes Part 2: %v\n", numberOfPathes2)

}
