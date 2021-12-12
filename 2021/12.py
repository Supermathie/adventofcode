#!/usr/bin/env python3

import collections
import itertools
from termcolor import colored
from timeit import timeit
import datetime

def part1(rooms):
  complete_paths = []
  partial_paths = [ ('start',) ]
 
  while len(partial_paths) > 0:
    cur_path = partial_paths.pop()
    cur_node = cur_path[-1]
    for next_node in rooms[cur_node]:
      if next_node == 'end':
        complete_paths.append(cur_path + (next_node,))
      elif next_node.islower():
        if next_node not in cur_path:
          partial_paths.append(cur_path + (next_node,))
      else:
        partial_paths.append(cur_path + (next_node,))
  return len(complete_paths)

def part2(rooms):
  complete_paths = []
  partial_paths = [ ('start',) ]

  def can_visit_second_small_cave(path):
    #lowers = sorted(filter(str.islower, path))
    #return not any(map(lambda x, y: x == y, lowers, lowers[1:]))
    visited = set()
    for cave in path:
      if cave.islower():
        if cave in visited:
          return False
        visited.add(cave)
    return True

 
  while len(partial_paths) > 0:
    cur_path = partial_paths.pop()
    cur_node = cur_path[-1]
    for next_node in rooms[cur_node]:
      if next_node == 'start':
        continue # start can only be visited once
      if next_node == 'end':
        complete_paths.append(cur_path + (next_node,))
      elif next_node.islower():
        if next_node not in cur_path or can_visit_second_small_cave(cur_path):
          partial_paths.append(cur_path + (next_node,))
      else:
        partial_paths.append(cur_path + (next_node,))
  return len(complete_paths)

def read_data(name):
  rooms = collections.defaultdict(list)

  with open(name, 'r') as f:
    for line in f:
      r1, r2 = line.strip().split('-')
      rooms[r1].append(r2)
      rooms[r2].append(r1)
  return rooms

def main():
   test0_data = read_data('12-test0.txt')
   test1_data = read_data('12-test1.txt')
   test2_data = read_data('12-test2.txt')
   data = read_data('12.txt')

   print(f'part1 test: {part1(test1_data)}')
   print(f'part2 test0: {part2(test0_data)}')
   print(f'part2 test1: {part2(test1_data)}')
   print(f'part2 test2: {part2(test2_data)}')
   
   s1 = datetime.datetime.now()
   p1 = part1(data)
   t1 = datetime.datetime.now() - s1
   s2 = datetime.datetime.now()
   p2 = part2(data)
   t2 = datetime.datetime.now() - s2

   print(f'part1 real: {p1} ({t1.total_seconds()}s)')
   print(f'part2 real: {p2} ({t2.total_seconds()}s)')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
