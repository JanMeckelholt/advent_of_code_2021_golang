package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part1(path string) (count uint, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	valueMap := make(map[int]map[int]uint64, 0)
	scanner := bufio.NewScanner(file)

	row := 0
	for scanner.Scan() {
		inputString := scanner.Text()
		line := make(map[int]uint64, len(inputString))
		for i := 0; len(inputString) > i; i++ {

			v, _ := strconv.ParseUint(string(inputString[i]), 10, 64)
			line[i] = v
		}
		valueMap[row] = line
		row++
	}
	return getRiskNumber(valueMap), nil
}
