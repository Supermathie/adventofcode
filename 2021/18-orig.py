#!/usr/bin/env python3

import functools
import operator as op
import re

from itertools import islice

class SnailNum:
  def __init__(self, dat):
    if isinstance(dat, str):
      self.contents = self._parse(dat)
    elif isinstance(dat, list):
      self.contents = dat
    elif isinstance(dat, tuple):
      self.contents = list(dat)
    elif isinstance(dat, SnailNum):
      self.contents = dat.contents.copy()
    else:
      raise ArgumentError(f'what is a {type(dat)}?')

  def __repr__(self):
    return f'SnailNum({repr(self.contents)})'

  def __add__(self, other):
    '''Adds

    >>> SnailNum([(1,1),(1,1)]) + SnailNum([(2,1),(2,1)])
    SnailNum([(1, 2), (1, 2), (2, 2), (2, 2)])
    '''
    l = [ (x[0], x[1]+1) for x in self.contents ]
    l.extend( (x[0], x[1]+1) for x in other.contents )
    n = SnailNum(l)
    return n.reduce()

  def _check_explode(self):
    for i, elem in enumerate(self.contents):
      if elem[1] == 5: # explode
        if i > 0:
          self.contents[i-1] = (self.contents[i-1][0] + elem[0], self.contents[i-1][1])
        if i+2 < len(self.contents):
          self.contents[i+2] = (self.contents[i+2][0] + self.contents[i+1][0], self.contents[i+2][1])
        del self.contents[i]
        self.contents[i] = (0, elem[1] - 1)
        return True
    return False

  def _check_split(self):
    for i, elem in enumerate(self.contents):
      if elem[0] >= 10: # split
        self.contents[i:i+1] = (elem[0] // 2, elem[1] + 1), ( (elem[0]+1) // 2, elem[1] + 1)
        return True
    return False
    

  def magnitude(self):
    
  def reduce(self):
    '''Reduces

    >>> SnailNum('[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]').reduce()
    SnailNum([(0, 4), (7, 4), (4, 3), (7, 4), (8, 4), (6, 4), (0, 4), (8, 2), (1, 2)])

    >>> SnailNum('[[[[[9,8],1],2],3],4]').reduce()
    SnailNum([(0, 4), (9, 4), (2, 3), (3, 2), (4, 1)])

    >>> SnailNum('[11,0]').reduce()
    SnailNum([(5, 2), (6, 2), (0, 1)])

    >>> SnailNum('[0,10]').reduce()
    SnailNum([(0, 1), (5, 2), (5, 2)])
    '''
    while True:
      if self._check_explode():
        continue
      if self._check_split():
        continue
      break
    return self

  @staticmethod
  def _parse(s):
    l = []
    depth = 0

    regex = '|'.join((
      '(?P<L>\[)',
      '(?P<R>\])',
      '(?P<N>\d+)',
      '(?P<C>,)',
      '(?P<E>.)',
    ))
    for m in re.finditer(regex, s):
      if m.lastgroup == 'L':
        depth += 1
      elif m.lastgroup == 'R':
        depth -= 1
        if depth < 0:
          raise ValueError(f'unmatched "]" at position {m.start()}')
      elif m.lastgroup == 'N':
        l.append( (int(m.group()), depth) )
      elif m.lastgroup == 'C':
        pass
      else:
        raise ValueError(f'unknown "{m.group()}" at position {m.start()}')
    if depth != 0:
      raise ValueError(f'expected "]" at end of string')
    return l
    

def part1(data):
  return functools.reduce(op.add, data)

def part2(data):
  pass

def read_data(name):
  with open(name, 'r') as f:
    return [ SnailNum(line) for line in f ]

def main():
   test_data = read_data('18-test.txt')
   print(f'part1 test: {part1(test_data)}')
   print(f'part2 test: {part2(test_data)}')

   data = read_data('18.txt')
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
