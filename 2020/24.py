#!/usr/bin/env python3

import collections
import sys
import re


# moves = { # https://www.redblobgames.com/grids/hexagons/#neighbors-axial
#   'e': [1, 0],
#   'w': [-1, 0],
#   'ne': [1, 1],
#   'se': [0, -1],
#   'nw': [0, 1],
#   'sw': [-1, -1],
# }

moves = {
  'e': [2, 0],
  'w': [-2, 0],
  'ne': [1, 1],
  'se': [1, -1],
  'nw': [-1, 1],
  'sw': [-1, -1],
}

def round(tiles):
  countTiles = collections.defaultdict(lambda: 0)
  for tile in tiles:
    for move in moves.values():
      countTiles[(tile[0] + move[0], tile[1] + move[1])] += 1
  newTiles = {}
  for newTile, v in countTiles.items():
    if newTile in tiles:
      if v > 2:
        pass # should be white
      else:
        newTiles[newTile] = 1
    else:
      if v == 2:
        newTiles[newTile] = 1
      else:
        pass # should be white
  return newTiles

tiles = {}
for line in open(sys.argv[1], 'r'):
  x = 0
  y = 0
  for d in re.findall('(s[ew]|n[ew]|e|w)', line):
    dx, dy = moves[d]
    x += dx
    y += dy
  if (x,y) in tiles:
    del(tiles[(x,y)])
  else:
    tiles[(x, y)] = 1

print("Part 1:", len(tiles))


for _ in range(100):
  tiles = round(tiles)

print("Part 2:", len(tiles))
