#!/usr/bin/env python3

import collections
import itertools
from termcolor import colored

class CorruptError(Exception): pass
class IncompleteError(Exception): pass

openers = {
    '(': (')', 1),
    '[': (']', 2),
    '{': ('}', 3),
    '<': ('>', 4),
}

closers = {
    ')': ('(', 3),
    ']': ('[', 57),
    '}': ('{', 1197),
    '>': ('<', 25137),
}

def parse(data):
  stack = []
  for c in data:
    if c in openers.keys():
      stack.append(c)
    else:
      if stack.pop() != closers[c][0]:
        raise CorruptError(c)
  if len(stack) > 0:
    raise IncompleteError(stack)

def part1(data):
  score = 0
  for line in data:
    try:
      parse(line)
    except IncompleteError:
      pass
    except CorruptError as e:
      score += closers[e.args[0]][1]
    else:
      raise "uhoh? valid line?"
  return score

def part2(data):
  scores = []
  for line in data:
    try:
      parse(line)
    except IncompleteError as e:
      score = 0
      stack = e.args[0]
      for c in reversed(stack):
        score *= 5
        score += openers[c][1]
      scores.append(score)
    except CorruptError:
      pass
    else:
      raise "uhoh? valid line?"
  scores = sorted(scores)
  return scores[(len(scores)-1)//2]

def read_data(name):
  with open(name, 'r') as f:
    return [ line.strip() for line in f ]

def main():
   test_data = read_data('10-test.txt')
   data = read_data('10.txt')

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
