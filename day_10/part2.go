package main

func part2(path string) (score uint, err error) {
	navSys, err := readIn(path)
	if err != nil {
		return 0, err
	}
	incompleteLines := getIncompleteLines(navSys)

	return getScorePart2(incompleteLines), nil
}

func getIncompleteLines(navSys []string) []string {
	var incompleteLines = make([]string, 0)
	for _, line := range navSys {
		firstWrongC, line := checkLine(line)
		if _, ok := pointsCorrupt[firstWrongC]; !ok {
			incompleteLines = append(incompleteLines, line)
		}
	}
	return incompleteLines
}

func getScorePart2(incompleteLines []string) uint {
	var scores = make([]uint, 0)
	for _, line := range incompleteLines {
		lineScore := uint(0)
		for _, c := range reverse(line) {
			lineScore = lineScore*5 + pointsIncomplete[c]
		}
		scores = append(scores, lineScore)
	}
	sortedScores := sort(scores)
	if len(sortedScores) > 0 {
		return sortedScores[len(sortedScores)/2]
	}
	return 0
}
