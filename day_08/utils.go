package main

import "strings"

func getAllOutputs(valueMap map[string]string) []string {
	outputs := make([]string, 0, len(valueMap))
	for _, v := range valueMap {
		output := strings.Split(v, " ")
		outputs = append(outputs, output...)
	}
	return outputs
}

func count1478(valueArray []string) uint {
	count := uint(0)
	for _, v := range valueArray {
		if len(v) == len(One) || len(v) == len(Four) || len(v) == len(Seven) || len(v) == len(Eight) {
			count++
		}
	}
	return count
}
