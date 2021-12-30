package main

import (
	"bufio"
	"log"
	"os"
)

func part2(path string) (sumUnmarked, finalNumber uint64, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputStrings := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputStrings = append(inputStrings, scanner.Text())
	}
	drawingNumbers, err := inputToDrawingNumbers(inputStrings)
	if err != nil {
		return 0, 0, err
	}

	boards, err := inputToBoards(inputStrings)
	if err != nil {
		return 0, 0, err
	}
	var winningBoards []int
	for _, number := range drawingNumbers {
		for ib, board := range boards {
			if !contains(winningBoards, ib) {
				for ir, row := range board {
					for in, n := range row {
						if n == number {
							boards[ib][ir][in] = 999
							if checkWin(boards[ib]) {
								if len(winningBoards)+1 == len(boards) {
									return getBoardSum(boards[ib]), number, nil
								}
								winningBoards = append(winningBoards, ib)
							}
						}

					}
				}
			}
		}
	}
	return 0, 0, nil
}
