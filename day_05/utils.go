package main

import (
	"strconv"
	"strings"
)

func rowStrsToLines(inputStrings []string) ([]Line, error) {
	lines := make([]Line, 0, len(inputStrings))
	for _, rowStr := range inputStrings {
		line, err := rowStrToLine(rowStr)
		if err != nil {
			return lines, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func rowStrToLine(rowStr string) (Line, error) {
	var line Line
	pointStrings := strings.Split(rowStr, " -> ")
	point1, err := pointStrToPoint(pointStrings[0])
	if err != nil {
		return line, err
	}
	point2, err := pointStrToPoint(pointStrings[1])
	if err != nil {
		return line, err
	}
	return Line{point1, point2}, nil
}

func pointStrToPoint(pointStr string) (Point, error) {
	var point Point
	p := strings.Split(pointStr, ",")
	x1, err := strconv.ParseUint(p[0], 10, 64)
	y1, err := strconv.ParseUint(p[1], 10, 64)
	if err != nil {
		return point, err
	}
	return Point{x1, y1}, nil
}

func linesToPoints(lines []Line) []Point {
	points := make([]Point, 0)
	for _, line := range lines {
		ps := lineToPoints(line)
		points = append(points, ps...)
	}

	return points
}
func lineToPoints(line Line) []Point {
	points := make([]Point, 0)
	if line.a.x == line.b.x {
		start := min(line.a.y, line.b.y)
		end := max(line.a.y, line.b.y)
		for i := start; i <= end; i++ {
			points = append(points, Point{line.a.x, i})
		}
	} else if line.a.y == line.b.y {
		start := min(line.a.x, line.b.x)
		end := max(line.a.x, line.b.x)
		for i := start; i <= end; i++ {
			points = append(points, Point{i, line.a.y})
		}
	}

	return points
}

func min(a, b uint64) uint64 {
	if a < b {
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

func getDuplicateFrequency(points []Point) map[Point]uint64 {

	duplicateFrequency := make(map[Point]uint64)

	for _, point := range points {
		_, exists := duplicateFrequency[point]

		if exists {
			duplicateFrequency[point] += 1
		} else {
			duplicateFrequency[point] = 1
		}
	}

	return duplicateFrequency
}

func countDuplicates(duplicateFrequency map[Point]uint64) uint64 {
	count := uint64(0)
	for _, frequncy := range duplicateFrequency {

		if frequncy > 1 {
			count++
		}
	}

	return count
}
