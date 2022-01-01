package main

func advanceOne(values []uint) []uint {
	newValues := make([]uint, 0)
	for _, value := range values {
		if value > 0 {
			value--
			newValues = append(newValues, value)
		} else if value == 0 {
			newValues = append(newValues, 6, 8)
		}
	}
	return newValues
}

func getDuplicateFrequency(values []uint) map[uint]uint64 {

	duplicateFrequency := make(map[uint]uint64)

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
