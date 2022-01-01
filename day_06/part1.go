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

	values := make([]uint, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := scanner.Text()
		valueStrs := strings.Split(inputString, ",")
		for _, valueStr := range valueStrs {
			value, err := strconv.ParseUint(valueStr, 10, 64)
			if err != nil {
				return 0, err
			}
			values = append(values, uint(value))
		}

	}

	for i := 0; i < 80; i++ {
		values = advanceOne(values)
	}

	return uint64(len(values)), nil
}
