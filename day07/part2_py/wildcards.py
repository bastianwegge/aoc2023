from collections import Counter
from enum import Enum


class HandType(Enum):
    HIGH_CARD = "0"
    ONE_PAIR = "1"
    TWO_PAIRS = "2"
    THREE_OF_A_KIND = "3"
    FULL_HOUSE = "4"
    FOUR_OF_A_KIND = "5"
    FIVE_OF_A_KIND = "6"


def handle_joker(hand):
    if hand == "JJJJJ":
        return "22222"

    counted_cards = Counter(hand)
    most_common_cards = counted_cards.most_common(2)

    first_most_common_card = most_common_cards[0][0]
    second_most_common_card = most_common_cards[1][0]

    if first_most_common_card != "J":
        return hand.replace("J", first_most_common_card)
    return hand.replace("J", second_most_common_card)


def get_hand_type(hand):
    has_joker = "J" in hand

    if has_joker:
        hand = handle_joker(hand)

    counted_cards = Counter(hand)
    different_cards = len(counted_cards)

    if different_cards == 5:
        return HandType.HIGH_CARD

    if different_cards == 4:
        return HandType.ONE_PAIR

    if different_cards == 1:
        return HandType.FIVE_OF_A_KIND

    if different_cards == 3:
        if sorted(counted_cards.values()) == [1, 1, 3]:
            return HandType.THREE_OF_A_KIND
        return HandType.TWO_PAIRS
    if different_cards == 2:
        if sorted(counted_cards.values()) == [2, 3]:
            return HandType.FULL_HOUSE
        return HandType.FOUR_OF_A_KIND


def create_ranking(input_games):
    type_ranking = []
    for game in input_games:
        hand, bid = game.split(" ")
        type = get_hand_type(hand)
        type_ranking.append((type.value + hand, int(bid)))

    return sorted(
        type_ranking, key=lambda x: x[0].translate(str.maketrans("TJQKA", f"A0BCD"))
    )


def process(input_games: str):
    ranking = create_ranking(input_games)
    total = sum([index * ranking[index - 1][1] for index in range(1, len(ranking) + 1)])
    return total


if "__main__" == __name__:
    with open("input.txt") as f:
        input = f.read().splitlines()
    print(process(input))