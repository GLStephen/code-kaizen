# puzzle location

https://adventofcode.com/2023/day/7

# critical BASIC elements

- Each hand wins an amount equal to its bid multiplied by its rank
- Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank

# determining RANK
- Each hand wins an amount equal to its bid multiplied by its rank
- where the weakest hand gets rank 1
- the second-weakest hand gets rank 2
- and so on up to the strongest hand.

Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

# sample

32T3K 765

T55J5 684

KK677 28

KTJJT 220

QQQJA 483

# known result

6440

# first blush algorithm

- determine base rank: one pair, two pair, three pair, four of a kind, full house, flush, straight, three of a kind, two pair, one pair, high card, etc.
- resolve ties: