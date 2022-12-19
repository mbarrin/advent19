package main

import (
	"advent19/internal/computer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputer(t *testing.T) {
	tests := map[string]struct {
		computer *computer.Computer
		input    []int
		output   int
	}{
		"input equals output":             {computer: computer.NewComputer([]int{3, 0, 4, 0, 99}), input: []int{666}, output: 666},
		"if 8: true":                      {computer: computer.NewComputer([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}), input: []int{8}, output: 1},
		"if 8: false":                     {computer: computer.NewComputer([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}), input: []int{7}, output: 0},
		"less than 8: true":               {computer: computer.NewComputer([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}), input: []int{7}, output: 1},
		"less than 8: false ":             {computer: computer.NewComputer([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}), input: []int{8}, output: 0},
		"if 8: true (immediate) ":         {computer: computer.NewComputer([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}), input: []int{8}, output: 1},
		"if 8: false (immediate)":         {computer: computer.NewComputer([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}), input: []int{7}, output: 0},
		"less than 8: true (immediate)":   {computer: computer.NewComputer([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}), input: []int{7}, output: 1},
		"less than 8: false (immediate)":  {computer: computer.NewComputer([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}), input: []int{8}, output: 0},
		"jump if false: true":             {computer: computer.NewComputer([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}), input: []int{0}, output: 0},
		"jump if false: false":            {computer: computer.NewComputer([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}), input: []int{6}, output: 1},
		"jump if true: true (immediate)":  {computer: computer.NewComputer([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}), input: []int{0}, output: 0},
		"jump if true: false (immediate)": {computer: computer.NewComputer([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}), input: []int{6}, output: 1},

		"large less than 8": {
			computer: computer.NewComputer([]int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			}),
			input:  []int{0},
			output: 999,
		},
		"large is 8": {
			computer: computer.NewComputer([]int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			}),
			input:  []int{8},
			output: 1000,
		},
		"large greater than 8": {
			computer: computer.NewComputer([]int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			}),
			input:  []int{16},
			output: 1001,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.computer.SetInputs(tc.input)
			go tc.computer.Compute()
			actual := <-tc.computer.OutputChannel()
			assert.Equal(t, tc.output, actual)
		})
	}
}
