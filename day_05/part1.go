package main

import (
	"bufio"
	"log"
	"os"
)

func part1(path string) (numberOfDuplicates uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputStrings := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputStrings = append(inputStrings, scanner.Text())
	}
	lines, err := rowStrsToLines(inputStrings)
	if err != nil {
		return 0, err
	}
	points := linesToPoints(lines)
	duplicateFrequency := getDuplicateFrequency(points)
	numberOfDuplicates = countDuplicates(duplicateFrequency)

	return numberOfDuplicates, nil
}
