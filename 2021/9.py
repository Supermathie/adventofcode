#!/usr/bin/env python3

import collections
import itertools
from termcolor import colored

def neighbour_vals(data, x, y):
  n = []
  for dx, dy in ( (-1,0), (1,0), (0,-1), (0,1) ):
    try:
      if y+dy < 0 or x+dx < 0:
        continue
      n.append( data[y + dy][x + dx] )
    except IndexError:
      pass
  return n

def part1(data):
  low_points = []
  for y in range(len(data)):
    for x in range(len(data[0])):
      h = data[y][x]
      n = neighbour_vals(data, x, y)
      #print(f'{x},{y} {h} {n}')
      if all(map(lambda i: h < i, n)):
        low_points.append(h)
        print(colored(f'{h}', 'green', attrs=['bold']), end='')
      else:
        pass
        print(f'{h}', end='')
    print()
  #print(low_points)
  return sum(low_points) + len(low_points)

def neighbours(data, x, y):
  n = set()
  for dx, dy in ( (-1,0), (1,0), (0,-1), (0,1) ):
    try:
      if y+dy < 0 or x+dx < 0 or y+dy > len(data)-1 or x+dx > len(data[0])-1:
        continue
      n.add( (x+dx, y+dy) )
    except IndexError:
      pass
  return n

def make_basin(data, x, y):
  nodes_to_process = {(x,y)}
  nodes = set()
  while len(nodes_to_process) > 0:
    node = nodes_to_process.pop()
    if node in nodes or data[node[1]][node[0]] == 9:
      continue
    nodes.add(node)
    nodes_to_process |= neighbours(data, *node)
  return nodes

def part2(data):
  low_points = []
  basins = {}
  for y in range(len(data)):
    for x in range(len(data[0])):
      h = data[y][x]
      n = neighbours(data, x, y)
      if all(map(lambda i: h < data[i[1]][i[0]], n)):
        low_points.append((x,y))
  for low_point in low_points:
    basins[low_point] = make_basin(data, *low_point)
  sizes = sorted(map(lambda s: len(s), basins.values()))
  return sizes[-1] * sizes[-2] * sizes[-3]

def read_data(name):
  with open(name, 'r') as f:
    return [ [ int(x) for x in line.strip() ] for line in f ]

def main():
   test_data = read_data('9-test.txt')
   test2_data = read_data('9-test2.txt')
   data = read_data('9.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part1 test: {part1(test2_data)}')
   print(f'part2 test: {part2(test_data)}')

   print(f'part1 real: {part1(data)}')
   print(f'part2 real: {part2(data)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
