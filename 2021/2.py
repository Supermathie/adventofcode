#!/usr/bin/env python3

import itertools

data = open('2.txt').readlines()

def part1(data):
   pos = 0
   depth = 0


   for cmd in data:
     d, l = cmd.split()
     x = int(l)

     if d == 'forward':
        pos = pos + x
     elif d == 'down':
        depth = depth + x
     elif d == 'up':
        depth = depth - x
     else:
       raise "uhoh"
   return pos * depth

def part2(data):
   pos = 0
   depth = 0
   aim = 0


   for cmd in data:
     d, l = cmd.split()
     x = int(l)

     if d == 'forward':
        pos = pos + x
        depth = depth + aim*x
     elif d == 'down':
        aim = aim + x
     elif d == 'up':
        aim = aim - x
     else:
       raise "uhoh"
   return pos * depth

print(part1(data))
print(part2(data))

  
