package main

import (
	"bufio"
	"log"
	"os"
)

var openingCharacters = []rune{'(', '[', '<', '{'}
var closingCharacters = []rune{')', ']', '>', '}'}
var points = map[rune]uint{')': 3, ']': 57, '>': 25137, '}': 1197}

func readIn(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	navSystem := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := scanner.Text()
		navSystem = append(navSystem, inputString)
	}
	return navSystem, nil
}

func Count(line string, characters []rune) uint {
	var num uint
	for _, c := range line {
		if contains(characters, c) {
			num++
		}
	}
	return num
}

func contains(chars []rune, c rune) bool {
	for _, a := range chars {
		if a == c {
			return true
		}
	}
	return false
}

func getIndex(chars []rune, c rune) int {
	for i, a := range chars {
		if a == c {
			return i
		}
	}
	return -1
}

func checkLine(line string) rune {
	remainingLine := line
	for len(remainingLine) >= 2 {
		found := false
	ChangeRemainingLine:
		for i, c := range remainingLine {
			if contains(closingCharacters, c) {
				found = true
				if i > 0 && getIndex(openingCharacters, rune(remainingLine[i-1])) == getIndex(closingCharacters, c) {
					remainingLine = remainingLine[:i-1] + remainingLine[i+1:]
					break ChangeRemainingLine
				} else {
					return c
				}
			}
		}
		if !found {
			return '0'
		}
	}
	return '0'
}

func getScore(navSys []string) uint {
	var num uint
	for _, line := range navSys {
		firstWrongC := checkLine(line)
		if point, ok := points[firstWrongC]; ok {
			num += point
		}
	}
	return num
}
