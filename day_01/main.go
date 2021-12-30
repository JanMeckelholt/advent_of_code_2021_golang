package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var previous uint64
	log.Println(&previous)
	count := 0
	first := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if !first && next > previous {
			count++
		}
		first = false
		previous = next
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
