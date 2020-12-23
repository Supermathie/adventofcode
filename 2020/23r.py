#!/usr/bin/env python3

import sys

def game(state, num_cups, moves):
  state.extend(range(len(state)+1, num_cups+1))
  cur = 0
  for i in range(moves):
    cur_cup = state[cur]
    # print(f"{i+1}: cur:{cur} {state}")
    destination = state[cur]-1
    if cur < num_cups-3:
      selected = state[cur+1:cur+4]
      state[cur+1:cur+4] = []
    elif cur == num_cups-3:
      selected = state[num_cups-2:num_cups] + state[0:1]
      state[num_cups-2:num_cups] = []
      state[0:1] = []
    elif cur == num_cups-2:
      selected = state[num_cups-1:num_cups] + state[0:2]
      state[num_cups-1:num_cups] = []
      state[0:2] = []
    elif cur == num_cups-1:
      selected = state[0:3]
      state[0:3] = []
    # print(f"selected:{selected}")
    while destination not in state:
      destination -= 1
      if destination < 1:
        destination = num_cups
    # print(f"dest:{destination}")
    dest_index = state.index(destination)
    state[dest_index+1:dest_index+1] = selected
    cur = state.index(cur_cup)
    cur = (cur + 1) % 9
  return state

state = [3, 8, 9, 1, 2, 5, 4, 6, 7]
finalState = game(state[:], len(state), 100)
output = "".join((str(i) for i in finalState))
print(f"Part 1 test: {output}")

state = [1, 3, 7, 8, 2, 6, 4, 9, 5]
finalState = game(state[:], len(state), 100)
output = "".join((str(i) for i in finalState))
print(f"Part 1: {output}")

state = [1, 3, 7, 8, 2, 6, 4, 9, 5]
finalState = game(state[:], 1_000_000, 10_000_000)
cup1 = finalState.index(1)
output = finalState[cup1+1] * finalState[cup1+2]
print(f"Part 2: {output}")
