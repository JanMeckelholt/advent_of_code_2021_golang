package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readIn(path string) (map[uint]map[uint]uint, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	valueMap := make(map[uint]map[uint]uint, 0)
	scanner := bufio.NewScanner(file)

	var row uint = 0
	for scanner.Scan() {
		inputString := scanner.Text()
		line := make(map[uint]uint, len(inputString))
		var i uint
		for i = 0; uint(len(inputString)) > i; i++ {

			v, _ := strconv.ParseUint(string(inputString[i]), 10, 64)
			line[i] = uint(v)
		}
		valueMap[row] = line
		row++
	}
	return valueMap, nil
}

func getRiskNumber(valueMap map[uint]map[uint]uint) uint {
	var riskNumber uint = 0
	for row, line := range valueMap {
		for index, v := range line {
			if leftHigher(line, index) && rightHigher(line, index) && aboveHigher(valueMap, row, index) && beloweHigher(valueMap, row, index) {
				riskNumber += uint(v) + 1
			}
		}
	}
	return riskNumber
}

func getLowPoints(valueMap map[uint]map[uint]uint) []Point {
	var lowPoints = make([]Point, 0)
	for row, line := range valueMap {
		for index, _ := range line {
			if leftHigher(line, index) && rightHigher(line, index) && aboveHigher(valueMap, row, index) && beloweHigher(valueMap, row, index) {
				lowPoints = append(lowPoints, Point{uint(index), uint(row)})
			}
		}
	}
	return lowPoints
}

func leftHigher(line map[uint]uint, index uint) bool {
	if index == 0 {
		return true
	}
	if line[index-1] > line[index] {
		return true
	}
	return false
}

func rightHigher(line map[uint]uint, index uint) bool {
	if index == uint(len(line)-1) {
		return true
	}
	if line[index+1] > line[index] {
		return true
	}
	return false
}

func aboveHigher(valueMap map[uint]map[uint]uint, row, index uint) bool {
	if row == 0 {
		return true
	}
	if valueMap[row-1][index] > valueMap[row][index] {
		return true
	}
	return false
}

func beloweHigher(valueMap map[uint]map[uint]uint, row, index uint) bool {
	if row == uint(len(valueMap))-1 {
		return true
	}
	if valueMap[row+1][index] > valueMap[row][index] {
		return true
	}
	return false
}

func remove(s []uint, i uint) []uint {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
