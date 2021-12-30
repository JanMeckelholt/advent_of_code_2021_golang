package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part1(path string) (gamma uint64, ypsilon uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	ypsilon = uint64(0)
	var countArray [12]int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sensorString := scanner.Text()
		for i, _ := range sensorString {
			countArray[i] = stringToCount0(sensorString, i, countArray[i])
		}
	}
	gammaString := countsToGammaString(countArray)
	ypsilonString := gammaStringToYpsilonString(gammaString)
	log.Printf("Gamma-String: %v\nYpsilon-String: %v\n", gammaString, ypsilonString)
	gamma, err = strconv.ParseUint(gammaString, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	ypsilon, err = strconv.ParseUint(ypsilonString, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	return gamma, ypsilon, nil
}

func countsToGammaString(countArray [12]int64) string {
	returnString := ""
	for i, _ := range countArray {
		returnString += countToString(countArray[i])
	}
	return returnString
}

func countToString(count int64) string {
	if count > 0 {
		return string('0')
	} else {
		return string('1')
	}
}

func gammaStringToYpsilonString(gammaString string) string {
	returnString := ""
	for _, character := range gammaString {
		if string(character) == "0" {
			returnString += "1"
		} else {
			returnString += "0"
		}
	}
	return returnString
}
