package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part2(path string) (oxygen uint64, co2 uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	oxygen = uint64(0)
	co2 = uint64(0)

	strings := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	//Find Oxygen
	position := 0
	oxygenStrings := strings
	for len(oxygenStrings) > 1 {
		count := int64(0)
		for _, reading := range oxygenStrings {
			count = stringToCount0(reading, position, count)
		}
		oxygenStrings = filterStrings(oxygenStrings, position, countToCriteria(count, "1"))
		position++
	}
	oxygen, err = strconv.ParseUint(oxygenStrings[0], 2, 64)
	if err != nil {
		return 0, 0, err
	}

	//Find CO2
	position = 0
	co2Strings := strings
	for len(co2Strings) > 1 {
		count := int64(0)
		for _, reading := range co2Strings {
			count = stringToCount0(reading, position, count)
		}
		co2Strings = filterStrings(co2Strings, position, countToCriteria(-count, "0"))
		position++
	}
	co2, err = strconv.ParseUint(co2Strings[0], 2, 64)
	if err != nil {
		return 0, 0, err
	}
	return oxygen, co2, nil
}

func filterStrings(strings []string, position int, criteria string) []string {
	returnStrings := make([]string, 0)
	for _, string := range strings {
		if string[position:position+1] == criteria {
			returnStrings = append(returnStrings, string)
		}
	}
	return returnStrings
}

func countToCriteria(count int64, valueWhenEqual string) string {
	if count == 0 {
		return valueWhenEqual
	}
	if count > 0 {
		return "0"
	}
	return "1"

}
