package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(path string) (position uint, minFuel uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := make([]uint, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := scanner.Text()
		valueStrs := strings.Split(inputString, ",")
		for _, valueStr := range valueStrs {
			value, err := strconv.ParseUint(valueStr, 10, 64)
			if err != nil {
				return 0, 0, err
			}
			values = append(values, uint(value))
		}

	}
	minValue, maxValue := MinMax(values)
	duplicateFrequency := getDuplicateFrequency(values)
	minFuel = calcFuel(values, minValue)
	position = minValue
	for i := minValue; i <= maxValue; i++ {
		_, exists := duplicateFrequency[i]
		if exists {
			fuel := calcFuel(values, i)
			if fuel < minFuel {
				minFuel = fuel
				position = i
			}
		}
	}

	log.Printf("values: %v\n", values)

	return position, minFuel, nil
}
