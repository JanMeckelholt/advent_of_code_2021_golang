package main

func part1(path string) (score int, err error) {
	eMap, err := readIn(path)
	if err != nil {
		return 0, err
	}

	return getNumberOfFlashes(eMap), nil
}
