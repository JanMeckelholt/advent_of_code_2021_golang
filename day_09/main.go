package main

import (
	"fmt"
	"log"
)

func main() {

	risknumber, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nRisknumber: %v\n", risknumber)

}
