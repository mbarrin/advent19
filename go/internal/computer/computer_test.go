package computer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	tests := map[string]struct {
		commands []int
		output   []int
	}{
		"one": {
			commands: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			output:   []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		"two": {
			commands: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			output:   []int{1219070632396864},
		},
		"three": {
			commands: []int{104, 1125899906842624, 99},
			output:   []int{1125899906842624},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out := make(chan int, len(tc.commands))

			c := NewComputer(tc.commands, nil, nil, out)
			c.Compute()
			close(out)

			actual := []int{}
			for output := range out {
				actual = append(actual, output)
			}
			assert.Equal(t, tc.output, actual)
		})
	}

}
