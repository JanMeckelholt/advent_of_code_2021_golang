package main

func part1(path string) (score uint, err error) {
	navSys, err := readIn(path)
	if err != nil {
		return 0, err
	}

	return getScore(navSys), nil
}
