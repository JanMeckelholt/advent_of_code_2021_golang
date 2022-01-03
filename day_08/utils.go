package main

import (
	"errors"
	"log"
	"math"
	"strings"
)

func getAllOutputs(valueMap map[string]string) []string {
	outputs := make([]string, 0)
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

func allPosibilites() map[string]bool {
	return map[string]bool{"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true}
}

func findWiring(valueArray []string) map[string]string {
	map1 := make(map[string]string, 8)
	possibleAs := allPosibilites()
	possibleBs := allPosibilites()
	possibleCs := allPosibilites()
	possibleDs := allPosibilites()
	possibleEs := allPosibilites()
	possibleFs := allPosibilites()
	possibleGs := allPosibilites()
	for _, v := range valueArray {
		if len(v) == len(One) {
			for _, j := range Eight {
				if !strings.Contains(v, j) {
					delete(possibleCs, j)
					delete(possibleFs, j)
				}
			}
		}
		if len(v) == len(Four) {
			for _, j := range Eight {
				if !strings.Contains(v, j) {
					delete(possibleCs, j)
					delete(possibleBs, j)
					delete(possibleDs, j)
					delete(possibleFs, j)
				}
			}
		}
		if len(v) == len(Seven) {
			for _, j := range Eight {
				if !strings.Contains(v, j) {
					delete(possibleCs, j)
					delete(possibleFs, j)
					delete(possibleAs, j)
				}
			}
		}
	}
	for k, _ := range possibleCs {
		for j, _ := range possibleAs {
			if k == j {
				delete(possibleAs, j)
			}
		}
	}
	for k, _ := range possibleAs {
		map1["a"] = k
	}
	delete(possibleBs, map1["a"])
	delete(possibleCs, map1["a"])
	delete(possibleDs, map1["a"])
	delete(possibleEs, map1["a"])
	delete(possibleFs, map1["a"])
	delete(possibleGs, map1["a"])
	/*	sixOrNineOrZero1 := make(map[string]bool, 0)
		sixOrNineOrZero2 := make(map[string]bool, 0)
		sixOrNineOrZero3 := make(map[string]bool, 0)*/
	zeroOrNine1 := ""
	zeroOrNine2 := ""
	six := ""
	c := ""

	for _, v := range valueArray {
		if len(v) == len(Six) {
			count := 0
			for k, _ := range possibleCs {
				for _, j := range v {
					if k == string(j) {
						count++
					}
				}
			}
			if count == 1 {
				six = v
			} else if count == 2 {
				if len(zeroOrNine1) == 0 {
					zeroOrNine1 = v
				} else {
					zeroOrNine2 = v
				}
			}

		}
	}

	for k, _ := range possibleCs {
		found := false
		for _, j := range six {
			if k == string(j) {
				found = true
			}
		}
		if !found { // k == c
			c = k
		}
	}
	for _, j := range Eight {
		if j != c {
			delete(possibleCs, j)
		}
	}
	delete(possibleFs, c)

	for k, _ := range possibleFs {
		map1["f"] = k
	}
	for k, _ := range possibleCs {
		map1["c"] = k
	}
	delete(possibleBs, map1["c"])
	delete(possibleDs, map1["c"])
	delete(possibleEs, map1["c"])
	delete(possibleGs, map1["c"])

	delete(possibleBs, map1["f"])
	delete(possibleDs, map1["f"])
	delete(possibleEs, map1["f"])
	delete(possibleGs, map1["f"])

	//zero := ""
	nine := ""
	d := ""
	four := ""
	for _, v := range valueArray {
		if len(v) == len(Four) {
			four = v
		}
	}
	found := false
	for _, j := range four {
		if !found {
			exists1 := strings.Contains(zeroOrNine1, string(j))
			exists2 := strings.Contains(zeroOrNine2, string(j))
			if !exists1 {
				//zero = zeroOrNine1
				nine = zeroOrNine2
				found = true
				d = string(j)
			}
			if !exists2 {
				//zero = zeroOrNine2
				nine = zeroOrNine1
				found = true
				d = string(j)
			}
		}
	}

	for _, j := range Eight {
		if j != d {
			delete(possibleDs, j)
		}
	}
	for k, _ := range possibleDs {
		map1["d"] = k
	}
	delete(possibleBs, map1["d"])
	delete(possibleEs, map1["d"])
	delete(possibleGs, map1["d"])

	for _, j := range four {
		if string(j) != map1["d"] && string(j) != map1["c"] && string(j) != map1["f"] { //j==b
			for _, r := range Eight {
				if r != string(j) {
					delete(possibleBs, r)
				}
			}
		}
	}

	for k, _ := range possibleBs {
		map1["b"] = k
	}
	delete(possibleEs, map1["b"])
	delete(possibleGs, map1["b"])
	for _, k := range nine {
		if string(k) != map1["a"] && string(k) != map1["b"] && string(k) != map1["c"] && string(k) != map1["d"] && string(k) != map1["f"] { //string(k)==g
			for _, r := range Eight {
				if r != string(k) {
					delete(possibleGs, r)
				}
			}
		}
	}
	for k, _ := range possibleGs {
		map1["g"] = k
	}
	delete(possibleEs, map1["g"])
	for k, _ := range possibleEs {
		map1["e"] = k
	}
	result := make(map[string]string, 8)
	for k, v := range map1 {
		result[v] = k
	}
	return result
}

func getOutputValue(wiringMap map[string]string, outputs []string) (uint, error) {
	sum := uint(0)
	for k, output := range outputs {
		newString := convertStr(wiringMap, output)
		digit, err := strToDigit(newString)
		if err != nil {
			return 0, err
		}
		sum += uint(math.Pow(10, float64(3-k))) * digit
	}
	return sum, nil
}

func convertStr(wiringMap map[string]string, input string) string {
	newString := ""
	for _, v := range input {
		newString += wiringMap[string(v)]
	}
	return newString
}

func strToDigit(s string) (uint, error) {
	if checkDigit(s, Zero) {
		return 0, nil
	}
	if checkDigit(s, One) {
		return 1, nil
	}
	if checkDigit(s, Two) {
		return 2, nil
	}
	if checkDigit(s, Three) {
		return 3, nil
	}
	if checkDigit(s, Four) {
		return 4, nil
	}
	if checkDigit(s, Five) {
		return 5, nil
	}
	if checkDigit(s, Six) {
		return 6, nil
	}
	if checkDigit(s, Seven) {
		return 7, nil
	}
	if checkDigit(s, Eight) {
		return 8, nil
	}
	if checkDigit(s, Nine) {
		return 9, nil
	}
	log.Printf("___NOT FOUND____\n%v", s)
	return 0, errors.New("not Found")
}

func checkDigit(s string, digit []string) bool {
	if len(s) != len(digit) {
		return false
	}
	match := true
	for _, c := range digit {
		found := false
		for _, v := range s {
			if string(v) == c {
				found = true
			}
		}

		if !found {
			match = false
		}
	}
	if match {

		return true
	}
	return false
}
