#!/usr/bin/python3

import copy
import re
import sys

class Almanac:

	def __init__(self):
		self.maps = {}
		self.map_order = []
		self.seeds = []
		self.input_min = None
		self.input_max = None

	def parse_input(self):

		for line in sys.stdin:
			if 'seeds' in line:
				self.seeds = [int(seed) for seed in line.split(':')[1].split()]

			elif 'map' in line:
				map_name = line.split()[0]
				self.maps[map_name] = []
				self.map_order.append(map_name)
			elif re.match('^[0-9 ]+$', line):
				destination, source, length = line.split()
				m = {
						'destination': int(destination),
						'source': int(source),
						'length': int(length),
						'source_max': int(source) + int(length)
					}

				if self.input_min is None or m['source'] < self.input_min:
					self.input_min = m['source']
				if self.input_max is None or m['source_max'] > self.input_max:
					self.input_max = m['source_max']
				self.maps[map_name].append(m)

		for map_name, maps in self.maps.items():
			self.maps[map_name] = sorted(maps, key=lambda m: m['source'])
		print(self.maps)

	def seed_location(self, seed_number):
		output = seed_number
		for map_name in self.map_order:
			#print(f'{output} {map_name} ', end='')
			a = 0
			z = len(self.maps[map_name]) - 1
			while a <= z:
				guess = int((z - a) / 2) + a
				#print(f'{a} {z} {guess}')
				curr_map = self.maps[map_name][guess]
				if output >= curr_map['source'] and output < (curr_map['source'] + curr_map['length']):
					output = curr_map['destination'] + (output - curr_map['source'])
					break
				elif output < curr_map['source']:
					z = guess - 1
				else:
					a = guess + 1
		#print(output)
		return(output)


if __name__ == '__main__':
	almanac = Almanac()
	almanac.parse_input()
	min_location = None
	for seed in almanac.seeds:
		location = almanac.seed_location(seed)
		if min_location is None or location < min_location:
			min_location = location
	print(f'Part 1 minimum location is {min_location}.')

	min_location = None
	seeds = copy.deepcopy(almanac.seeds)
	while(seeds):
		print(seeds)
		start_seed = seeds.pop(0)
		length = seeds.pop(0)
		for seed in range(start_seed, start_seed + length):
			if seed < almanac.input_min or seed > almanac.input_max:
				location = seed
			else:
				location = almanac.seed_location(seed)
			if min_location is None or location < min_location:
				min_location = location
	print(f'Part 2 minimum location is {min_location}.')
