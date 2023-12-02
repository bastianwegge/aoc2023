package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func Test_process(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "actual",
			input: inputTest,
			want:  8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processAssignment(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Bag
	}{
		{
			name:  "blue assignment",
			input: "blue 4",
			want:  Bag{r: 0, g: 0, b: 4},
		},
		{
			name:  "red assignment",
			input: "red 4",
			want:  Bag{r: 4, g: 0, b: 0},
		},
		{
			name:  "green assignment",
			input: "green 4",
			want:  Bag{r: 0, g: 4, b: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processAssignment(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
