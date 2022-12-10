package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cmdRegex = regexp.MustCompile(`^(\w)(\d+)$`)

type point struct {
	rowID, columnID int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	first := build(lines[0])
	crossed := intersections(lines[1], first)
	fmt.Println("part 1:", nearest(crossed))
	fmt.Println("part 2:", distance(crossed))
}

func distance(points map[point][]int) int {
	smallest := math.Inf(0)
	for _, distances := range points {
		local := 0
		for _, loc := range distances {
			local += loc
		}
		if float64(local) < smallest {
			smallest = float64(local)
		}
	}
	return int(smallest)
}

func nearest(points map[point][]int) int {
	smallest := math.Inf(0)

	for point := range points {
		distance := float64(manhattenDistance(point.rowID, point.columnID))
		if distance < smallest {
			smallest = distance
		}
	}

	return int(smallest)
}

func intersections(line string, existing map[point]int) map[point][]int {
	current := &point{0, 0}
	points := map[point][]int{}
	location := 1
	commands := strings.Split(line, ",")

	for _, x := range commands {
		cmd := cmdRegex.FindStringSubmatch(x)
		dir := cmd[1]
		num, _ := strconv.Atoi(cmd[2])

		switch dir {
		case "U":
			for i := 0; i < num; i++ {
				current = &point{current.rowID + 1, current.columnID}
				if existing[*current] > 0 {
					points[*current] = append(points[*current], existing[*current])
					points[*current] = append(points[*current], location)
				}
				location++
			}
		case "D":
			for i := 0; i < num; i++ {
				current = &point{current.rowID - 1, current.columnID}
				if existing[*current] > 0 {
					points[*current] = append(points[*current], existing[*current])
					points[*current] = append(points[*current], location)
				}
				location++
			}
		case "R":
			for i := 0; i < num; i++ {
				current = &point{current.rowID, current.columnID + 1}
				if existing[*current] > 0 {
					points[*current] = append(points[*current], existing[*current])
					points[*current] = append(points[*current], location)
				}
				location++
			}
		case "L":
			for i := 0; i < num; i++ {
				current = &point{current.rowID, current.columnID - 1}
				if existing[*current] > 0 {
					points[*current] = append(points[*current], existing[*current])
					points[*current] = append(points[*current], location)
				}
				location++
			}
		}
	}

	return points
}

func build(line string) map[point]int {
	points := map[point]int{}
	location := 1
	commands := strings.Split(line, ",")

	current := &point{0, 0}

	for _, x := range commands {
		cmd := cmdRegex.FindStringSubmatch(x)
		dir := cmd[1]
		num, _ := strconv.Atoi(cmd[2])

		switch dir {
		case "U":
			for i := 0; i < num; i++ {
				current = &point{current.rowID + 1, current.columnID}
				points[*current] = location
				location++
			}
		case "D":
			for i := 0; i < num; i++ {
				current = &point{current.rowID - 1, current.columnID}
				points[*current] = location
				location++
			}
		case "R":
			for i := 0; i < num; i++ {
				current = &point{current.rowID, current.columnID + 1}
				points[*current] = location
				location++
			}
		case "L":
			for i := 0; i < num; i++ {
				current = &point{current.rowID, current.columnID - 1}
				points[*current] = location
				location++
			}
		}
	}

	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattenDistance(row, col int) int {
	return abs(row) + abs(col)
}
