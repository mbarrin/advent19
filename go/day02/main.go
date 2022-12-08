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

	commands[1] = 12
	commands[2] = 2

	commands = computer(commands)

	fmt.Println("part 1:", commands[0])
}

func computer(commands []int) []int {
	for i := 0; i <= len(commands); i++ {
		switch commands[i] {
		case 1:
			commands[commands[i+3]] = commands[commands[i+1]] + commands[commands[i+2]]
			i += 3
		case 2:
			commands[commands[i+3]] = commands[commands[i+1]] * commands[commands[i+2]]
			i += 3
		case 99:
			return commands
		}
	}

	return []int{}
}
