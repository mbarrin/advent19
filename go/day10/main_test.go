package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadData(t *testing.T) {
	tests := map[string]struct {
		input  string
		output map[asteroid]bool
	}{
		"one": {
			input: "sample.txt",
			output: map[asteroid]bool{
				{rowID: 0, colID: 1}: true,
				{rowID: 0, colID: 4}: true,
				{rowID: 2, colID: 0}: true,
				{rowID: 2, colID: 1}: true,
				{rowID: 2, colID: 2}: true,
				{rowID: 2, colID: 3}: true,
				{rowID: 2, colID: 4}: true,
				{rowID: 3, colID: 4}: true,
				{rowID: 4, colID: 3}: true,
				{rowID: 4, colID: 4}: true,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := loadData(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestVisible(t *testing.T) {
	tests := map[string]struct {
		input  map[asteroid]bool
		output int
	}{
		"one": {
			input:  loadData("sample.txt"),
			output: 8,
		},
		"two": {
			input:  loadData("sample2.txt"),
			output: 210,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := visible(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}
