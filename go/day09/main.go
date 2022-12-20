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

	in, out := make(chan int, 10), make(chan int, 10)

	c := computer.NewComputer(commands, nil, in, out)
	in <- 1
	c.Compute()
	close(out)

	for x := range out {
		fmt.Println(x)
	}
}
