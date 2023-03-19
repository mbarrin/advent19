package main

import (
	"advent19/internal/computer"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	tmp := strings.Split(string(input), ",")

	commands := []int{}

	for _, x := range tmp {
		code, _ := strconv.Atoi(x)
		commands = append(commands, code)
	}

	fmt.Println("part 1:", run(commands, 1))
	fmt.Println("part 2:", run(commands, 2))
}

func run(commands []int, value int) int {
	in, out := make(chan int, 10), make(chan int, 10)

	c := computer.NewComputer(commands, nil, in, out)
	in <- value
	c.Compute()
	close(out)

	var output int
	for x := range out {
		output = x
	}

	return output
}
