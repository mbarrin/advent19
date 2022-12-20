package computer

import (
	"fmt"
	"sync"
)

type Ops []int

type Computer struct {
	commands      Ops
	inputs        []int
	outputs       []int
	i             int
	r             int
	inputChannel  chan int
	outputChannel chan int
	wg            *sync.WaitGroup
}

func NewComputer(commands Ops, wg *sync.WaitGroup, in, out chan int) *Computer {
	local := make(Ops, 10240)
	copy(local, commands)

	return &Computer{
		commands:      local,
		inputs:        []int{},
		outputs:       []int{},
		i:             0,
		r:             0,
		inputChannel:  in,
		outputChannel: out,
		wg:            wg,
	}
}

func (c *Computer) LastOutput() int {
	return c.outputs[len(c.outputs)-1]
}

func (c *Computer) SetInputs(inputs []int) {
	c.inputs = inputs
}

func (c *Computer) value(index int, mode byte) int {
	if mode == 48 {
		return c.commands[index]
	} else if mode == 49 {
		return index
	} else if mode == 50 {
		return c.r + c.commands[index]
	}
	return 0
}

func (c *Computer) add(firstMode, secondMode byte) {
	c.commands[c.commands[c.i+3]] = c.commands[c.value(c.i+1, firstMode)] + c.commands[c.value(c.i+2, secondMode)]
	c.i += 3
}

func (c *Computer) mult(firstMode, secondMode byte) {
	c.commands[c.commands[c.i+3]] = c.commands[c.value(c.i+1, firstMode)] * c.commands[c.value(c.i+2, secondMode)]
	c.i += 3
}

func (c *Computer) jmp(firstmode, secondMode byte) {
	if c.commands[c.value(c.i+1, firstmode)] != 0 {
		c.i = c.commands[c.value(c.i+2, secondMode)] - 1
	} else {
		c.i += 2
	}
}

func (c *Computer) jne(firstmode, secondMode byte) {
	if c.commands[c.value(c.i+1, firstmode)] == 0 {
		c.i = c.commands[c.value(c.i+2, secondMode)] - 1
	} else {
		c.i += 2
	}
}

func (c *Computer) lessThan(firstmode, secondMode byte) {
	if c.commands[c.value(c.i+1, firstmode)] < c.commands[c.value(c.i+2, secondMode)] {
		c.commands[c.commands[c.i+3]] = 1
	} else {
		c.commands[c.commands[c.i+3]] = 0
	}
	c.i += 3
}

func (c *Computer) equals(firstmode, secondMode byte) {
	if c.commands[c.value(c.i+1, firstmode)] == c.commands[c.value(c.i+2, secondMode)] {
		c.commands[c.commands[c.i+3]] = 1
	} else {
		c.commands[c.commands[c.i+3]] = 0
	}
	c.i += 3
}

func (c *Computer) Compute() int {
	for {
		// 0 pad to length 4 and split into an array
		cmd := fmt.Sprintf("%04d", c.commands[c.i])
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
			param1 := c.value(c.i+1, paramOneMode)
			c.commands[param1] = c.inputs[0]
			c.inputs = c.inputs[1:]
			c.i += 1

		case "04":
			// write output
			param1 := c.value(c.i+1, paramOneMode)
			output := c.commands[param1]
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

		case "09":
			// update r
			c.r += c.value(c.i+1, paramOneMode)
			c.i += 1

		case "99":
			if c.wg != nil {
				c.wg.Done()
			}
			return 0
		}

		c.i++
	}
}
