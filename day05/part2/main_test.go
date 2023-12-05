package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func Test_createRangeMap(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Conversion
	}{
		{
			name:  "creates Conversion from input",
			input: "50 98 2\n52 50 4",
			want: Conversion{
				SourceToDest: []SourceToDestination{
					{Source: 98, Dest: 50, Count: 2},
					{Source: 50, Dest: 52, Count: 4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createRangeMap(tt.input); reflect.DeepEqual(got, tt.want) != true {
				t.Errorf("got %v want %v", got, tt.want)
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
			want:  35,
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
