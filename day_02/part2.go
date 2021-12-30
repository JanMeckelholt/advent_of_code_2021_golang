package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(path string) (horizontal uint64, depth uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	horizontal = uint64(0)
	depth = uint64(0)
	aim := uint64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction := strings.Fields(scanner.Text())
		step, err := strconv.ParseUint(direction[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		switch Direction(direction[0]) {
		case DirectionForward:
			horizontal += step
			depth += aim * step
		case DirectionUp:
			aim -= step
		case DirectionDown:
			aim += step
		default:
			return 0, 0, errors.New("invalid input file")
		}
	}

	return horizontal, depth, nil
}
