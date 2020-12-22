#!/usr/bin/env python3

import sys

def game(p1, p2, recurse=False):
  # print(f"start game between {p1} and {p2}")
  seen = set()
  while len(p1) > 0 and len(p2) > 0:
    h = hash((tuple(p1), tuple(p2)))
    if h in seen:
      # print(f"game over (repetition): {p1} {p2}")
      return 1, p1
    seen.add(h)

    # print(f"round state: {p1} {p2}")
    c1, c2 = p1.pop(), p2.pop()
    if recurse and len(p1) >= c1 and len(p2) >= c2:
      winner, _ = game(p1[-c1:], p2[-c2:])
    elif c1 > c2:
      winner = 1
    else:
      winner = 2
    if winner == 1:
      p1.insert(0, c1)
      p1.insert(0, c2)
    else:
      p2.insert(0, c2)
      p2.insert(0, c1)

  # print(f"game over: {p1} {p2}")
  if len(p1) > 0:
    return 1, p1
  else:
    return 2, p2


p1 = []
p2 = []

with open(sys.argv[1], 'r') as f:
  f.readline()
  for line in f:
    if line.strip() == "":
      break
    p1.insert(0, int(line.strip()))
  
  f.readline()
  for line in f:
    p2.insert(0, int(line.strip()))

winner, deck = game(p1[:], p2[:])
winner_sum = sum(((i+1)*x for i, x in enumerate(deck)))
print(f"Part 1: {winner_sum}")

winner, deck = game(p1[:], p2[:], recurse=True)
winner_sum = sum(((i+1)*x for i, x in enumerate(deck)))
print(f"Part 2: {winner_sum}")
