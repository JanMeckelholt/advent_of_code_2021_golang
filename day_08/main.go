package main

import (
	"fmt"
	"log"
)

func main() {

	count, err1 := part1("./input.txt")
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Printf("Result Part1:\nDigits 1, 4, 7 and 8 appear %v times in output\n", count)

}
