#!/usr/bin/env python3

import collections
import itertools
from termcolor import colored

dneigh = (
 (-1,-1), (-1,0), (-1, 1),
 ( 0,-1),         ( 0, 1),
 ( 1,-1), ( 1,0), ( 1, 1),
)

dom = None
rng = None

def print_fish(data):
  for y in rng:
    for x in dom:
      e = data[y][x]
      if e == 0:
        print(colored(f'{e}', 'white', attrs=['bold']), end='')
      else:
        print(f'{e}', end='')
    print()
  print('\n')

def do_step(data):
  flashed = set()
  flashing = set()
  for x in dom:
    for y in rng:
      data[y][x] += 1
      if data[y][x] > 9:
        flashing.add((x,y))

  while len(flashing) > 0:
    x, y = flashing.pop()
    if data[y][x] > 9:
      flashed.add((x, y))
      data[y][x] = 0
      for nx, ny in neighbours(data, x, y) - flashed:
        data[ny][nx] += 1
        if data[ny][nx] > 9:
          flashing.add((nx,ny))
  #print_fish(data)
  return flashed
  
def neighbours(data, x, y):
  n = set()
  for dx, dy in dneigh:
    if x+dx in dom and y+dy in rng:
      n.add( (x+dx, y+dy) )
  return n

def part1(data):
  global dom, rng
  dom = range(len(data[0]))
  rng = range(len(data))

  total_flashes = 0
  for i in range(100):
    total_flashes += len(do_step(data))
  return total_flashes

def part2(data):
  global dom, rng
  dom = range(len(data[0]))
  rng = range(len(data))

  for i in itertools.count(1):
    flashers = do_step(data)
    if len(flashers) == len(data) * len(data[0]):
      return i

def read_data(name):
  with open(name, 'r') as f:
    return [ [ int(x) for x in line.strip() ] for line in f ]

def main():
   test_data = read_data('11-test.txt')
   print(f'part1 test: {part1(test_data)}')
   test_data = read_data('11-test.txt')
   print(f'part2 test: {part2(test_data)}')

   data = read_data('11.txt')
   print(f'part1 real: {part1(data)}')
   data = read_data('11.txt')
   print(f'part2 real: {part2(data)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
