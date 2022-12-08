package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	tmp := strings.Split(string(input), ",")

	var commands []int

	for _, x := range tmp {
		code, _ := strconv.Atoi(x)
		commands = append(commands, code)
	}

	one := make([]int, len(commands))
	copy(one, commands)
	one[1], one[2] = 12, 2

	fmt.Println("part 1:", computer(one))

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			two := make([]int, len(commands))
			copy(two, commands)

			two[1], two[2] = i, j
			if computer(two) == 19690720 {
				fmt.Println("part 2:", 100*i+j)
				break
			}
		}
	}
}

func computer(commands []int) int {
	for i := 0; i <= len(commands); i++ {
		cmd := commands[i]
		param1 := commands[i+1]
		param2 := commands[i+2]
		output := commands[i+3]

		switch cmd {
		case 1:
			commands[output] = commands[param1] + commands[param2]
			i += 3
		case 2:
			commands[output] = commands[param1] * commands[param2]
			i += 3
		case 99:
			return commands[0]
		}
	}

	return 0
}
