package main

import (
	"advent19/internal/computer"
	"fmt"
	"os"
	"sort"
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

	//fmt.Println("part 1:", maxSignal(commands, []int{0, 1, 2, 3, 4}))
	fmt.Println("part 2:", maxSignal(commands, []int{5, 6, 7, 8, 9}))
}

func feedbackSignal(commands, phases []int) int {
	for _, phases := range permutations(phases) {
		wg := new(sync.WaitGroup)

		amps := []*computer.Computer{}

		output := 0
		for i, x := range phases {
			amp := computer.NewComputer(i, commands, wg)
			amp.SetInputs([]int{x, output})
			amps = append(amps, amp)
		}

		amps[0].SetInputChannel(amps[4].OutputChannel())
		amps[0].SetOutputChannel(amps[1].InputChannel())

		amps[1].SetInputChannel(amps[0].OutputChannel())
		amps[1].SetOutputChannel(amps[2].InputChannel())

		amps[2].SetInputChannel(amps[1].OutputChannel())
		amps[2].SetOutputChannel(amps[3].InputChannel())

		amps[3].SetInputChannel(amps[2].OutputChannel())
		amps[3].SetOutputChannel(amps[4].InputChannel())

		amps[4].SetInputChannel(amps[3].OutputChannel())
		amps[4].SetOutputChannel(amps[0].InputChannel())

		for _, x := range amps {
			go x.Compute()
		}

		wg.Wait()
	}

	//sort.Ints(signals)

	//return signals[len(signals)-1]
	return 0
}

func maxSignal(commands, phases []int) int {
	signals := []int{}

	for _, phases := range permutations(phases) {
		output := 0
		wg := new(sync.WaitGroup)

		amp := computer.NewComputer(1, commands, wg)
		amp.SetInputs(append([]int{phases[0]}, 0))
		go amp.Compute()

		output = <-*amp.OutputChannel()

		for i := 1; i < len(phases); i++ {
			amp := computer.NewComputer(1, commands, wg)

			amp.SetInputs(append([]int{phases[i]}, output))
			go amp.Compute()
			output = <-*amp.OutputChannel()
			wg.Wait()
		}

		signals = append(signals, output)
	}

	sort.Ints(signals)

	return signals[len(signals)-1]
}

func permutations(nums []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(nums, len(nums))

	return res
}
