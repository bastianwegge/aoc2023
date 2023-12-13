package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

//
//func Test_findArrangements(t *testing.T) {
//	tests := []struct {
//		name  string
//		input string
//		want  int
//	}{
//		{
//			name:  "findArrangements",
//			input: inputTest,
//			want:  8,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := findArrangements(tt.input); got != tt.want {
//				t.Errorf("findArrangements() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_process(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "process",
			input: inputTest,
			want:  21,
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
