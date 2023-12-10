package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func Test_findStartPoint(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  Point
	}{
		{
			name:  "findStartPoint",
			input: []string{".....", ".S...", "....."},
			want:  Point{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStartPoint(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "findStartPoint",
			input: inputTest,
			want:  4,
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
