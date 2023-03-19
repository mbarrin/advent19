package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type asteroid struct {
	rowID float64
	colID float64
}

func main() {
	asteroids := loadData("input.txt")

	fmt.Println("part 1:", visible(asteroids))
}

func loadData(filename string) map[asteroid]bool {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	asteroids := map[asteroid]bool{}

	row := 0
	for scanner.Scan() {
		for col, field := range strings.Split(scanner.Text(), "") {
			if field == "#" {
				new := newAsteroid(float64(row), float64(col))
				asteroids[new] = true
			}
		}
		row++
	}

	return asteroids
}

func visible(asteroids map[asteroid]bool) int {
	max := 0
	for a := range asteroids {
		lines := map[float64]int{}
		local := map[asteroid]bool{}

		for k, v := range asteroids {
			local[k] = v
		}

		delete(local, a)

		for l := range local {
			angle := math.Atan2((l.colID-a.colID), (l.rowID-a.rowID)) * (180 / math.Pi)
			lines[angle] += 1
		}

		if len(lines) > max {
			max = len(lines)
		}
	}
	return max
}

func newAsteroid(row, col float64) asteroid {
	return asteroid{rowID: row, colID: col}
}

func (a *asteroid) String() string {
	return fmt.Sprintf("row: %f, col: %f", a.rowID, a.colID)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
