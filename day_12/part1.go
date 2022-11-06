package main

func part1(path string) (numberOfPathes int, err error) {
	points, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return getNumberOfPathes(*points, (*points)["start"], 0), nil
}
