package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func part1(path string) (count uint, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	valueMap := make(map[string]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := scanner.Text()
		valueStrs := strings.Split(inputString, " | ")
		valueMap[valueStrs[0]] = valueStrs[1]
	}
	outputs := getAllOutputs(valueMap)
	count = count1478(outputs)
	return count, nil
}
