package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func Test_fillWaysToBeatTheRecord(t *testing.T) {
	tests := []struct {
		name  string
		input Race
		want  int
	}{
		{
			name: "actual",
			input: Race{
				time:   7,
				record: 9,
			},
			want: 4,
		},
		{
			name: "actual",
			input: Race{
				time:   15,
				record: 40,
			},
			want: 8,
		},
		{
			name: "actual",
			input: Race{
				time:   30,
				record: 200,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.fillWaysToBeatTheRecord(); got != tt.want {
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
			name:  "actual",
			input: inputTest,
			want:  288,
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
