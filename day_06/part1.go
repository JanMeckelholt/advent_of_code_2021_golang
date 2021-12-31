package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(path string) (finalNumber uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := make([]uint64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := scanner.Text()
		valueStrs := strings.Split(inputString, ",")
		for _, valueStr := range valueStrs {
			value, err := strconv.ParseUint(valueStr, 10, 64)
			if err != nil {
				return 0, err
			}
			values = append(values, value)
		}

	}

	for i := 0; i < 80; i++ {
		values = advanceOne(values)
	}

	return uint64(len(values)), nil
}

func advanceOne(values []uint64) []uint64 {
	newValues := make([]uint64, 0)
	for _, value := range values {
		if value > 0 {
			value--
			newValues = append(newValues, value)
		} else if value == 0 {
			newValues = append(newValues, 6, 8)
		}
	}
	return newValues
}
