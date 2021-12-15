#!/usr/bin/env python3

import collections
import itertools
import queue

def dijkstra(data, start = None, end = None):
  dom = range(len(data[0]))
  rng = range(len(data))
  start = start or (0,0)
  end = end or (rng.stop-1, dom.stop-1)

  distances = [ [ None for _ in dom ] for _ in rng ]
  visited = [ [ False for _ in dom ] for _ in rng ]
  candidates = set(start,)

  def neighbours(x, y):
    dneigh = (
              ( 0,-1),
     (-1, 0),         ( 1, 0),
              ( 0, 1),
    )
    for dx, dy in dneigh:
      if x+dx in dom and y+dy in rng:
         yield (x+dx, y+dy)
  
  distances[0][0] = 0
  visited[0][0] = True
  candidates = {(0, 0)}

  while not visited[rng.stop-1][dom.stop-1]:
    cur = min(candidates, key = lambda x: distances[x[1]][x[0]])
    candidates.remove(cur)
    visited[cur[1]][cur[0]] = True
    for n in neighbours(*cur):
      if visited[n[1]][n[0]]:
        continue
      candidates.add(n)
      new_dist = distances[cur[1]][cur[0]] + data[n[1]][n[0]] 
      if distances[n[1]][n[0]] is None or new_dist < distances[n[1]][n[0]]:
        distances[n[1]][n[0]] = new_dist
  return distances[end[1]][end[0]]

def part1(data):
  return dijkstra(data)

def part2(data):
  mapdata = [ [ (data[y % len(data)][x % len(data[0])] + x // len(data[0]) + y // len(data) - 1) % 9 + 1 for x in range(len(data[0])*5) ] for y in range(len(data)*5) ]
  return dijkstra(mapdata)

def read_data(name):
  with open(name, 'r') as f:
    return [ [ int(x) for x in line.strip() ] for line in f ]

def main():
   test_data = read_data('15-test.txt')
   data = read_data('15.txt')

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
