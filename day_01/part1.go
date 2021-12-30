package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part1(path string) (uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var previous uint64
	count := uint64(0)
	first := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return 0, err
		}
		if !first && next > previous {
			count++
		}
		first = false
		previous = next
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}
