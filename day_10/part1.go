package main

func part1(path string) (score uint, err error) {
	navSys, err := readIn(path)
	if err != nil {
		return 0, err
	}

	return getScorePart1(navSys), nil
}

func getScorePart1(navSys []string) uint {
	var num uint
	for _, line := range navSys {
		firstWrongC, _ := checkLine(line)
		if point, ok := pointsCorrupt[firstWrongC]; ok {
			num += point
		}
	}
	return num
}
