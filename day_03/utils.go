package main

func stringToCount0(inputString string, position int, counter int64) int64 {
	if inputString[position:position+1] == string('0') {
		return counter + 1
	} else {
		return counter - 1
	}
}
