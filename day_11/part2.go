package main

func part2(path string) (iteration int, err error) {
	eMap, err := readIn(path)
	if err != nil {
		return 0, err
	}

	return getInterationZero(eMap), nil
}
