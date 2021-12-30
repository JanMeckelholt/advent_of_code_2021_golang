package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
	log.Printf("Drawing numbers: %v\n", drawingNumbers)

	boards, err := inputToBoards(inputStrings)
	if err != nil {

	}

	log.Printf("boards: %v", boards)

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
	log.Printf("Boards after Drawing: %v", boards)

	return 0, 0, nil
}

func checkWin(board [5][5]uint64) bool {
	win := checkWinRow(board)
	if !win {
		win = checkWinRow(transpose(board))
	}
	return win
}

func checkWinRow(board [5][5]uint64) bool {
	for _, row := range board {
		noWin := false
		for _, n := range row {
			if n != 999 {
				noWin = true
			}
		}
		if !noWin {
			return true
		}
	}
	return false
}

func transpose(board [5][5]uint64) [5][5]uint64 {
	var result [5][5]uint64
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			result[i][j] = board[j][i]
		}
	}
	return result
}

func inputToDrawingNumbers(inputStrings []string) ([]uint64, error) {

	var returnArray []uint64

	for _, s := range strings.Split(inputStrings[0], ",") {
		element, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return returnArray, err
		}
		returnArray = append(returnArray, element)
	}
	return returnArray, nil
}

func inputToBoards(inputStrings []string) ([][5][5]uint64, error) {
	rows := inputStrings[1:]
	boards := make([][5][5]uint64, 0)
	rowNumber := 0
	var board [5][5]uint64
	for _, row := range rows {
		if row == "" {
			if rowNumber != 0 {
				rowNumber = 0
				boards = append(boards, board)
			}
		} else {
			numbers := strings.Split(row, " ")
			numberOfSpaces := 0
			for i, number := range numbers {

				if number == "" {
					numberOfSpaces++
				} else {
					n, err := strconv.ParseUint(number, 10, 64)
					if err != nil {
						return nil, err
					}
					board[rowNumber][i-numberOfSpaces] = n
				}
			}
			rowNumber++
		}
	}
	return boards, nil
}

func getBoardSum(board [5][5]uint64) uint64 {
	sum := uint64(0)
	for _, row := range board {
		for _, n := range row {
			if n != 999 {
				sum += n
			}
		}
	}
	return sum
}
