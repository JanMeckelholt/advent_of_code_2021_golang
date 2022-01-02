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
		valueStrs := strings.Split(inputString, "|")
		//input := strings.Split(valueStrs[0]," ")
		//output := strings.Split(valueStrs[1]," ")
		valueMap[valueStrs[0]] = valueStrs[1]
		log.Println(valueStrs[1])

	}
	//log.Printf("Input: %v\n", valueMap)
	log.Printf("Input: %v\n", valueMap["cgaed gcdbfa gcfaed gfcde gadfceb cdbfeg acg eacf eabgd ca"])
	outputs := getAllOutputs(valueMap)
	count = count1478(outputs)
	return count, nil
}
