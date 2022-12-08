package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var total, totalTwo int

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += fuel(mass)
		totalTwo += moduleFuel(mass)
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", totalTwo)
}

func moduleFuel(mass int) int {
	if mass <= 0 {
		return 0
	}

	return fuel(mass) + moduleFuel(fuel(mass))
}

func fuel(mass int) int {
	tmp := (mass / 3) - 2
	if tmp > 0 {
		return tmp
	}

	return 0
}
