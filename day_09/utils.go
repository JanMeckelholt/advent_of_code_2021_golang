package main

func getRiskNumber(valueMap map[int]map[int]uint64) uint {
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

func leftHigher(line map[int]uint64, index int) bool {
	if index == 0 {
		return true
	}
	if line[index-1] > line[index] {
		return true
	}
	return false
}

func rightHigher(line map[int]uint64, index int) bool {
	if index == len(line)-1 {
		return true
	}
	if line[index+1] > line[index] {
		return true
	}
	return false
}

func aboveHigher(valueMap map[int]map[int]uint64, row, index int) bool {
	if row == 0 {
		return true
	}
	if valueMap[row-1][index] > valueMap[row][index] {
		return true
	}
	return false
}

func beloweHigher(valueMap map[int]map[int]uint64, row, index int) bool {
	if row == len(valueMap)-1 {
		return true
	}
	if valueMap[row+1][index] > valueMap[row][index] {
		return true
	}
	return false
}
