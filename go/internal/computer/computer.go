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

func (c *Computer) get(index int, mode byte) int {
	if mode == 48 {
		return c.commands[c.commands[index]]
	} else if mode == 49 {
		return c.commands[index]
	} else if mode == 50 {
		return c.commands[c.r+c.commands[index]]
	}
	return 0
}

func (c *Computer) set(index int, mode byte, value int) {
	if mode == 48 {
		c.commands[c.commands[index]] = value
	} else if mode == 50 {
		c.commands[c.r+c.commands[index]] = value
	}
}

func (c *Computer) add(firstMode, secondMode, thirdMode byte) {
	c.set(c.i+3, thirdMode, (c.get(c.i+1, firstMode) + c.get(c.i+2, secondMode)))
	c.i += 3
}

func (c *Computer) mult(firstMode, secondMode, thirdMode byte) {
	c.set(c.i+3, thirdMode, (c.get(c.i+1, firstMode) * c.get(c.i+2, secondMode)))
	c.i += 3
}

func (c *Computer) read(firstMode byte) {
	if len(c.inputs) == 0 {
		c.inputs = append(c.inputs, <-c.inputChannel)
	}
	c.set(c.i+1, firstMode, c.inputs[0])
	c.inputs = c.inputs[1:]
	c.i += 1
}

func (c *Computer) write(firstMode byte) {
	output := c.get(c.i+1, firstMode)
	c.outputs = append(c.outputs, output)
	c.outputChannel <- output
	c.i += 1
}

func (c *Computer) jmp(firstmode, secondMode byte) {
	if c.get(c.i+1, firstmode) != 0 {
		c.i = c.get(c.i+2, secondMode) - 1
	} else {
		c.i += 2
	}
}

func (c *Computer) jne(firstmode, secondMode byte) {
	if c.get(c.i+1, firstmode) == 0 {
		c.i = c.get(c.i+2, secondMode) - 1
	} else {
		c.i += 2
	}
}

func (c *Computer) lessThan(firstmode, secondMode, thirdMode byte) {
	if c.get(c.i+1, firstmode) < c.get(c.i+2, secondMode) {
		c.set(c.i+3, thirdMode, 1)
	} else {
		c.set(c.i+3, thirdMode, 0)
	}
	c.i += 3
}

func (c *Computer) equals(firstmode, secondMode, thirdMode byte) {
	if c.get(c.i+1, firstmode) == c.get(c.i+2, secondMode) {
		c.set(c.i+3, thirdMode, 1)
	} else {
		c.set(c.i+3, thirdMode, 0)
	}
	c.i += 3
}

func (c *Computer) Compute() int {
	for {
		// 0 pad to length 4 and split into an array
		cmd := fmt.Sprintf("%05d", c.commands[c.i])
		paramOneMode, paramTwoMode, paramThreeMode := cmd[2], cmd[1], cmd[0]

		// recreate the opcode
		switch string(cmd[3:]) {
		case "01":
			c.add(paramOneMode, paramTwoMode, paramThreeMode)
		case "02":
			c.mult(paramOneMode, paramTwoMode, paramThreeMode)
		case "03":
			c.read(paramOneMode)
		case "04":
			c.write(paramOneMode)
		case "05":
			c.jmp(paramOneMode, paramTwoMode)
		case "06":
			c.jne(paramOneMode, paramTwoMode)
		case "07":
			c.lessThan(paramOneMode, paramTwoMode, paramThreeMode)
		case "08":
			c.equals(paramOneMode, paramTwoMode, paramThreeMode)
		case "09":
			// update r
			c.r += c.get(c.i+1, paramOneMode)
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
