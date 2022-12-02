#!/usr/bin/env python3

import functools
import operator as op
import re
import sys

from itertools import islice

class SnailNum:
  class Exploded(Exception): pass
  class ExplodedL(Exception): pass
  class ExplodedR(Exception): pass
  class Split(Exception): pass

  def __init__(self, *args, parent=None):
    self.parent = parent
    if len(args) == 1:
      dat = args[0]
      if isinstance(dat, list):
        self.l, self.r = dat
      elif isinstance(dat, tuple):
        self.l, self.r = dat
      elif isinstance(dat, SnailNum):
        self.l, self.r = dat.copy()
      else:
        raise ValueError(f'what is a {type(dat)}?')
    elif len(args) == 2:
      self.l, self.r = args
    else:
      raise ValueError()

  def __repr__(self):
    return f'SnailNum({self.l},{self.r})'

  def __str__(self):
    return f'[{self.l},{self.r}]'

  def __add__(self, other):
    '''Adds

    >>> SnailNum(1,1) + SnailNum(2,2)
    SnailNum([1,1],[2,2])

    >>> str(SnailNum.parse('[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]') + SnailNum.parse('[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]'))
    '[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]'
    '''
    s = SnailNum(self, other)
    self.parent = s
    other.parent = s
    return s


  def _add_to_left(self, val):
    if isinstance(self.l, int):
      self.l += val
    else:
      self.l._add_to_left(val)

  def _add_to_right(self, val):
    if isinstance(self.r, int):
      self.r += val
    else:
      self.r._add_to_right(val)
    
  def _check_explode(self, depth):
    #print(f'depth:{depth}, l:{self.l}, r:{self.r}', file=sys.stderr)
    if depth > 3:
      # handle l
      prev = self
      cur = self.parent
      while cur is not None and prev is cur.l:
        prev = cur
        cur = cur.parent
      if cur is not None:
        if isinstance(cur.l, int):
          cur.l += self.l
        else:
          cur._add_to_right(self.l)

      # handle r
      prev = self
      cur = self.parent
      while cur is not None and prev is cur.r:
        prev = cur
        cur = cur.parent
      if cur is not None:
        if isinstance(cur.r, int):
          cur.r += self.r
        else:
          cur._add_to_left(self.r)
      
      # replace self with 0
      if self.parent.l is self:
        self.parent.l = 0
      if self.parent.r is self:
        self.parent.r = 0

      raise self.Exploded()
    if isinstance(self.l, SnailNum):
      self.l._check_explode(depth+1)
    if isinstance(self.r, SnailNum):
      self.r._check_explode(depth+1)

  def _check_split(self):
    if isinstance(self.l, int):
      if self.l >= 10:
        self.l = SnailNum(self.l // 2, (self.l+1) // 2, parent=self)
        raise self.Split()
    else:
      self.l._check_split()

    if isinstance(self.r, int):
      if self.r >= 10:
        self.r = SnailNum(self.r // 2, (self.r+1) // 2, parent=self)
        raise self.Split()
    else:
      self.r._check_split()

  def magnitude(self):
    pass

    
  def reduce(self):
    '''Reduces

    >>> SnailNum.parse('[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]').reduce()
    SnailNum([[[0,7],4],[[7,8],[6,0]]],[8,1])

    >>> SnailNum.parse('[[[[[9,8],1],2],3],4]').reduce()
    SnailNum([[[0,9],2],3],4)

    >>> SnailNum.parse('[[0,11],0]').reduce()
    SnailNum([0,[5,6]],0)

    >>> SnailNum(11,0).reduce()
    SnailNum([5,6],0)

    >>> SnailNum(0,10).reduce()
    SnailNum(0,[5,5])
    '''
    print(f'reducing: {self}', file=sys.stderr)
    while True:
      try:
        if self._check_explode(0):
          continue
        if self._check_split():
          continue
      except (self.Exploded, self.ExplodedL, self.ExplodedR, self.Split) as e:
        print(f'restarting after {type(e).__name__}: {self}', file=sys.stderr)
        pass
      else:
        break
    return self

  @staticmethod
  def _parse_gen(s):
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
      yield m

  @staticmethod
  def _parse_con(g):

    self = SnailNum(0, 0)
    sym = next(g)
    if sym.lastgroup == 'L':
      l = SnailNum._parse_con(g)
      l.parent = self
    elif sym.lastgroup == 'N':
      l = int(sym.group())
    else:
      raise ValueError(f'unexpected "{sym.group()}" at position {sym.start()}')

    sym = next(g)
    if sym.lastgroup != 'C':
      raise ValueError(f'unexpected "{sym.group()}" at position {sym.start()}')

    sym = next(g)
    if sym.lastgroup == 'L':
      r = SnailNum._parse_con(g)
      r.parent = self
    elif sym.lastgroup == 'N':
      r = int(sym.group())
    else:
      raise ValueError(f'unexpected "{sym.group()}" at position {sym.start()}')

    sym = next(g)
    if sym.lastgroup != 'R':
      raise ValueError(f'unexpected "{sym.group()}" at position {sym.start()}')

    self.l, self.r = l,r
    return self

  @staticmethod
  def parse(s):
    '''Parse!

    >>> SnailNum.parse('[1,2]')
    SnailNum(1,2)

    >>> SnailNum.parse('[[1,2],3]')
    SnailNum([1,2],3)
    '''
    gen = SnailNum._parse_gen(s)
    sym = next(gen)
    if sym.lastgroup != 'L':
      raise ValueError(f'unexpected "{sym.group()}" at position {sym.start()}')
    n = SnailNum._parse_con(gen)
    return n

  def magnitude(self):
    '''Magnitude!

    >>> SnailNum(1,2).magnitude()
    7

    >>> SnailNum.parse('[[9,1],[1,9]]').magnitude()
    129

    >>> SnailNum.parse('[[1,2],[[3,4],5]]').magnitude()
    143

    >>> SnailNum.parse('[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]').magnitude()
    3488
    '''
    if isinstance(self.l, SnailNum):
      l = self.l.magnitude()
    else:
      l = self.l

    if isinstance(self.r, SnailNum):
      r = self.r.magnitude()
    else:
      r = self.r

    return 3*l + 2*r

def part1(data):
  n = functools.reduce(op.add, data)
  return n.magnitude()

def part2(data):
  pass

def read_data(name):
  with open(name, 'r') as f:
    return [ SnailNum.parse(line) for line in f ]

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
