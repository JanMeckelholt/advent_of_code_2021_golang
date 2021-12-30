package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part2(path string) (uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := uint64(0)
	measurementNumber := uint64(0)
	measure1 := uint64(0)
	measure2 := uint64(0)
	measure3 := uint64(0)
	sumPrevious := uint64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		actual, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return 0, err
		}
		measurementNumber++
		log.Println(measurementNumber)
		if measurementNumber == 1 {
			log.Println("test")
			measure1 = actual
		}
		if measurementNumber == 2 {
			measure2 = actual
		}
		if measurementNumber == 3 {
			measure3 = actual
		}
		if measurementNumber == 3 {
			sumPrevious = measure1 + measure2 + measure3
		}
		if measurementNumber > 3 {
			measure1 = measure2
			measure2 = measure3
			measure3 = actual
			sumActual := measure1 + measure2 + measure3
			if sumActual > sumPrevious {
				count++
			}
			sumPrevious = sumActual
		}

	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}
