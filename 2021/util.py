#!/usr/bin/env python3

import functools
import io
import operator as op

from itertools import islice

def bits_to_uint(bits):
  '''Treats a big-endian bitstream of arbitrary size as an unsigned integer

  >>> bits_to_uint(())
  0

  >>> bits_to_uint((1,))
  1

  >>> bits_to_uint((0,1))
  1

  >>> bits_to_uint((1,0))
  2

  >>> bits_to_uint((1,0,1,0))
  10

  >>> bits_to_uint((1,1,1,1,0,0,1))
  121
  '''
  val = 0
  for bit in bits:
    val <<= 1
    val |= bit
  return val

if __name__ == '__main__':
  import sys
  if len(sys.argv) > 1 and sys.argv[1] == '--test':
      import doctest
      doctest.testmod()
  elif len(sys.argv) > 1 and sys.argv[1] == '-i':
      from IPython import embed
      embed()
