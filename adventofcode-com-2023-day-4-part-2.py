#!/usr/bin/python3

import sys

copies = {}
for line in sys.stdin:
	title, content = line.strip().split(': ')
	card_number = int(title.split()[1])
	winners_string, numbers_string = content.split(' | ')
	winners = [int(n) for n in winners_string.split()]
	numbers = [int(n) for n in numbers_string.split()]

	if card_number in copies:
		copies[card_number] += 1
	else:
		copies[card_number] = 1

	print(f'Scratching card {card_number}')
	card_score = 0
	for number in numbers:
		if number in winners:
			card_score += 1

	for copy in range(card_number + 1, card_number + 1 + card_score):
		print(f'Adding a copy of {copy}')
		if copy in copies:
			copies[copy] += copies[card_number]
		else:
			copies[copy] = copies[card_number]

sum = 0
for number, count in copies.items():
	print(f'{number}: {count}')
	sum += count

print(sum)
