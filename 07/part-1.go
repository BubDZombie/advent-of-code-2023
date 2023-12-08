package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Bid int
	Cards CardCountList
	CardOrder []int
	Power int
}

type HandList []Hand
func (h HandList) Len() int { return len(h) }
func (h HandList) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h HandList) Less(i, j int) bool {
	if h[i].Power < h[j].Power {
		return true
	} else if h[i].Power == h[j].Power {
		for k, _ := range h[i].CardOrder {
			if h[i].CardOrder[k] < h[j].CardOrder[k] {
				return true
			} else if h[i].CardOrder[k] > h[j].CardOrder[k] {
				return false
			}
		}
	}
	return false
}

type CardCount struct {
	Value int
	Count int
}

type CardCountList []CardCount
func (c CardCountList) Len() int { return len(c) }
func (c CardCountList) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardCountList) Less(i, j int) bool {
	if (c[i].Count < c[j].Count) || ((c[i].Count == c[j].Count) && (c[i].Value < c[j].Value)) {
		return true
	}
	return false
}

func makeHand(cards string, bid int) Hand {
	values := map[string]int {
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	cardMap := make(map[int]int)
	var cardOrder []int
	for _, card := range cards {
		value := values[string(card)]
		_, ok := cardMap[value]
		if ok {
			cardMap[value]++
		} else {
			cardMap[value] = 1
		}
		cardOrder = append(cardOrder, value)
	}
	var counts CardCountList
	for value, count := range cardMap {
		counts = append(counts, CardCount{value, count})
	}
	sort.Sort(sort.Reverse(counts))
	var power int
	if counts[0].Count == 5 {
		power = 7
	} else if counts[0].Count == 4 && counts[1].Count == 1 {
		power = 6
	} else if counts[0].Count == 3 && counts[1].Count == 2 {
		power = 5
	} else if counts[0].Count == 3 && counts[1].Count == 1 {
		power = 4
	} else if counts[0].Count == 2 && counts[1].Count == 2 {
		power = 3
	} else if counts[0].Count == 2 && counts[1].Count == 1 {
		power = 2
	} else {
		power = 1
	}
	return Hand{bid, counts, cardOrder, power}
}

func main() {
	var hands HandList
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		line_split := strings.Fields(line)
		cards := line_split[0]
		bid, _ := strconv.Atoi(line_split[1])
		hands = append(hands, makeHand(cards, bid))
	}
	sort.Sort(hands)
	winnings := 0
	for rank, hand := range hands {
		fmt.Printf("%+v %i\n", hand, rank + 1)
		//fmt.Printf("%i * %i = %i, winnings = %i\n", hand.Bid, rank + 1, (rank + 1) * hand.Bid, winnings)
		winnings += (rank + 1) * hand.Bid
	}
	fmt.Printf("Part 1 solution: %i\n", winnings)
}
