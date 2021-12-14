#!/usr/bin/env python3

import collections
import itertools
import datetime

def do_step1(state, rules):
  new_state = ''
  for e1, e2 in zip(state, state[1:]):
    yield e1
    try:
      yield rules[(e1,e2)]
    except KeyError:
      pass
  yield state[-1]

def part1(data):
  state, rules = data
  for i in range(10):
    state = ''.join(do_step1(state, rules))
    #print(state)
  counts = collections.defaultdict(lambda: 0)
  for i in state:
    counts[i] += 1
  #sizes = { v: k for k, v in counts.items() }
  return max(counts.values()) - min(counts.values())

def do_step2(state, rules):
  new_state = collections.defaultdict(lambda: 0)
  for (e1, e2), count in state.items():
    try:
      enew = rules[e1,e2]
      new_state[e1,enew] += count
      new_state[enew,e2] += count
    except KeyError:
      new_state[e1,e2] = state[e1,e2]
  return new_state

def part2(data):
  initial, rules = data
  paircounts = collections.defaultdict(lambda: 0)
  for e1, e2 in zip(initial, initial[1:]):
    paircounts[e1,e2] += 1

  for i in range(40):
    paircounts = do_step2(paircounts, rules)

  counts = collections.defaultdict(lambda: 0)
  for (e1, e2), count in paircounts.items():
    counts[e1] += count
    counts[e2] += count
  # everything is double counted except the start and end of the initial state
  counts[initial[0]] += 1
  counts[initial[-1]] += 1
  counts = { k: v//2 for k, v in counts.items() }
  #print(dict(counts))
  return max(counts.values()) - min(counts.values())

def read_data(name):
  rules = dict()

  with open(name, 'r') as f:
    initial = f.readline().strip()
    f.readline()
    for line in f:
      pair, ele = line.strip().split(' -> ')
      rules[tuple(pair)] = ele
  return initial, rules

def main():
   test_data = read_data('14-test.txt')
   data = read_data('14.txt')

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
