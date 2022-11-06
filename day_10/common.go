package main

import (
	"bufio"
	"log"
	"os"
)

var openingCharacters = []rune{'(', '[', '<', '{'}
var closingCharacters = []rune{')', ']', '>', '}'}
var pointsCorrupt = map[rune]uint{')': 3, ']': 57, '>': 25137, '}': 1197}
var pointsIncomplete = map[rune]uint{'(': 1, '[': 2, '<': 4, '{': 3}

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

func checkLine(line string) (errorCharacter rune, remaining string) {
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
					return c, ""
				}
			}
		}
		if !found {
			return rune('0'), remainingLine
		}
	}
	return rune('0'), remainingLine
}

func sort(unsorted []uint) []uint {
	complete := false
	sorted := unsorted
	for !complete {
		changedSomething := false
		for i, v := range sorted {
			if len(sorted)-1 > i && sorted[i+1] < v {
				sorted[i] = sorted[i+1]
				sorted[i+1] = v
				changedSomething = true
			}
		}
		if !changedSomething {
			complete = true
		}
	}
	return sorted
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
