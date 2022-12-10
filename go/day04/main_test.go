package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	tests := map[string]struct {
		input  int
		exact  bool
		output int
	}{
		"one":   {input: 111111, exact: false, output: 1},
		"two":   {input: 223450, exact: false, output: 0},
		"three": {input: 123789, exact: false, output: 0},

		"four":  {input: 112233, exact: true, output: 1},
		"five":  {input: 123444, exact: true, output: 0},
		"six":   {input: 111122, exact: true, output: 1},
		"seven": {input: 223333, exact: true, output: 1},
		"eight": {input: 222333, exact: true, output: 0},
		"nine":  {input: 223311, exact: true, output: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := valid(tc.input, tc.exact)
			assert.Equal(t, tc.output, actual)
		})
	}
}
