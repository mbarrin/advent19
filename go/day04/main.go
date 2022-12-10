package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	start := "137683"
	end := "596253"
	total, totalTwo := 0, 0

	s, _ := strconv.Atoi(start)
	e, _ := strconv.Atoi(end)

	for i := s; i < e; i++ {
		total += valid(fmt.Sprintf("%d", i), false)
		totalTwo += valid(fmt.Sprintf("%d", i), true)
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", totalTwo)
}

func valid(number string, exact bool) int {
	pairsVal := map[int]int{}
	lowest := 0

	for _, x := range strings.Split(number, "") {
		num, _ := strconv.Atoi(x)
		if num < lowest {
			return 0
		}
		lowest = num
		pairsVal[num]++
	}

	for _, v := range pairsVal {
		if exact {
			if v == 2 {
				return 1
			}
		} else {
			if v >= 2 {
				return 1
			}
		}
	}

	return 0
}
