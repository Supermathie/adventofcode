#!/usr/bin/env python3

import io
import itertools
import operator as op
import util

BORDER_SIZE = 2

def neighbours(x, y, dom = None, rng = None):
  dneigh = (
    (-1,-1), ( 0,-1), ( 1,-1),
    (-1, 0), ( 0, 0), ( 1, 0),
    (-1, 1), ( 0, 1), ( 1, 1),
  )
  for dx, dy in dneigh:
    if dom is not None and x+dx not in dom:
      continue
    if rng is not None and y+dy not in rng:
      continue
    yield (x+dx, y+dy)

def do_iter(algorithm, pixels):
  dom = range(
    min(map(op.itemgetter(0), pixels)) - BORDER_SIZE*2,
    max(map(op.itemgetter(0), pixels)) + BORDER_SIZE*2 + 1,
  )
  rng = range(
    min(map(op.itemgetter(1), pixels)) - BORDER_SIZE*2,
    max(map(op.itemgetter(1), pixels)) + BORDER_SIZE*2 + 1,
  )

  new = set()
  for x, y in itertools.product(dom, rng):
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

def do_iter_2(algorithm, pixels):
  dom = range(
    min(map(op.itemgetter(0), pixels)) - BORDER_SIZE,
    max(map(op.itemgetter(0), pixels)) + BORDER_SIZE + 1,
  )
  rng = range(
    min(map(op.itemgetter(1), pixels)) - BORDER_SIZE,
    max(map(op.itemgetter(1), pixels)) + BORDER_SIZE + 1,
  )
  pixels = do_iter(algorithm, pixels)
  #print_image(pixels)
  pixels = do_iter(algorithm, pixels)
  #print_image(pixels)
  
  if algorithm[0] == 1:
    # we'll have a border around the outside of the image we need to remove
    pixels = pixels.intersection(set(itertools.product(dom, rng)))

  return pixels

def part1(data):
  algorithm, pixels = data
  return len(do_iter_2(algorithm, pixels))

def part2(data):
  algorithm, pixels = data
  for _ in range(25):
    pixels = do_iter_2(algorithm, pixels)
  return len(pixels)

def read_data(name_or_data):
  if isinstance(name_or_data, str):
     buf = open(name_or_data, 'r')
  elif isinstance(name_or_data, bytes):
     buf = io.BytesIO(name_or_data)
  else:
     raise ValueError(name_or_data)

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
   test2_data = read_data('20-test2.txt')
   data = read_data('20.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part1 test2: {part1(test2_data)}')
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
