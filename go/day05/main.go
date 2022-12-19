package main

import (
	"advent19/internal/computer"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	tmp := strings.Split(string(input), ",")

	commands := []int{}

	for _, x := range tmp {
		code, _ := strconv.Atoi(x)
		commands = append(commands, code)
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	computerOne := computer.NewComputer(1, commands, wg)
	computerOne.SetInputs([]int{1})

	computerTwo := computer.NewComputer(1, commands, wg)
	computerTwo.SetInputs([]int{5})

	go computerOne.Compute()
	var output int
	for output == 0 {
		output = <-*computerOne.OutputChannel()
	}
	fmt.Println("part 1:", output)

	go computerTwo.Compute()
	output = <-*computerTwo.OutputChannel()
	fmt.Println("part 2:", output)
}
