#!/usr/bin/env python3

import itertools

# Part 1
depths = (int(x) for x in open('1.txt'))
a1, a2 = itertools.tee(depths)
next(a2)
increases = 0
for depth, nextdepth in zip(a1, a2):
   if nextdepth > depth:
      increases += 1

print(increases)

# Part 2
depths = (int(x) for x in open('1.txt'))
a1, a2, a3 = itertools.tee(depths, 3)
next(a2)
next(a3); next(a3)

window1, window2 = itertools.tee((sum(a) for a in zip(a1,a2,a3)))
next(window2)

increases = 0
for depth, nextdepth in zip(window1, window2):
   if nextdepth > depth:
      increases += 1

print(increases)

