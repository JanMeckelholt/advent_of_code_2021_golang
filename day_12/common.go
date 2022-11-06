package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type Point = struct {
	Visited     bool
	Uppercase   bool
	Connections []string
}

func readIn(path string) (*map[string]Point, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var points = make(map[string]Point, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := scanner.Text()
		s := strings.Split(inputString, "-")
		if point, ok := points[s[0]]; ok {
			if strings.ToLower(s[1]) != "start" {
				point.Connections = append(point.Connections, s[1])
				points[s[0]] = point
			}
		} else if strings.ToLower(s[0]) != "end" && strings.ToLower(s[1]) != "start" {
			point := Point{
				Visited:     false,
				Uppercase:   unicode.IsUpper(rune(s[0][0])),
				Connections: []string{s[1]},
			}
			points[s[0]] = point
		}

		if point, ok := points[s[1]]; ok {
			if strings.ToLower(s[0]) != "start" {
				point.Connections = append(point.Connections, s[0])
				points[s[1]] = point
			}
		} else if strings.ToLower(s[1]) != "end" && strings.ToLower(s[0]) != "start" {
			point := Point{
				Visited:     false,
				Uppercase:   unicode.IsUpper(rune(s[1][0])),
				Connections: []string{s[0]},
			}
			points[s[1]] = point
		}

	}
	return &points, nil
}

func getNumberOfPathes(m map[string]Point, current Point, count int, twice bool) int {

	for _, dest := range current.Connections {
		if dest == "end" {
			count++
		} else {
			newPoint := m[dest]
			t := twice
			if twice && newPoint.Visited && !newPoint.Uppercase {
				//deadend -> no count increase
				continue
			}
			if !newPoint.Uppercase && newPoint.Visited {
				t = true
			}
			newPoint.Visited = true
			newMap := make(map[string]Point, len(m))
			for k, v := range m {
				newMap[k] = v
			}
			newMap[dest] = newPoint
			count = getNumberOfPathes(newMap, newPoint, count, t)
		}
	}
	return count
}
