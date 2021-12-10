#!/usr/bin/env python3

import collections
import itertools

def step_day(pop):
  new_pop = collections.defaultdict(lambda: 0)
  for age, num in pop.items():
    if age == 0:
      new_pop[8] += num # new fish!
      new_pop[6] += num # old fish!
    else:
      new_pop[age-1] += num
  return new_pop

def sim(pop, days):
  for _ in range(days):
    pop = step_day(pop)
  return sum(pop.values())


def read_data(name):
  pop = collections.defaultdict(lambda: 0)
  with open(name, 'r') as f:
    for fish in map(int, f.readline().split(',')):
      pop[fish] += 1
  return pop

def main():
   test_pop = read_data('6-test.txt')
   pop = read_data('6.txt')

   print(f'part1 test (18): {sim(test_pop, 18)}')
   print(f'part1 test (80): {sim(test_pop, 80)}')
   print(f'part2 test (256): {sim(test_pop, 256)}')

   print(f'part1 real: {sim(pop, 80)}')
   print(f'part2 real: {sim(pop, 256)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
