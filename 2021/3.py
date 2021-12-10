#!/usr/bin/env python3

import itertools

def part1(data):
   bit_count = [0] * len(data[0])
   for datum in data:
      for pos, bit in enumerate(datum):
         if bit == '1':
            bit_count[pos] += 1
   
   gamma   = ''.join(map(lambda x: '1' if x > len(data)/2 else '0', bit_count))
   epsilon = ''.join(map(lambda x: '1' if x < len(data)/2 else '0', bit_count))

   return int(gamma, 2) * int(epsilon, 2)

def part2(data):
   def split(data, position, bias_most):
      ones = []
      zeroes = []
      for datum in data:
         if datum[position] == '0':
            zeroes.append(datum)
         else:
            ones.append(datum)
      #print(f'split: {zeroes}, {ones}, {position}, {bias_most}')
      if bias_most:
         return ones if len(ones) >= len(zeroes) else zeroes
      else:
         return zeroes if len(zeroes) <= len(ones) else ones   
   
   data_most = data.copy()
   data_least = data.copy()

   bit = itertools.cycle(range(len(data[0])))
   while len(data_most) > 1:
      data_most = split(data_most, next(bit), True)
   
   bit = itertools.cycle(range(len(data[0])))
   while len(data_least) > 1:
      data_least = split(data_least, next(bit), False)

   return int(data_most[0], 2) * int(data_least[0], 2)
   


test_data = list(x.strip() for x in open('3-test.txt').readlines())
print(f'part1 test: {part1(test_data)}')
print(f'part2 test: {part2(test_data)}')

data = list(x.strip() for x in open('3.txt').readlines())

print(f'part1 real: {part1(data)}')
print(f'part2 real: {part2(data)}')

  
