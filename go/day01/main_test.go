package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleFuel(t *testing.T) {
	tests := map[string]struct {
		input  int
		output int
	}{
		"one":   {input: 14, output: 2},
		"two":   {input: 1969, output: 966},
		"three": {input: 100756, output: 50346},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := moduleFuel(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}
