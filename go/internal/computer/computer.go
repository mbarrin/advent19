package computer

import (
	"fmt"
	"sync"
)

type Ops []int

type Computer struct {
	Commands      Ops
	inputs        []int
	outputs       []int
	i             int
	inputChannel  chan int
	outputChannel chan int
	wg            *sync.WaitGroup
}

func NewComputer(commands Ops, wg *sync.WaitGroup, in, out chan int) *Computer {
	local := make(Ops, len(commands))
	copy(local, commands)

	return &Computer{
		Commands:      local,
		inputs:        []int{},
		outputs:       []int{},
		i:             0,
		inputChannel:  in,
		outputChannel: out,
		wg:            wg,
	}
}

func (c *Computer) LastOutput() int {
	return c.outputs[len(c.outputs)-1]
}

func (c *Computer) InputChannel() chan int {
	return c.inputChannel
}

func (c *Computer) OutputChannel() chan int {
	return c.outputChannel
}

func (c *Computer) SetInputs(inputs []int) {
	c.inputs = inputs
}

func (c *Computer) value(index int, mode byte) int {
	if mode == 48 {
		return c.Commands[index]
	} else if mode == 49 {
		return index
	}
	return 0
}

func (c *Computer) add(firstMode, secondMode byte) {
	c.Commands[c.Commands[c.i+3]] = c.Commands[c.value(c.i+1, firstMode)] + c.Commands[c.value(c.i+2, secondMode)]
	c.i += 3
}

func (c *Computer) mult(firstMode, secondMode byte) {
	c.Commands[c.Commands[c.i+3]] = c.Commands[c.value(c.i+1, firstMode)] * c.Commands[c.value(c.i+2, secondMode)]
	c.i += 3
}

func (c *Computer) jmp(firstmode, secondMode byte) {
	if c.Commands[c.value(c.i+1, firstmode)] != 0 {
		c.i = c.Commands[c.value(c.i+2, secondMode)] - 1
	} else {
		c.i += 2
	}
}

func (c *Computer) jne(firstmode, secondMode byte) {
	if c.Commands[c.value(c.i+1, firstmode)] == 0 {
		c.i = c.Commands[c.value(c.i+2, secondMode)] - 1
	} else {
		c.i += 2
	}
}

func (c *Computer) lessThan(firstmode, secondMode byte) {
	if c.Commands[c.value(c.i+1, firstmode)] < c.Commands[c.value(c.i+2, secondMode)] {
		c.Commands[c.Commands[c.i+3]] = 1
	} else {
		c.Commands[c.Commands[c.i+3]] = 0
	}
	c.i += 3
}

func (c *Computer) equals(firstmode, secondMode byte) {
	if c.Commands[c.value(c.i+1, firstmode)] == c.Commands[c.value(c.i+2, secondMode)] {
		c.Commands[c.Commands[c.i+3]] = 1
	} else {
		c.Commands[c.Commands[c.i+3]] = 0
	}
	c.i += 3
}

func (c *Computer) Compute() int {
	for c.i < len(c.Commands) {
		// 0 pad to length 4 and split into an array
		cmd := fmt.Sprintf("%04d", c.Commands[c.i])
		paramOneMode := cmd[1]
		paramTwoMode := cmd[0]

		// recreate the opcode
		switch string(cmd[2:]) {
		case "01":
			// add
			c.add(paramOneMode, paramTwoMode)

		case "02":
			// mult
			c.mult(paramOneMode, paramTwoMode)

		case "03":
			// read input
			if len(c.inputs) == 0 {
				c.inputs = append(c.inputs, <-c.inputChannel)
			}
			c.Commands[c.Commands[c.i+1]] = c.inputs[0]
			c.inputs = c.inputs[1:]
			c.i += 1

		case "04":
			// write output
			param1 := c.value(c.i+1, paramOneMode)
			output := c.Commands[param1]
			c.outputs = append(c.outputs, output)
			c.outputChannel <- output
			c.i += 1

		case "05":
			// jump if true
			c.jmp(paramOneMode, paramTwoMode)

		case "06":
			// jump if false
			c.jne(paramOneMode, paramTwoMode)

		case "07":
			// less than
			c.lessThan(paramOneMode, paramTwoMode)

		case "08":
			// equal to
			c.equals(paramOneMode, paramTwoMode)

		case "99":
			c.wg.Done()
			return 0
		}

		c.i++
	}
	return -1
}
