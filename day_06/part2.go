package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(path string) (finalNumber uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valuesAtStart := make([]uint, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := scanner.Text()
		valueStrs := strings.Split(inputString, ",")
		for _, valueStr := range valueStrs {
			value, err := strconv.ParseUint(valueStr, 10, 64)
			if err != nil {
				return 0, err
			}
			valuesAtStart = append(valuesAtStart, uint(value))
		}

	}

	duplicateFrequencyAtStart := getDuplicateFrequency(valuesAtStart)
	log.Printf("duplicateFrequencyAtStart: %v,\n", duplicateFrequencyAtStart)

	uniqueValuesAtStart := getUniqueValues(duplicateFrequencyAtStart)
	log.Printf("uniqueValuesAtStart: %v\n", uniqueValuesAtStart)
	valuesAfter128 := advanceX(uniqueValuesAtStart, duplicateFrequencyAtStart, 128)
	log.Printf("number of values after 128: %v\n", len(valuesAfter128))

	duplicateFrequencyAfter128 := getDuplicateFrequency(valuesAfter128)
	uniqueValuesAfter128 := getUniqueValues(duplicateFrequencyAfter128)

	count := uint64(0)
	for _, value := range uniqueValuesAfter128 {
		singleValue := make([]uint, 0)
		singleValue = append(singleValue, value)
		for i := uint(0); i < 128; i++ {
			//log.Printf("i: %v\n", i)
			singleValue = advanceOne(singleValue)
		}
		count += duplicateFrequencyAfter128[value] * uint64(len(singleValue))

	}
	return count, nil
}

func advanceX(uniqueValues []uint, duplicateFrequency map[uint]uint64, repetitions uint) []uint {
	valuesAfter := make([]uint, 0)
	for _, value := range uniqueValues {
		singleValue := make([]uint, 0)
		singleValue = append(singleValue, value)
		for i := uint(0); i < repetitions; i++ {
			//log.Printf("i: %v\n", i)
			singleValue = advanceOne(singleValue)
		}
		for j := uint64(0); j < duplicateFrequency[value]; j++ {
			//log.Printf("j: %v", j)
			valuesAfter = append(valuesAfter, singleValue...)
		}
	}
	return valuesAfter
}

func getUniqueValues(duplicateFrequencyAtStart map[uint]uint64) []uint {
	uniqueValues := make([]uint, 0)
	for value, _ := range duplicateFrequencyAtStart {
		uniqueValues = append(uniqueValues, value)
	}
	return uniqueValues
}
