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
	fmt.Printf("Result: \nNumber of Pathes: %v\n", numberOfPathes)

}
