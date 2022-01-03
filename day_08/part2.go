package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func part2(path string) (sum uint, err error) {
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

	sum = uint(0)
	for k, v := range valueMap {
		allInputs := make([]string, 0)
		inputs := strings.Split(k, " ")
		outputs := strings.Split(v, " ")
		allInputs = append(allInputs, inputs...)

		wiringMap := findWiring(inputs)
		outputValue, err := getOutputValue(wiringMap, outputs)
		if err != nil {
			return 0, err
		}
		sum += outputValue
	}

	return sum, nil
}
