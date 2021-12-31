#!/usr/bin/env python3

import collections
import itertools
import io
import re

Step = collections.namedtuple('Step', ['on', 'x', 'y', 'z', 'cube'])

def part1(steps):
  pixels = set()
  for step in steps:

    for x, y, z in itertools.product(
      set(step.x) & set(range(-50,51)),
      set(step.y) & set(range(-50,51)),
      set(step.z) & set(range(-50,51)),
    ):
      if -50 <= x <= 50 and -50 <= y <= 50 and -50 <= z <= 50:
        if step.on:
          pixels.add((x, y, z))
        else:
          try:
            pixels.remove((x, y, z))
          except KeyError:
            pass
  return len(pixels)

class Cube:
  def __init__(self, x1, x2, y1, y2, z1, z2, on=True):
    if x1 > x2:
      raise ValueError(f'given x from {x1} to {x2}')
    if y1 > y2:
      raise ValueError(f'given y from {y1} to {y2}')
    if z1 > z2:
      raise ValueError(f'given z from {z1} to {z2}')
    self.x1 = x1
    self.x2 = x2
    self.y1 = y1
    self.y2 = y2
    self.z1 = z1
    self.z2 = z2
    self.on = on

  def __repr__(self):
    return f"Cube({self.x1}, {self.x2}, {self.y1}, {self.y2}, {self.z1}, {self.z2}{'' if self.on else ', on=False'})"

  def __eq__(self, other):
    return all((
      self.x1 == other.x1,
      self.x2 == other.x2,
      self.y1 == other.y1,
      self.y2 == other.y2,
      self.z1 == other.z1,
      self.z2 == other.z2,
    ))
  
  def volume(self):
    return (self.x2-self.x1+1) * (self.y2-self.y1+1) * (self.z2-self.z1+1) * (1 if self.on else -1)

  def overlaps(self, other):
    '''Does self overlap, at all, other?
    
    >>> Cube(0, 0, 0, 0, 0, 0).overlaps(Cube(0, 0, 0, 0, 0, 0))
    True

    >>> Cube(0, 0, 0, 0, 0, 0).overlaps(Cube(-1, 1, -1, 1, -1, 1))
    True

    >>> Cube(-1, 1, -1, 1, -1, 1).overlaps(Cube(0, 0, 0, 0, 0, 0))
    True

    >>> Cube(0, 0, 0, 0, 0, 0).overlaps(Cube(0, 0, 0, 0, 1, 1))
    False
    '''
    return all((
      self.x1 <= other.x1 <= self.x2 or self.x1 <= other.x2 <= self.x2 or other.x1 <= self.x1 <= other.x2,
      self.y1 <= other.y1 <= self.y2 or self.y1 <= other.y2 <= self.y2 or other.y1 <= self.y1 <= other.y2,
      self.z1 <= other.z1 <= self.z2 or self.z1 <= other.z2 <= self.z2 or other.z1 <= self.z1 <= other.z2,
    ))
  
  def mask(self, other):
    '''Returns an iterable of cubes constructed from self that do not overlap other
    
    >>> Cube(0, 0, 0, 0, 0, 0).mask(Cube(0, 0, 0, 0, 0, 0))
    ()

    >>> Cube(0, 0, 0, 0, 0, 0).mask(Cube(-1, 1, -1, 1, -1, 1))
    ()

    >>> Cube(-1, 1, 0, 0, 0, 0).mask(Cube(0, 0, 0, 0, 0, 0))
    (Cube(-1, -1, 0, 0, 0, 0), Cube(1, 1, 0, 0, 0, 0))

    >>> Cube(0, 0, -1, 1, 0, 0).mask(Cube(0, 0, 0, 0, 0, 0))
    (Cube(0, 0, -1, -1, 0, 0), Cube(0, 0, 1, 1, 0, 0))

    >>> Cube(0, 0, 0, 0, -1, 1).mask(Cube(0, 0, 0, 0, 0, 0))
    (Cube(0, 0, 0, 0, -1, -1), Cube(0, 0, 0, 0, 1, 1))

    >>> len(Cube(-1, 1, -1, 1, -1, 1).mask(Cube(0, 0, 0, 0, 0, 0)))
    26

    >>> Cube(0, 0, 0, 0, 0, 0) in Cube(-1, 1, -1, 1, -1, 1).mask(Cube(0, 0, 0, 0, 0, 0))
    False
    '''

    if not self.overlaps(other):
      return self,
    
    new = []
    for i, ((x1, x2), (y1, y2), (z1, z2)) in enumerate(itertools.product(
      ((self.x1, other.x1-1), (other.x1, other.x2), (other.x2+1, self.x2)),
      ((self.y1, other.y1-1), (other.y1, other.y2), (other.y2+1, self.y2)),
      ((self.z1, other.z1-1), (other.z1, other.z2), (other.z2+1, self.z2)),
    )):
      if i != 13: # cube != other, but faster
        try:
          cube = Cube(x1, x2, y1, y2, z1, z2)
        except ValueError:
          pass
        else:
          new.append(cube)
    return tuple(new)

  def intersect(self, other, on=True):
    if not self.overlaps(other):
      raise ValueError("cubes do not intersect")
    return Cube(
      max(self.x1, other.x1), min(self.x2, other.x2),
      max(self.y1, other.y1), min(self.y2, other.y2),
      max(self.z1, other.z1), min(self.z2, other.z2),
      on
    )

def part2(steps):
  cubes = []
  for step in steps:
    new_cubes = []
    for cube in cubes:
      if cube.overlaps(step.cube):
        new_cubes.append(cube.intersect(step.cube, not cube.on))
    if step.cube.on:
      new_cubes.append(step.cube)
    cubes.extend(new_cubes)
  
  return sum(map(Cube.volume, cubes))

def part2_attempt1(steps):
  cubes = []
  for step in steps:
    new_cubes = []
    for cube in cubes:
      new_cubes.extend(cube.mask(step.cube))
    if step.on:
      new_cubes.append(step.cube)
    cubes = new_cubes
  
  return sum(map(Cube.volume, cubes))

def read_data(name_or_data):
  '''Reads the data, as you might expect

  >>> read_data(b'on x=10..12,y=10..12,z=10..12')[0]
  Step(on=True, x=range(10, 13), y=range(10, 13), z=range(10, 13), cube=...)

  >>> read_data(b'off x=-12..-10,y=-12..-10,z=-12..-10')[0]
  Step(on=False, x=range(-12, -9), y=range(-12, -9), z=range(-12, -9), cube=...)

  '''
  if isinstance(name_or_data, str):
     buf = open(name_or_data, 'r')
  elif isinstance(name_or_data, bytes):
     buf = io.TextIOWrapper(io.BytesIO(name_or_data))
  else:
     raise ValueError(name_or_data)

  def parse(line):
    m = re.match('^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$', line.strip())
    if m is None:
      print(f'{line.strip()} did not parse')
    g = m.groups()
    x1, x2, y1, y2, z1, z2 = map(int, g[1:])
    return Step(
      g[0] == 'on',
      range(x1, x2 + 1),
      range(y1, y2 + 1),
      range(z1, z2 + 1),
      Cube(x1, x2, y1, y2, z1, z2, on=g[0] == 'on')
    )
  return tuple(parse(line) for line in buf)

def main():
   test_data = read_data('22-test.txt')
   test2_data = read_data('22-test2.txt')
   test3_data = read_data('22-test3.txt')
   data = read_data('22.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part1 test2: {part1(test2_data)}')
   print(f'part1 test3: {part1(test3_data)}')
   print(f'part2 test3: {part2(test3_data)}')

   print(f'part1 real: {part1(data)}')
   print(f'part2 real: {part2(data)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod(optionflags=doctest.ELLIPSIS)
elif len(sys.argv) > 1 and sys.argv[1] == '-i':
    from IPython import embed
    embed()
else:
    main()
