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

	fmt.Println("part 1:", process(commands, 1))

	fmt.Println("part 2:", process(commands, 5))
}

func process(commands []int, val int) int {
	in := make(chan int, 10)
	out := make(chan int, 10)

	wg := new(sync.WaitGroup)

	computer := computer.NewComputer(commands, wg, in, out)
	computer.SetInputs([]int{val})

	wg.Add(1)
	go computer.Compute()

	wg.Wait()
	close(in)
	close(out)

	var output int

	for {
		tmp, ok := <-out
		if ok {
			output = tmp
		} else {
			break
		}
	}

	return output
}
