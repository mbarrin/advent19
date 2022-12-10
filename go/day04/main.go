package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	start, end := 137683, 596253
	total, totalTwo := 0, 0

	for i := start; i < end; i++ {
		total += valid(i, false)
		totalTwo += valid(i, true)
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", totalTwo)
}

func valid(number int, exact bool) int {
	pairsVal := map[int]int{}
	lowest := 0

	for _, x := range strings.Split(fmt.Sprintf("%d", number), "") {
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
