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

def decode_packet(bitstream):
  '''Decodes a packet as per specifications on https://adventofcode.com/2021/day/16

  Return: packet_version, packet_type, packet_value

  >>> decode_packet(make_bitstream(io.BytesIO(b'D2FE28')))
  (6, 4, 2021)

  >>> decode_packet(make_bitstream(io.BytesIO(b'38006F45291200')))
  (1, 6, ((6, 4, 10), (2, 4, 20)))

  >>> decode_packet(make_bitstream(io.BytesIO(b'EE00D40C823060')))
  (7, 3, ((2, 4, 1), (4, 4, 2), (1, 4, 3)))
  '''
  bitstream = iter(bitstream)

  V = bits_to_uint(islice(bitstream, 3))
  T = bits_to_uint(islice(bitstream, 3))
  if T == 4: # literal value
    more_packets = True
    bits = []
    while more_packets:
      more_packets = next(bitstream) == 1
      bits.extend(islice(bitstream, 4))
    return V, T, bits_to_uint(bits)

  # not a literal
  L = next(bitstream)
  subpackets = []
  if L == 0: # 15 bits -> total length in bits of subpackets
    Llen = bits_to_uint(islice(bitstream, 15))
    subpackets_stream = islice(bitstream, Llen)
    try:
      while True:
        subpackets.append(decode_packet(subpackets_stream))
    except StopIteration:
      pass
  else: # 11 bits -> number of sub-packets
    Lnum = bits_to_uint(islice(bitstream, 11))
    while len(subpackets) < Lnum:
      subpackets.append(decode_packet(bitstream))
  return V, T, tuple(subpackets)

def part1(data):
  packet = decode_packet(data)
  def sum_versions(packet):
    ver, type_, val = packet
    if isinstance(val, tuple):
      return ver + sum(map(sum_versions, val))
    else:
      return ver
  return sum_versions(packet)

def process_packet(packet):
  _, T, val = packet
  if T == 4:
    return val
  return {
    0: sum,
    1: lambda x: functools.reduce(op.mul, x),
    2: min,
    3: max,
    5: lambda x: int(op.gt(*x)),
    6: lambda x: int(op.lt(*x)),
    7: lambda x: int(op.eq(*x)),
  }[T](map(process_packet, val))

def part2(data):
  '''
  Does the thing!

  >>> part2(make_bitstream(io.BytesIO(b'C200B40A82')))
  3

  >>> part2(make_bitstream(io.BytesIO(b'04005AC33890')))
  54

  >>> part2(make_bitstream(io.BytesIO(b'880086C3E88112')))
  7

  >>> part2(make_bitstream(io.BytesIO(b'CE00C43D881120')))
  9

  >>> part2(make_bitstream(io.BytesIO(b'D8005AC2A8F0')))
  1

  >>> part2(make_bitstream(io.BytesIO(b'F600BC2D8F')))
  0

  >>> part2(make_bitstream(io.BytesIO(b'9C005AC2F8F0')))
  0

  >>> part2(make_bitstream(io.BytesIO(b'9C0141080250320F1802104A08')))
  1
  '''
  packet = decode_packet(data)
  return process_packet(packet)

def make_bitstream(buf):
  c = buf.read(1)
  while c not in (b'\n', b''):
    h = int(c, 16)
    yield (h & 0b1000) >> 3
    yield (h & 0b0100) >> 2
    yield (h & 0b0010) >> 1
    yield (h & 0b0001) >> 0
    c = buf.read(1)

def read_data(name):
  return make_bitstream(open(name, 'rb'))

def main():
   test_data = read_data('16-test.txt')
   print(f'part1 test: {part1(test_data)}')
   test_data = read_data('16-test.txt')
   print(f'part2 test: {part2(test_data)}')

   data = read_data('16.txt')
   print(f'part1 real: {part1(data)}')
   data = read_data('16.txt')
   print(f'part2 real: {part2(data)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
