package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type asteroid struct {
	rowID float64
	colID float64
}

type base struct {
	location asteroid
	count    int
	angles   map[float64]int
}

func main() {
	asteroids := loadData("input.txt")

	base := visible(asteroids)
	fmt.Println("part 1:", base.count)

	fmt.Println("part 2:", destroy(asteroids, base))
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

func destroy(asteroids map[asteroid]bool, base base) float64 {
	sortedKeys := []float64{}
	total := 0
	for k := range base.angles {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Float64s(sortedKeys)

	for _, a := range sortedKeys {
		base.angles[a] -= 1
		total++
		if base.angles[a] == 0 {
			delete(base.angles, a)
		}

		if total == 200 {
			return a
		}
	}

	return 0
}

func visible(asteroids map[asteroid]bool) base {
	max := base{}
	for a := range asteroids {
		lines := map[float64]int{}
		local := map[asteroid]bool{}

		for k, v := range asteroids {
			local[k] = v
		}

		delete(local, a)

		for l := range local {
			angle := math.Atan2((l.colID-a.colID), (l.rowID-a.rowID)) * (180 / math.Pi)
			angle = math.Mod((angle + 360), 360)
			lines[angle] += 1
		}

		if len(lines) > max.count {
			max.location = a
			max.count = len(lines)
			max.angles = lines
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
