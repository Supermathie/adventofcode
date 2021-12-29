#!/usr/bin/env python3

import functools
import io
import operator as op
import math
import util

import itertools

def neighbours(x, y, dom = None, rng = None):
  dneigh = (
    (-1,-1), ( 0,-1), ( 1,-1),
    (-1, 0), ( 0, 0), ( 1, 0),
    (-1, 1), ( 0, 1), ( 1, 1),
  )
  for dx, dy in dneigh:
    if dom is not None and x+dx not in dom:
      continue
    if rng is not None and x+dx not in rng:
      continue
    yield (x+dx, y+dy)

def do_iter(algorithm, pixels):
  dom = range(
    min(map(op.itemgetter(0), pixels)) - 1,
    max(map(op.itemgetter(0), pixels)) + 2,
  )
  rng = range(
    min(map(op.itemgetter(1), pixels)) - 1,
    max(map(op.itemgetter(1), pixels)) + 2,
  )

  new = set()
  for x in dom:
    for y in rng:
      val = util.bits_to_uint((1 if n in pixels else 0 for n in neighbours(x, y)))
      if algorithm[val] == 1:
        new.add((x,y))

  return new

def print_image(pixels):
  dom = range(
    min(map(op.itemgetter(0), pixels)),
    max(map(op.itemgetter(0), pixels)) + 1,
  )
  rng = range(
    min(map(op.itemgetter(1), pixels)),
    max(map(op.itemgetter(1), pixels)) + 1,
  )
  for y in rng:
    for x in dom:
      print('#' if (x,y) in pixels else '.', end='')
    print('')
  print('')

def part1(data):
  algorithm, pixels = data
  #print_image(pixels)
  pixels = do_iter(algorithm, pixels)
  #print_image(pixels)
  pixels = do_iter(algorithm, pixels)
  #print_image(pixels)
  return len(pixels)

def part2(data):
  algorithm, pixels = data


def read_data(name_or_data):
  if isinstance(name_or_data, str):
     buf = open(name_or_data, 'r')
  elif isinstance(name_or_data, bytes):
     buf = io.BytesIO(name_or_data)
  else:
     raise ArgumentError

  algorithm = [ {'.': 0, '#': 1}[c] for c in buf.readline().strip() ]
  if buf.readline() != '\n': # blank line
    raise ValueError("expected blank line")

  pixels = set()
  for y, line in enumerate(buf):
    for x, c in enumerate(line.strip()):
      if c == '#':
        pixels.add((x,y))
  
  return algorithm, pixels

def main():
   test_data = read_data('20-test.txt')
   data = read_data('20.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part2 test: {part2(test_data)}')

   print(f'part1 real: {part1(data)}')
   print(f'part2 real: {part2(data)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
elif len(sys.argv) > 1 and sys.argv[1] == '-i':
    from IPython import embed
    embed()
else:
    main()
