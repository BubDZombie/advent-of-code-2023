#!/usr/bin/python3

import re
import sys

class Board:
	def __init__(self):
		self.rows = []
		self.gears = {}

	def add_row(self, row_string):
		row = []
		for char in row_string:
			row.append(char)
		self.rows.append(row)

	def check_adjacent(self, row, column):
		matches = []
		gears = []
		for check_row in [row - 1, row, row + 1]:
			for check_column in [column - 1, column, column + 1]:
				if (check_row >= 0
						and check_row < len(self.rows)
						and check_column >= 0
						and check_column < len(self.rows[check_row])
						and re.match('[^\.0-9]', self.rows[check_row][check_column])):
					char = self.rows[check_row][check_column]
					matches.append(f'{check_row}:{check_column}:{char}')
					print(f'match {check_row}:{check_column}:{char}')
					if char == '*':
						gears.append(f'{check_row}:{check_column}:{char}')
						print(f'gear {check_row}:{check_column}:{char}')
		return (matches, gears)

	def part_sum(self):
		sum = 0
		number_string = ''
		is_adjacent = False
		gear_set = set()
		for row_num, row in enumerate(self.rows):
			for col_num, char in enumerate(row):
				if re.match('[0-9]', char):
					number_string += char
					matches, gears = self.check_adjacent(row_num, col_num)
					if matches:
						is_adjacent = True
					for gear in gears:
						gear_set.add(gear)
				if not re.match('[0-9]', char) or col_num == (len(row) - 1):
					if number_string and is_adjacent:
						print(f'Adding symbol adjacent number {number_string} to sum.')
						sum += int(number_string)
					if number_string and gear_set:
						for gear in gear_set:
							print(gear)
							if gear not in self.gears:
								self.gears[gear] = [int(number_string)]
							else:
								self.gears[gear].append(int(number_string))
					number_string = ''
					gear_set = set()
					is_adjacent = False
		print(f'Part 1 sum is {sum}')

		gear_sum = 0
		for gear, numbers in self.gears.items():
			if len(numbers) == 2:
				ratio = numbers[0] * numbers[1]
				print(f'{gear} is a gear with ratio {ratio}.')
				gear_sum += ratio
		print(f'Part 2 gear sum is {gear_sum}')

if __name__ == '__main__':
	board = Board()
	for line in sys.stdin:
		board.add_row(line.strip())
	board.part_sum()
