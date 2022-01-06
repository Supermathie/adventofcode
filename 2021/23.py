#!/usr/bin/env python3

import collections
import itertools
import io
import re
import textwrap
import queue

class State:

  MOVE_COSTS = {
    'A': 1,
    'B': 10,
    'C': 100,
    'D': 1000,
  }

  # room exit locations
  A_pos = 2
  B_pos = 4
  C_pos = 6
  D_pos = 8
  room_pos = {A_pos, B_pos, C_pos, D_pos}

  def __init__(self, hallway, A, B, C, D):
    self.hallway = tuple(hallway)
    self.A = tuple(A)
    self.B = tuple(B)
    self.C = tuple(C)
    self.D = tuple(D)

  def __repr__(self):
    return f"State({self.hallway}, {self.A}, {self.B}, {self.C}, {self.D})"

  def __eq__(self, other):
    return all((
      self.hallway == other.hallway,
      self.A == other.A,
      self.B == other.B,
      self.C == other.C,
      self.D == other.D,
    ))
  
  def __lt__(self, other):
    return hash(self) < hash(other)
  
  def __hash__(self):
    return hash(f'{self.hallway}{self.A}{self.B}{self.C}{self.D}')

  def copy(self, **kwargs):
    other = State(self.hallway, self.A, self.B, self.C, self.D)
    for k,v in kwargs.items():
      setattr(other, k, tuple(v))
    return other

  def solved(self):
    return all((
      all(map(lambda c: c == 'A', self.A)),
      all(map(lambda c: c == 'B', self.B)),
      all(map(lambda c: c == 'C', self.C)),
      all(map(lambda c: c == 'D', self.D)),
    ))

  def __str__(self):
    s    = f'#############\n'
    s   += f"#{''.join(c or '.' for c in self.hallway)}#\n"
    for i, _ in enumerate(self.A):
      s += f"{'#' if i == 0 else ' '}{'#' if i == 0 else ' '}#{self.A[i] or '.'}#{self.B[i] or '.'}#{self.C[i] or '.'}#{self.D[i] or '.'}#{'#' if i == 0 else ''}{'#' if i == 0 else ''}\n"
    s   += f'  #########\n'
    return s

  def room(self, room):
    return getattr(self, room)

  def room_enterable(self, room):
    '''is room enterable? returns False or the number of steps to get into position
    
    >>> State( ('A', *(None,) * 10), [None, 'A'], ['B']*2, ['C']*2, ['D']*2 ).room_enterable('A')
    1

    >>> State( ('A', 'A', *(None,) * 9), [None, None], ['B']*2, ['C']*2, ['D']*2 ).room_enterable('A')
    2

    >>> State( ('A', *(None,) * 10), [None, 'A', 'A', 'A'], ['B']*4, ['C']*4, ['D']*4 ).room_enterable('A')
    1

    >>> State( ('B', *(None,) * 10), [None, 'A', 'A', 'A'], 'ABBB', ['C']*4, ['D']*4 ).room_enterable('B')
    False

    >>> State( ('A', *(None,) * 10), [None, 'A', 'B', 'A'], ['B']*4, ['C']*4, ['D']*4 ).room_enterable('A')
    False
    '''
    contents = self.room(room)
    if not all(map(lambda x: x == room or x is None, self.room(room))):
      return False
    
    try:
      return contents.index(room)
    except ValueError:
      return len(contents)

  def valid_moves(self):
    '''iters valid moves
    
    >>> next(State( ('A', *(None,) * 10), [None, 'A'], ['B']*2, ['C']*2, ['D']*2 ).valid_moves())
    (3, State((None, None, None, None, None, None, None, None, None, None, None), ('A', 'A'), ('B', 'B'), ('C', 'C'), ('D', 'D')))
    '''
    for i, c in enumerate(self.hallway):
      if c:
        slot = self.room_enterable(c)
        if slot:
          room_pos = getattr(self, f'{c}_pos')
          if ( i < room_pos and not any(self.hallway[i+1:room_pos]) ) or \
            ( i > room_pos and not any(self.hallway[room_pos:i]) ):
              num_steps = abs(room_pos - i) # move to the room entrance
              num_steps += slot             # move into the room
              new_hallway = ( *self.hallway[:i], None, *self.hallway[i+1:] )
              new_room = list(getattr(self, c))
              new_room[slot-1] = c
              yield num_steps * self.MOVE_COSTS[c], self.copy(**{'hallway': new_hallway, c: new_room})
    
    # leave from rooms?
    for room in 'A', 'B', 'C', 'D':
      contents = self.room(room)
      if all(map(lambda x: x == room or x is None, contents)):
        # desired end state, don't change
        continue
      for slot in range(len(contents)):
        if contents[slot] is not None:
          c = contents[slot]
          new_room = list(contents)
          new_room[slot] = None
          room_pos = getattr(self, f'{room}_pos')
          for i in set(range(0, room_pos)) - self.room_pos:
            if not any(self.hallway[i:room_pos]):
              num_steps = room_pos - i + slot + 1
              new_hallway = ( *self.hallway[:i], c, *self.hallway[i+1:] )
              yield num_steps * self.MOVE_COSTS[c], self.copy(**{'hallway': new_hallway, room: new_room})

          for i in set(range(room_pos + 1, len(self.hallway))) - self.room_pos:
            if not any(self.hallway[room_pos:i+1]):
              num_steps = i - room_pos + slot + 1
              new_hallway = ( *self.hallway[:i], c, *self.hallway[i+1:] )
              yield num_steps * self.MOVE_COSTS[c], self.copy(**{'hallway': new_hallway, room: new_room})
          break # once we find an occupied space, stop looking


def dijkstra(start_state, end_check, neighbours):
  visited = {}

  candidates = queue.PriorityQueue()
  candidates.put((0, start_state, None))

  while not candidates.empty():
    dist, cur, prev = candidates.get_nowait()
    #print('Evaluating distance {dist}:')
    #print(textwrap.indent(str(cur), '  '))
    if cur in visited:
      continue
    visited[cur] = dist, prev

    if end_check(cur):
      steps = []
      i = cur
      while i is not None:
        steps.append((visited[i][0], i))
        i = visited[i][1]
      return tuple(reversed(steps))

    for n_dist, n in neighbours(cur):
      if n not in visited:
        candidates.put((dist + n_dist, n, cur))
  raise RuntimeError('end state not found!')

def part1(data):
  initial = data
  steps = dijkstra(initial, State.solved, State.valid_moves)
  if False:
    for cost, state in steps:
      print(f'Distance: {cost}')
      print(textwrap.indent(str(state), '  '))

  return steps[-1][0]

def part2(data):
  initial = State(
    data.hallway,
    (data.A[0], 'D', 'D', data.A[1]),
    (data.B[0], 'C', 'B', data.B[1]),
    (data.C[0], 'B', 'A', data.C[1]),
    (data.D[0], 'A', 'C', data.D[1]),
  )
  steps = dijkstra(initial, State.solved, State.valid_moves)
  if False:
    for cost, state in steps:
      print(f'Distance: {cost}')
      print(textwrap.indent(str(state), '  '))

  return steps[-1][0]

def read_data(name_or_data):
  '''Reads the data, as you might expect

  >>> read_data(b'#############\\n#...........#\\n###B#C#B#D###\\n  #A#D#C#A#\\n  #########  \\n')
  State((None, None, None, None, None, None, None, None, None, None, None), ('B', 'A'), ('C', 'D'), ('B', 'C'), ('D', 'A'))
  '''
  if isinstance(name_or_data, str):
     buf = open(name_or_data, 'r')
  elif isinstance(name_or_data, bytes):
     buf = io.TextIOWrapper(io.BytesIO(name_or_data))
  else:
     raise ValueError(name_or_data)

  if next(buf).strip() != '#' * 13:
    raise ValueError("unexpected data")
  if next(buf).strip() != '#' + '.' * 11 + '#':
    raise ValueError("unexpected data")

  line = next(buf).strip()
  m = re.match('^###([ABCD])#([ABCD])#([ABCD])#([ABCD])###$', line)
  if m is None:
    print(f'{line} did not parse')
  g = m.groups()
  room_A = [g[0], None]
  room_B = [g[1], None]
  room_C = [g[2], None]
  room_D = [g[3], None]

  line = next(buf).rstrip()
  m = re.match('^  #([ABCD])#([ABCD])#([ABCD])#([ABCD])#$', line)
  if m is None:
    print(f'{line} did not parse')
  g = m.groups()
  room_A[1] = g[0]
  room_B[1] = g[1]
  room_C[1] = g[2]
  room_D[1] = g[3]
  return State([None for _ in range(11)], room_A, room_B, room_C, room_D)

def main():
   test_data = read_data('23-test.txt')
   data = read_data('23.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part2 test: {part2(test_data)}')

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
