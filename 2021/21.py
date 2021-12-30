#!/usr/bin/env python3

import collections
import itertools
import io

class GameWon(Exception): pass

class PlayerState:
  def __init__(self, pos, score=0):
    self.pos = pos
    self.score = score
  
  def take_turn(self, die):
    _, roll1 = next(die)
    _, roll2 = next(die)
    times_rolled, roll3 = next(die)

    self.pos += roll1 + roll2 + roll3
    while self.pos > 10:
      self.pos -= 10
    
    self.score += self.pos
    if self.score >= 1000:
      raise GameWon(times_rolled)

def part1(data):
  p1, p2 = data
  p1 = PlayerState(p1)
  p2 = PlayerState(p2)

  die = enumerate(itertools.cycle(range(1, 101)), start=1)

  while True:
    try:
      p1.take_turn(die)
    except GameWon as e:
      times_rolled, = e.args
      return p2.score * times_rolled

    try:
      p2.take_turn(die)
    except GameWon as e:
      times_rolled, = e.args
      return p1.score * times_rolled

GameState = collections.namedtuple('GameState', ['p1', 's1', 'p2', 's2'])

def do_p1_turn(states):
  new_states = collections.defaultdict(lambda: 0)
  for state, num_universes in states.items():
    if state.s1 >= 21 or state.s2 >= 21:
      # game is over
      new_states[state] += num_universes
      continue
    for roll in itertools.product((1,2,3), repeat=3):
      new_pos = (state.p1 + sum(roll) - 1) % 10 + 1
      new_state = GameState(new_pos, state.s1 + new_pos, state.p2, state.s2)
      new_states[new_state] += num_universes
  return new_states

def do_p2_turn(states):
  new_states = collections.defaultdict(lambda: 0)
  for state, num_universes in states.items():
    if state.s1 >= 21 or state.s2 >= 21:
      # game is over
      new_states[state] += num_universes
    else:
      # game is not over
      for roll in itertools.product((1,2,3), repeat=3):
        new_pos = (state.p2 + sum(roll) - 1) % 10 + 1
        new_state = GameState(state.p1, state.s1, new_pos, state.s2 + new_pos)
        new_states[new_state] += num_universes
  return new_states

def part2(data):
  p1, p2 = data
  states = collections.defaultdict(lambda: 0)
  states[GameState(p1, 0, p2, 0)] = 1

  while any(map(lambda x: x.s1 < 21 and x.s2 < 21, states)):
    states = do_p1_turn(states)
    states = do_p2_turn(states)

  v1 = sum(itertools.compress(states.values(), map(lambda s: s.s1 >= 21, states)))
  v2 = sum(itertools.compress(states.values(), map(lambda s: s.s2 >= 21, states)))
  return max(v1, v2)

def read_data(name_or_data):
  if isinstance(name_or_data, str):
     buf = open(name_or_data, 'r')
  elif isinstance(name_or_data, bytes):
     buf = io.BytesIO(name_or_data)
  else:
     raise ValueError(name_or_data)

  p1 = int(buf.readline().strip().split(': ')[1])
  p2 = int(buf.readline().strip().split(': ')[1])
  return p1, p2

def main():
   test_data = read_data('21-test.txt')
   data = read_data('21.txt')

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
