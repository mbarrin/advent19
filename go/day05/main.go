package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ops []int

func main() {
	input, _ := os.ReadFile("input.txt")

	tmp := strings.Split(string(input), ",")

	var commands, commandsTwo ops

	for _, x := range tmp {
		code, _ := strconv.Atoi(x)
		commands = append(commands, code)
		commandsTwo = append(commandsTwo, code)
	}

	fmt.Println("part 1:", computer(commands, 1))
	fmt.Println("part 2:", computer(commandsTwo, 5))
}

func (o ops) value(index int, mode byte) int {
	if mode == 48 {
		return o[index]
	} else if mode == 49 {
		return index
	}
	return 0
}

func (o ops) add(i int, firstMode, secondMode byte) int {
	o[o[i+3]] = o[o.value(i+1, firstMode)] + o[o.value(i+2, secondMode)]
	return i + 3
}

func (o ops) mult(i int, firstMode, secondMode byte) int {
	o[o[i+3]] = o[o.value(i+1, firstMode)] * o[o.value(i+2, secondMode)]
	return i + 3
}

func (o ops) jmp(i int, firstmode, secondMode byte) int {
	if o[o.value(i+1, firstmode)] != 0 {
		i = o[o.value(i+2, secondMode)] - 1
	} else {
		i += 2
	}
	return i
}

func (o ops) jne(i int, firstmode, secondMode byte) int {
	if o[o.value(i+1, firstmode)] == 0 {
		i = o[o.value(i+2, secondMode)] - 1
	} else {
		i += 2
	}
	return i
}

func (o ops) lessThan(i int, firstmode, secondMode byte) int {
	if o[o.value(i+1, firstmode)] < o[o.value(i+2, secondMode)] {
		o[o[i+3]] = 1
	} else {
		o[o[i+3]] = 0
	}
	return i + 3
}

func (o ops) equals(i int, firstmode, secondMode byte) int {
	if o[o.value(i+1, firstmode)] == o[o.value(i+2, secondMode)] {
		o[o[i+3]] = 1
	} else {
		o[o[i+3]] = 0
	}
	return i + 3
}

func computer(commands ops, input int) []int {
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
			i = commands.add(i, paramOneMode, paramTwoMode)

		case "02":
			// mult
			i = commands.mult(i, paramOneMode, paramTwoMode)

		case "03":
			// read input
			commands[commands[i+1]] = input
			i += 1

		case "04":
			// write output
			param1 := commands.value(i+1, paramOneMode)
			output = append(output, commands[param1])
			i += 1

		case "05":
			// jump if true
			i = commands.jmp(i, paramOneMode, paramTwoMode)

		case "06":
			// jump if false
			i = commands.jne(i, paramOneMode, paramTwoMode)

		case "07":
			i = commands.lessThan(i, paramOneMode, paramTwoMode)

		case "08":
			// equal to
			i = commands.equals(i, paramOneMode, paramTwoMode)

		case "99":
			return output
		}
	}
	return output
}
