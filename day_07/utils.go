package main

func min(a, b uint64) uint64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b uint64) uint64 {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) uint {
	if a < 0 {
		return uint(-a)
	}
	return uint(a)
}

func fuelBetween(a, b uint) uint64 {
	fuel := uint64(0)
	x := min(uint64(a), uint64(b))
	y := max(uint64(a), uint64(b))
	for i := uint64(1); i <= y-x; i++ {
		fuel += i
	}
	return fuel
}

func MinMax(array []uint) (uint, uint) {
	var max uint = array[0]
	var min uint = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func getDuplicateFrequency(values []uint) map[uint]uint {
	duplicateFrequency := make(map[uint]uint)
	for _, value := range values {
		_, exists := duplicateFrequency[value]

		if exists {
			duplicateFrequency[value] += 1
		} else {
			duplicateFrequency[value] = 1
		}
	}
	return duplicateFrequency
}

func calcFuel(values []uint, position uint) uint64 {
	fuel := uint64(0)
	for _, value := range values {
		fuel += uint64(abs(int(value - position)))
	}
	return fuel
}

func calcFuel2(values []uint, position uint) uint64 {
	fuel := uint64(0)
	for _, value := range values {
		fuel += fuelBetween(value, position)
	}
	return fuel
}
