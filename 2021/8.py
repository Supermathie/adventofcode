#!/usr/bin/env python3

import collections
import itertools

def part1(data):
  return sum((sum(map(lambda x: len(x) in (2,3,4,7), output)) for _, output in data))

def solve(inputs, outputs):
  '''

  >>> solve('acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab'.split(), 'cdfeb fcadb cdfeb cdbaf'.split())
  5353
  '''
  inputs = [ set(i) for i in inputs ]
  outputs = [ frozenset(o) for o in outputs ]

  digits = {}
  for i in inputs:
    if len(i) == 2:
      digits[1] = i
    elif len(i) == 3:
      digits[7] = i
    elif len(i) == 4:
      digits[4] = i
    elif len(i) == 7:
      digits[8] = i

  inputs.remove(digits[1])
  inputs.remove(digits[4])
  inputs.remove(digits[7])
  inputs.remove(digits[8])
  
  seg_a = digits[7] - digits[1]

  freqs = { key: len(list(group)) for key, group in itertools.groupby(sorted(itertools.chain(*inputs))) }

  seg_e = set(next(filter(lambda k: freqs[k] == 3, freqs))) # 

  digits[9] = digits[8] - seg_e
  inputs.remove(digits[9])

  digits[0] = next(filter(lambda x: len(x) == 6 and digits[1] < x, inputs))
  inputs.remove(digits[0])

  digits[6] = next(filter(lambda x: len(x) == 6, inputs))
  inputs.remove(digits[6])

  digits[3] = next(filter(lambda x: digits[1] < x, inputs))
  inputs.remove(digits[3])

  digits[2] = next(filter(lambda x: seg_e < x, inputs))
  inputs.remove(digits[2])

  digits[5] = inputs[0]
  inputs.remove(digits[5])

  digits_encoded = { frozenset(v): k for k, v in digits.items() }

  return \
    digits_encoded[outputs[0]] * 1000 + \
    digits_encoded[outputs[1]] * 100  + \
    digits_encoded[outputs[2]] * 10   + \
    digits_encoded[outputs[3]] * 1

def part2(data):
  #for input_, output in data:
    #print(f'{" ".join(output)}: {solve(input_, output)}')
  return sum((solve(input_, output) for input_, output in data))

def read_data(name):
  data = []
  with open(name, 'r') as f:
    for line in f:
      inputs, outputs = line.split(' | ')
      data.append((inputs.split(), outputs.split()))
  return data

def main():
   test_data = read_data('8-test.txt')
   data = read_data('8.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part2 test: {part2(test_data)}')

   print(f'part1 real: {part1(data)}')
   print(f'part2 real: {part2(data)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
