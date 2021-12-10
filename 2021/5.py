#!/usr/bin/env python3

import collections
import itertools

def part1(vents):
  coverage = collections.defaultdict(lambda: 0)
  for c1, c2 in vents:
    if c1[0] == c2[0]:
      #print(f'vertical: {c1, c2}')
      x = c1[0]
      for y in range(min(c1[1], c2[1]), max(c1[1], c2[1])+1):
          coverage[x, y] += 1
    elif c1[1] == c2[1]:
      #print(f'horizontal: {c1, c2}')
      y = c1[1]
      for x in range(min(c1[0], c2[0]), max(c1[0], c2[0])+1):
          coverage[x, y] += 1
    else:
      pass
    #print(coverage)
  return sum(map(lambda v: v >= 2, coverage.values()))

def part2(vents):
  coverage = collections.defaultdict(lambda: 0)
  for c1, c2 in vents:
    if c1[0] == c2[0]:
      #print(f'vertical: {c1, c2}')
      x = c1[0]
      for y in range(min(c1[1], c2[1]), max(c1[1], c2[1])+1):
          coverage[x, y] += 1
    elif c1[1] == c2[1]:
      #print(f'horizontal: {c1, c2}')
      y = c1[1]
      for x in range(min(c1[0], c2[0]), max(c1[0], c2[0])+1):
          coverage[x, y] += 1
    else:
      x_step = 1 if c1[0] < c2[0] else -1
      y_step = 1 if c1[1] < c2[1] else -1
      for point in zip(
          range(c1[0], c2[0] + x_step, x_step),
          range(c1[1], c2[1] + y_step, y_step),
          ):
        coverage[point] += 1
    #print(coverage)
  return sum(map(lambda v: v >= 2, coverage.values()))

def read_data(name):
  with open(name, 'r') as f:
    def makepoint(p):
      return tuple(map(int, p.split(',')))

  return [ (makepoint(c1), makepoint(c2)) for c1, c2 in [ line.strip().split(' -> ') for line in f ] ]

def main():
   test_vents = read_data('5-test.txt')
   vents = read_data('5.txt')

   print(f'part1 test: {part1(test_vents)}')
   print(f'part2 test: {part2(test_vents)}')

   print(f'part1 real: {part1(vents)}')
   print(f'part2 real: {part2(vents)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
