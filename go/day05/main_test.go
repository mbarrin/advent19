package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputer(t *testing.T) {
	tests := map[string]struct {
		commands []int
		input    int
		output   []int
	}{
		"input equals output":             {commands: []int{3, 0, 4, 0, 99}, input: 666, output: []int{666}},
		"if 8: true":                      {commands: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 8, output: []int{1}},
		"if 8: false":                     {commands: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 7, output: []int{0}},
		"less than 8: true":               {commands: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: 7, output: []int{1}},
		"less than 8: false ":             {commands: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: 8, output: []int{0}},
		"if 8: true (immediate) ":         {commands: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 8, output: []int{1}},
		"if 8: false (immediate)":         {commands: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 7, output: []int{0}},
		"less than 8: true (immediate)":   {commands: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: 7, output: []int{1}},
		"less than 8: false (immediate)":  {commands: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: 8, output: []int{0}},
		"jump if false: true":             {commands: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: 0, output: []int{0}},
		"jump if false: false":            {commands: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: 6, output: []int{1}},
		"jump if true: true (immediate)":  {commands: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 0, output: []int{0}},
		"jump if true: false (immediate)": {commands: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 6, output: []int{1}},

		"large less than 8": {
			commands: []int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			},
			input:  0,
			output: []int{999},
		},
		"large is 8": {
			commands: []int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			},
			input:  8,
			output: []int{1000},
		},
		"large greater than 8": {
			commands: []int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			},
			input:  16,
			output: []int{1001},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computer(tc.commands, tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}
