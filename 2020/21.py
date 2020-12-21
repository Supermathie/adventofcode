#!/usr/bin/env python3

import itertools
import re

allergen_map = {}

for line in open('input/21'):
		m = re.match('^(.*) \(contains (.*)\)$', line)
		ingredients = set(m.groups()[0].split())
		allergens = m.groups()[1].split(', ')

		for allergen in allergens:
			allergen_map[allergen] = allergen_map.get(allergen, ingredients) & ingredients

allergen_ingredients = set(itertools.chain(*allergen_map.values()))
total = 0

for line in open('input/21'):
		m = re.match('^(.*) \(contains (.*)\)$', line)
		ingredients = set(m.groups()[0].split())
		total += len(ingredients - allergen_ingredients)

print(f"Part 1: {total}")

while(any((len(x) > 1 for x in allergen_map.values()))):
	for k, v in allergen_map.items():
		if len(v) == 1:
			for k2, v2 in allergen_map.items():
				if k2 != k:
					v2.difference_update(v)

print(f"Part 2: {','.join((allergen_map[k].pop() for k in sorted(allergen_map)))}")
