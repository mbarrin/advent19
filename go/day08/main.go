package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type layer struct {
	data   []string
	digits map[string]int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	chars := strings.Split(string(input), "")

	layers := createLayers(chars, 25, 6)

	lowestVal := math.Inf(1)
	lowest := &layer{}
	for _, l := range layers {
		if float64(l.digits["0"]) < lowestVal {
			lowestVal = float64(l.digits["0"])
			lowest = l
		}
	}

	fmt.Println("part 1:", lowest.digits["1"]*lowest.digits["2"])

	for i, x := range image(layers).data {
		if i%25 == 0 {
			fmt.Println()
		}
		fmt.Print(x)
	}
}

func createLayers(chars []string, rows, cols int) []*layer {
	layers := []*layer{}
	for i := 0; i < len(chars); i = i + (rows * cols) {
		l := newLayer(chars[i : i+(rows*cols)])
		l.countDigits()
		layers = append(layers, l)
	}
	return layers
}

// 0 is black, 1 is white, and 2 is transparent.
func image(layers []*layer) *layer {
	imageSlice := []string{}

	for i := 0; i < len(layers[0].data); i++ {
		for _, x := range layers {
			if x.data[i] == "0" {
				imageSlice = append(imageSlice, " ")
				break
			} else if x.data[i] == "1" {
				imageSlice = append(imageSlice, "█")
				break
			}
		}
	}

	return newLayer(imageSlice)
}

func newLayer(image []string) *layer {
	return &layer{
		data:   image,
		digits: map[string]int{},
	}
}

func (l *layer) countDigits() {
	for _, x := range l.data {
		l.digits[x]++
	}
}
