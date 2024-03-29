package main

func part1(path string) (riskNumber uint, err error) {
	valueMap, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return getRiskNumber(valueMap), nil
}
