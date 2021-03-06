package main

import (
	"bufio"
	"log"
	"os"
)

func part1(path string) (sumUnmarked, finalNumber uint64, err error) {
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

	for _, number := range drawingNumbers {
		for ib, board := range boards {
			for ir, row := range board {
				for in, n := range row {
					if n == number {
						boards[ib][ir][in] = 999
						if checkWin(boards[ib]) {
							log.Printf("Winner-Board:\n %v\n", boards[ib])
							log.Printf("Last Number: %v\n", number)
							return getBoardSum(boards[ib]), number, nil
						}
					}

				}
			}
		}
	}
	return 0, 0, nil
}
