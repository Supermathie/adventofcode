#!/usr/bin/env python3

import collections
import itertools

def move_cost1(crabs, pos):
    return sum(map(lambda x: abs(x - pos), crabs))

def part1(crabs):
    pos = sum(crabs) // len(crabs)
    cost = move_cost1(crabs, pos)
    while True:
        #print(f'{pos} - {cost}')
        new_cost = move_cost1(crabs, pos - 1)
        if new_cost < cost:
            pos, cost = pos - 1, new_cost
            continue

        new_cost = move_cost1(crabs, pos + 1)
        if new_cost < cost:
            pos, cost = pos + 1, new_cost
            continue

        return cost

def move_cost2(crabs, pos):
    return sum(map(lambda x: abs(x - pos)*(abs(x-pos)+1)//2, crabs))

def part2(crabs):
    pos = sum(crabs) // len(crabs)
    cost = move_cost2(crabs, pos)
    while True:
        #print(f'{pos} - {cost}')
        new_cost = move_cost2(crabs, pos - 1)
        if new_cost < cost:
            pos, cost = pos - 1, new_cost
            continue

        new_cost = move_cost2(crabs, pos + 1)
        if new_cost < cost:
            pos, cost = pos + 1, new_cost
            continue

        return cost

def read_data(name):
  with open(name, 'r') as f:
    return list(map(int, f.readline().split(',')))

def main():
   test_data = read_data('7-test.txt')
   data = read_data('7.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part2 test: {part2(test_data)}')

   print(f'part1 real: {part1(data)}')
   print(f'part2 real: {part2(data)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
