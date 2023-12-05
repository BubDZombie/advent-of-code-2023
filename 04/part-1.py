#!/usr/bin/python3

import sys

sum = 0
for line in sys.stdin:
	title, content = line.strip().split(': ')
	winners_string, numbers_string = content.split(' | ')
	winners = [int(n) for n in winners_string.split()]
	numbers = [int(n) for n in numbers_string.split()]

	card_score = 0
	for number in numbers:
		if number in winners:
			if card_score == 0:
				card_score = 1
			else:
				card_score *= 2
	sum += card_score
	print(f'{title} is worth {card_score} points.')

print(sum)
