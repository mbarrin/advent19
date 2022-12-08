package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputer(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output int
	}{
		"one":   {input: []int{1, 0, 0, 0, 99}, output: 2},
		"two":   {input: []int{2, 3, 0, 3, 99}, output: 2},
		"three": {input: []int{2, 4, 4, 5, 99, 0}, output: 2},
		"four":  {input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, output: 30},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computer(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}
