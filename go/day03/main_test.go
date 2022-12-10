package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var firstLine = "R8,U5,L5,D3"
var first = map[point]int{
	{0, 1}: 1, {0, 2}: 2, {0, 3}: 3, {0, 4}: 4, {0, 5}: 5, {0, 6}: 6, {0, 7}: 7, {0, 8}: 8,
	{1, 8}: 9, {2, 8}: 10, {3, 8}: 11, {4, 8}: 12, {5, 8}: 13,
	{5, 7}: 14, {5, 6}: 15, {5, 5}: 16, {5, 4}: 17, {5, 3}: 18,
	{4, 3}: 19, {3, 3}: 20, {2, 3}: 21,
}

var secondLine = "U7,R6,D4,L4"
var second = map[point]int{
	{1, 0}: 1, {2, 0}: 2, {3, 0}: 3, {4, 0}: 4, {5, 0}: 5, {6, 0}: 6, {7, 0}: 7,
	{7, 1}: 8, {7, 2}: 9, {7, 3}: 10, {7, 4}: 11, {7, 5}: 12, {7, 6}: 13,
	{6, 6}: 14, {5, 6}: 15, {4, 6}: 16, {3, 6}: 17,
	{3, 5}: 18, {3, 4}: 19, {3, 3}: 20, {3, 2}: 21,
}

func TestBuild(t *testing.T) {
	tests := map[string]struct {
		input  string
		output map[point]int
	}{
		"line one": {input: firstLine, output: first},
		"line two": {input: secondLine, output: second},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := build(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestIntersections(t *testing.T) {
	tests := map[string]struct {
		inputLine     string
		inputExisting map[point]int
		output        map[point][]int
	}{
		"one": {
			inputLine: secondLine, inputExisting: first,
			output: map[point][]int{{5, 6}: {15, 15}, {3, 3}: {20, 20}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := intersections(tc.inputLine, tc.inputExisting)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestDistance(t *testing.T) {
	tests := map[string]struct {
		input  map[point][]int
		output int
	}{
		"one": {input: map[point][]int{{5, 6}: {15, 15}, {3, 3}: {20, 20}}, output: 30},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := distance(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestNearest(t *testing.T) {
	tests := map[string]struct {
		input  map[point][]int
		output int
	}{
		"one": {input: map[point][]int{{5, 6}: {15, 15}, {3, 3}: {20, 20}}, output: 6},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := nearest(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestManhattenDistance(t *testing.T) {
	tests := map[string]struct {
		inputRow int
		inputCol int
		output   int
	}{
		"one": {inputRow: 3, inputCol: 3, output: 6},
		"two": {inputRow: -3, inputCol: 4, output: 7},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := manhattenDistance(tc.inputRow, tc.inputCol)
			assert.Equal(t, actual, tc.output)
		})
	}

}
