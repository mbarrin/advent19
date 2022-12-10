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

	var commands, commandsTwo []int

	for _, x := range tmp {
		code, _ := strconv.Atoi(x)
		commands = append(commands, code)
		commandsTwo = append(commandsTwo, code)
	}

	fmt.Println("part 1:", computer(commands, 1))
	fmt.Println("part 2:", computer(commandsTwo, 5))
}

func value(commands *[]int, index int, mode byte) int {
	// byte 48 == 0 == valueAt
	// byte 49 == 1 == value
	if mode == 48 {
		return (*commands)[index]
	} else if mode == 49 {
		return index
	}
	return 0
}

func add(commands *[]int, index int, firstMode, secondMode byte) int {
	return (*commands)[value(commands, index+1, firstMode)] + (*commands)[value(commands, index+2, secondMode)]
}

func mult(commands *[]int, index int, firstMode, secondMode byte) int {
	return (*commands)[value(commands, index+1, firstMode)] * (*commands)[value(commands, index+2, secondMode)]
}

func computer(commands []int, input int) []int {
	output := []int{}
	for i := 0; i < len(commands); i++ {
		// 0 pad to length 4 and split into an array
		cmd := fmt.Sprintf("%04d", commands[i])
		paramOneMode := cmd[1]
		paramTwoMode := cmd[0]

		// recreate the opcode
		switch string(cmd[2:]) {
		case "01":
			// add
			commands[commands[i+3]] = add(&commands, i, paramOneMode, paramTwoMode)
			i += 3

		case "02":
			// mult
			commands[commands[i+3]] = mult(&commands, i, paramOneMode, paramTwoMode)
			i += 3

		case "03":
			// read input
			commands[commands[i+1]] = input
			i += 1

		case "04":
			// write output
			param1 := value(&commands, i+1, paramOneMode)
			output = append(output, commands[param1])
			i += 1

		case "05":
			// jump if true
			if commands[value(&commands, i+1, paramOneMode)] != 0 {
				i = commands[value(&commands, i+2, paramTwoMode)] - 1
			} else {
				i += 2
			}

		case "06":
			// jump if false
			if commands[value(&commands, i+1, paramOneMode)] == 0 {
				i = commands[value(&commands, i+2, paramTwoMode)] - 1
			} else {
				i += 2
			}

		case "07":
			// less than
			if commands[value(&commands, i+1, paramOneMode)] < commands[value(&commands, i+2, paramTwoMode)] {
				commands[commands[i+3]] = 1
			} else {
				commands[commands[i+3]] = 0
			}
			i += 3

		case "08":
			// equal to
			if commands[value(&commands, i+1, paramOneMode)] == commands[value(&commands, i+2, paramTwoMode)] {
				commands[commands[i+3]] = 1
			} else {
				commands[commands[i+3]] = 0
			}
			i += 3

		case "99":
			return output
		}
	}
	return output
}
