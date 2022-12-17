package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllOrbits(t *testing.T) {
	tests := map[string]struct {
		input  *planet
		output int
	}{
		"single": {input: &planet{name: "a"}, output: 0},
		"double": {input: &planet{name: "a", children: []*planet{{name: "b"}}}, output: 1},
		"triple": {
			input: &planet{
				name: "a",
				children: []*planet{
					{
						name: "b",
						children: []*planet{
							{name: "c"},
						},
					},
				},
			},
			output: 3,
		},
		"split 1": {
			input: &planet{
				name: "a",
				children: []*planet{
					{name: "b"}, {name: "c"},
				},
			},
			output: 2,
		},
		"split 2": {
			input: &planet{
				name: "a",
				children: []*planet{
					{
						name: "b", children: []*planet{
							{name: "c"},
							{name: "d"},
						},
					},
				},
			},
			output: 5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tc.input.findAllOrbits()
			assert.Equal(t, tc.output, actual)
		})
	}
}
