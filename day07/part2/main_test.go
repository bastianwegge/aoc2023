package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func Test_handInitialization(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  Hand
	}{
		{
			name:  "five of a kind",
			input: []string{"A", "A", "A", "A", "A"},
			want:  Hand{cards: []string{"A", "A", "A", "A", "A"}, cardCount: map[string]int{"A": 5}},
		},
		{
			name:  "four of a kind",
			input: []string{"A", "A", "A", "A", "3"},
			want:  Hand{cards: []string{"A", "A", "A", "A", "3"}, cardCount: map[string]int{"A": 4, "3": 1}},
		},
		{
			name:  "three of a kind",
			input: []string{"A", "A", "A", "5", "1"},
			want:  Hand{cards: []string{"A", "A", "A", "5", "1"}, cardCount: map[string]int{"A": 3, "5": 1, "1": 1}},
		},
		{
			name:  "two of a kind",
			input: []string{"A", "A", "1", "2", "4"},
			want:  Hand{cards: []string{"A", "A", "1", "2", "4"}, cardCount: map[string]int{"A": 2, "1": 1, "2": 1, "4": 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createHand(tt.input); reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HandCountWildcards(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "five of a kind",
			input: []string{"J", "A", "A", "J", "A"},
			want:  2,
		},
		{
			name:  "four of a kind",
			input: []string{"A", "A", "A", "A", "3"},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createHand(tt.input).CountWildcards(); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HandGetType(t *testing.T) {
	tests := []struct {
		name  string
		input Hand
		want  int
	}{
		{
			name:  "five of a kind",
			input: *createHand([]string{"A", "A", "A", "A", "A"}),
			want:  FiveOfAKind,
		},
		{
			name:  "four of a kind",
			input: *createHand([]string{"A", "A", "A", "A", "3"}),
			want:  FourOfAKind,
		},
		{
			name:  "three of a kind",
			input: *createHand([]string{"A", "A", "A", "5", "1"}),
			want:  ThreeOfAKind,
		},
		{
			name:  "one pair",
			input: *createHand([]string{"A", "A", "1", "2", "4"}),
			want:  OnePair,
		},
		{
			name:  "full house",
			input: *createHand([]string{"1", "1", "1", "2", "2"}),
			want:  FullHouse,
		},
		{
			name:  "full house 2",
			input: *createHand([]string{"1", "2", "1", "2", "1"}),
			want:  FullHouse,
		},
		{
			name:  "two pair",
			input: *createHand([]string{"1", "1", "4", "2", "2"}),
			want:  TwoPair,
		},
		{
			name:  "two pair",
			input: *createHand([]string{"1", "2", "4", "6", "9"}),
			want:  HighCard,
		},
		{
			name:  "three of a kind with a joker",
			input: *createHand([]string{"1", "2", "3", "3", "J"}),
			want:  ThreeOfAKind,
		},
		{
			name:  "five of a kind with three jokers",
			input: *createHand([]string{"A", "A", "J", "J", "J"}),
			want:  FiveOfAKind,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.BestHand().HandType(); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HandCompare(t *testing.T) {
	tests := []struct {
		name   string
		input1 Hand
		input2 Hand
		want   int
	}{
		{
			name:   "four of a kind wins",
			input1: *createHand([]string{"A", "2", "A", "A", "A"}),
			input2: *createHand([]string{"2", "3", "A", "A", "A"}),
			want:   WIN,
		},
		{
			name:   "four of a kind but input1 first card is higher",
			input1: *createHand([]string{"A", "2", "A", "A", "A"}),
			input2: *createHand([]string{"2", "A", "A", "A", "A"}),
			want:   WIN,
		},
		{
			name:   "four of a kind but input2 first card is higher",
			input1: *createHand([]string{"2", "A", "A", "A", "A"}),
			input2: *createHand([]string{"A", "A", "A", "A", "3"}),
			want:   LOSS,
		},
		{
			name:   "four of a kind but draw",
			input1: *createHand([]string{"A", "A", "A", "A", "3"}),
			input2: *createHand([]string{"A", "A", "A", "A", "3"}),
			want:   DRAW,
		},
		{
			name:   "two pair but input1 second card is higher",
			input1: *createHand([]string{"K", "K", "6", "7", "7"}),
			input2: *createHand([]string{"K", "T", "J", "J", "T"}),
			want:   WIN,
		},
		{
			name:   "two pair but input1 second card is higher",
			input1: *createHand([]string{"Q", "Q", "Q", "J", "A"}),
			input2: *createHand([]string{"T", "5", "5", "J", "5"}),
			want:   WIN,
		},
		{
			name:   "two pair but input1 second card is higher",
			input1: *createHand([]string{"Q", "Q", "Q", "J", "A"}),
			input2: *createHand([]string{"K", "T", "J", "J", "T"}),
			want:   WIN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input1.Compare(tt.input2); got != tt.want {
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
			want:  5905,
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
