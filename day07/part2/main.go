package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var Cards = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

type Hand struct {
	cards         []string
	cardCount     map[string]int
	pos           int
	possibleHands []Hand
}

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func createHand(cards []string) *Hand {
	hand := Hand{
		cards:     cards,
		cardCount: make(map[string]int),
	}
	hand.possibleHands = hand.PossibleHands()
	if len(hand.possibleHands) > 0 {
		fmt.Println(hand.possibleHands[len(hand.possibleHands)-1])
	}
	for _, card := range cards {
		hand.cardCount[card]++
	}
	return &hand
}

func (h *Hand) CollectHands() []Hand {
	if len(h.possibleHands) == 0 {
		return []Hand{*h}
	}
	hands := make([]Hand, 0)
	hands = append(hands, *h)
	for _, possibleHand := range h.possibleHands {
		for _, hand := range possibleHand.CollectHands() {
			hands = append(hands, hand)
		}
	}
	return hands
}

func (h *Hand) BestHand() *Hand {
	allHands := h.CollectHands()
	if len(allHands) == 1 {
		return h
	}
	sort.Slice(allHands, func(i, j int) bool {
		return allHands[i].Compare(allHands[j]) == WIN
	})
	return &h.possibleHands[0]
}

func (h *Hand) PossibleHands() []Hand {
	possibleHands := make([]Hand, 0)
	for cardIndex, myCard := range h.cards {
		if myCard == "J" {
			for _, card := range Cards {
				if card == "J" {
					continue
				}
				newCards := make([]string, len(h.cards))
				copy(newCards, h.cards)
				newCards[cardIndex] = card
				possibleHands = append(possibleHands, *createHand(newCards))
			}
		}
	}
	return possibleHands
}

func (h *Hand) HandType() int {
	threeOfKind := false
	twoOfKind := false
	twoPair := false
	for _, count := range h.cardCount {
		if count == 5 {
			return FiveOfAKind
		}
		if count == 4 {
			return FourOfAKind
		}
		if count == 3 {
			threeOfKind = true
		}
		if count == 2 {
			if twoOfKind {
				twoPair = true
			} else {
				twoOfKind = true
			}
		}
	}
	if threeOfKind && twoOfKind {
		return FullHouse
	}
	if threeOfKind {
		return ThreeOfAKind
	}
	if twoPair {
		return TwoPair
	}
	if twoOfKind {
		return OnePair
	}

	return HighCard
}

const (
	WIN = iota
	DRAW
	LOSS
)

func (h *Hand) Compare(other Hand) int {
	handType := h.BestHand().HandType()
	otherHandType := other.HandType()
	if handType < otherHandType {
		return WIN
	}
	if handType > otherHandType {
		return LOSS
	}
	for i, card := range h.cards {
		if indexOf(card, Cards) > indexOf(other.cards[i], Cards) {
			return WIN
		}
		if indexOf(card, Cards) < indexOf(other.cards[i], Cards) {
			return LOSS
		}
	}
	return DRAW
}

func (h *Hand) CountWildcards() int {
	return h.cardCount["J"]
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func process(input string) int {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0)
	bids := make([]int, 0)
	// reading and parsing the input to Hands
	for lineIndex, line := range lines {
		cardsAndBid := strings.Split(line, " ")
		hand := *createHand(strings.Split(cardsAndBid[0], ""))
		hand.pos = lineIndex
		hands = append(hands, hand)
		bid, err := strconv.Atoi(cardsAndBid[1])
		if err != nil {
			panic(err)
		}
		bids = append(bids, bid)
	}

	sort.Slice(hands, func(i, j int) bool {
		//fmt.Println(hands[i].cards, hands[j].cards, hands[i].Compare(hands[j]) == WIN)
		return hands[i].Compare(hands[j]) != WIN
	})

	sum := 0
	for handIndex, hand := range hands {
		fmt.Println(handIndex, hand.cards, bids[hand.pos], "*", (handIndex + 1))
		sum += bids[hand.pos] * (handIndex + 1)
	}

	return sum
}

func main() {
	fmt.Println("Starting day06/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
