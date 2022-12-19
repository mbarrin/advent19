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

	fmt.Println("part 1:", maxSignal(commands, []int{0, 1, 2, 3, 4}))
	fmt.Println("part 2:", feedbackSignal(commands, []int{5, 6, 7, 8, 9}))
}

func feedbackSignal(commands, phases []int) int {
	signals := []int{}
	for _, phases := range permutations(phases) {
		wg := new(sync.WaitGroup)

		amps := []*computer.Computer{}

		ampACh := make(chan int, 5)
		ampBCh := make(chan int, 5)
		ampCCh := make(chan int, 5)
		ampDCh := make(chan int, 5)
		ampECh := make(chan int, 5)

		output := 0
		amp := computer.NewComputer(commands, wg, ampECh, ampACh)
		amp.SetInputs([]int{phases[0], output})
		amps = append(amps, amp)

		amp = computer.NewComputer(commands, wg, ampACh, ampBCh)
		amp.SetInputs([]int{phases[1]})
		amps = append(amps, amp)

		amp = computer.NewComputer(commands, wg, ampBCh, ampCCh)
		amp.SetInputs([]int{phases[2]})
		amps = append(amps, amp)

		amp = computer.NewComputer(commands, wg, ampCCh, ampDCh)
		amp.SetInputs([]int{phases[3]})
		amps = append(amps, amp)

		amp = computer.NewComputer(commands, wg, ampDCh, ampECh)
		amp.SetInputs([]int{phases[4]})
		amps = append(amps, amp)

		for _, x := range amps {
			wg.Add(1)
			go x.Compute()
		}

		wg.Wait()

		signals = append(signals, amps[4].LastOutput())
	}

	sort.Ints(signals)

	return signals[len(signals)-1]
}

func maxSignal(commands, phases []int) int {
	signals := []int{}

	for _, phases := range permutations(phases) {
		output := 0
		wg := new(sync.WaitGroup)

		in := make(chan int)
		out := make(chan int)

		amp := computer.NewComputer(commands, wg, in, out)
		amp.SetInputs(append([]int{phases[0]}, 0))
		wg.Add(1)
		go amp.Compute()

		output = <-out

		for i := 1; i < len(phases); i++ {
			amp := computer.NewComputer(commands, wg, in, out)

			amp.SetInputs(append([]int{phases[i]}, output))
			wg.Add(1)
			go amp.Compute()
			output = <-out
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
