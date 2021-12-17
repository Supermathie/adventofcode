#!/usr/bin/env python3

import functools
import io
import operator as op
import math

import itertools

def find_x(dom):
  '''Generates tuples of (inital x velocity, x position, time value, x has stopped moving?)'''
  for vx_0 in range(int(math.sqrt(dom.start)), dom.stop):
    x = 0
    for dx, t in zip(range(vx_0, 0, -1), itertools.count(1)):
      x += dx
      if x in dom:
        yield vx_0, x, t, dx == 1
      elif x > dom.stop:
        break

def x_valid_times(data):
  times = set()
  dx0min = None
  for vx_0, _, t, dx0 in data:
    times.add(t)
    if dx0:
      if dx0min is None or dx0min > t:
        dx0min = t
  return times, dx0min

def find_y(rng, valid_x_times, dx0min):
  for vy_0 in range(-abs(rng.start), max(abs(rng.start), abs(rng.stop))*2):
    y = 0
    max_y = 0
    vy = vy_0
    for t in itertools.count(1):
      y += vy
      max_y = max(max_y, y)
      vy -= 1
      if y in rng and (t in valid_x_times or t >= dx0min):
        yield vy_0, y, t, max_y
      elif y < rng.stop:
        break

def get_valid_values(dom, rng):
  valid_x_values = tuple(find_x(dom))
  valid_x_times, dx0min = x_valid_times(valid_x_values)
  valid_y_values = tuple(find_y(rng, valid_x_times, dx0min))
  return valid_x_values, valid_y_values

def part1(data):
  valid_x, valid_y = get_valid_values(*data)
  return max(valid_y, key=op.itemgetter(3))[3]

def part2(data):
  valid_x, valid_y = get_valid_values(*data)

  valid = set()
  for x, y in itertools.product(valid_x, valid_y):
    if x[2] == y[2] or (y[2] > x[2] and x[3]):
      valid.add((x[0], y[0]))
  return len(valid)

def main():
   # target area: x=20..30, y=-10..-5
   test_data = (range(20, 30+1), range(-10, -5+1))

   # target area: x=175..227, y=-134..-79
   data = (range(175, 227+1), range(-134, -79+1))

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
