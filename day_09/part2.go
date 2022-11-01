package main

type Point struct {
	X uint
	Y uint
}

func (p Point) isInSlice(slice []Point) bool {
	for _, pInSlice := range slice {
		if pInSlice.X == p.X && pInSlice.Y == p.Y {
			return true
		}
	}
	return false
}

func (p Point) checkSourounding(valueMap map[uint]map[uint]uint) []Point {
	surrounding := make([]Point, 0)

	//left
	if p.X > 0 && valueMap[p.Y][p.X-1] != 9 {
		surrounding = append(surrounding, Point{p.X - 1, p.Y})
	}

	//right
	if uint(len(valueMap[p.Y]))-1 > p.X && valueMap[p.Y][p.X+1] != 9 {
		surrounding = append(surrounding, Point{p.X + 1, p.Y})
	}

	//above
	if p.Y > 0 && valueMap[p.Y-1][p.X] != 9 {
		surrounding = append(surrounding, Point{p.X, p.Y - 1})
	}

	//belowe
	if uint(len(valueMap))-1 > p.Y && valueMap[p.Y+1][p.X] != 9 {
		surrounding = append(surrounding, Point{p.X, p.Y + 1})
	}
	return surrounding
}

func getMax3Product(basinsLength []uint) uint {
	var minIndex, minValue uint
	max3Basins := make([]uint, len(basinsLength))
	max3Basins = append(max3Basins, basinsLength...)
	for len(max3Basins) > 3 {
		minValue = max3Basins[0]
		minIndex = 0
		for k, v := range max3Basins {
			if minValue > v {
				minIndex = uint(k)
				minValue = v
			}
		}
		max3Basins = remove(max3Basins, minIndex)
	}
	switch len(max3Basins) {
	case 1:
		return max3Basins[0]
	case 2:
		return max3Basins[0] * max3Basins[1]
	case 3:
		return max3Basins[0] * max3Basins[1] * max3Basins[2]
	}
	return 0
}

func part2(path string) (basinSum uint, err error) {
	valueMap, err := readIn(path)
	if err != nil {
		return 0, err
	}
	lowPoints := getLowPoints(valueMap)
	var basinsLength = make([]uint, 0)

	for _, lP := range lowPoints {
		basin := make([]Point, 0)
		basin = append(basin, lP)
		num := 0
		for len(basin) > num {
			num = len(basin)

			for _, p := range basin {
				surrounding := p.checkSourounding(valueMap)
				for _, pSurrounding := range surrounding {
					if !pSurrounding.isInSlice(basin) {
						basin = append(basin, pSurrounding)
					}
				}
			}
		}
		basinsLength = append(basinsLength, uint(len(basin)))
	}
	return getMax3Product(basinsLength), nil
}
