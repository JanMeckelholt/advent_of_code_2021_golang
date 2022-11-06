package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readIn(path string) (*map[int]map[int]int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	eMap := make(map[int]map[int]int, 0)
	row := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := make(map[int]int)
		inputString := scanner.Text()
		for i, v := range inputString {
			parsedV, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return &eMap, err
			}
			line[i] = int(parsedV)
		}
		eMap[row] = line
		row++
	}
	return &eMap, nil
}

func getNumberOfFlashes(eMap *map[int]map[int]int) (n int) {

	for i := 0; i < 100; i++ {
		numberOfFlashes := doOneRound(eMap)
		n += numberOfFlashes
	}
	printEMap(eMap)
	return n
}

func doOneRound(eMap *map[int]map[int]int) (numberOfFlashes int) {
	for row, line := range *eMap {
		for i, v := range line {
			(*eMap)[row][i] = v + 1
		}
	}
	for row, line := range *eMap {
		for i, v := range line {
			if v > 9 {
				(*eMap)[row][i] = 0
				numberOfFlashes++
				n := checkSurrounding(eMap, row, i)
				numberOfFlashes += n
			}
		}
	}

	return numberOfFlashes
}

func checkSurrounding(eMap *map[int]map[int]int, row, col int) (numberOfFlashes int) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) {
				if col+i >= 0 && col+i <= len((*eMap)[0])-1 {
					if row+j >= 0 && row+j <= len(*eMap)-1 {
						if (*eMap)[row+j][col+i] < 9 && (*eMap)[row+j][col+i] != 0 {
							(*eMap)[row+j][col+i] = (*eMap)[row+j][col+i] + 1
						} else {
							if (*eMap)[row+j][col+i] >= 9 {
								(*eMap)[row+j][col+i] = 0
								numberOfFlashes++
								n := checkSurrounding(eMap, row+j, col+i)
								numberOfFlashes += n
							}
						}
					}
				}
			}
		}

	}
	return numberOfFlashes
}

func printEMap(eMap *map[int]map[int]int) {
	fmt.Println("**************")
	for i := 0; i < len(*eMap); i++ {
		sLine := ""
		line := (*eMap)[i]
		for j := 0; j < len(line); j++ {
			sLine += fmt.Sprint(line[j])
		}
		fmt.Printf("%s\n", sLine)

	}
	fmt.Println("**************")
}
