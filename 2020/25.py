#!/usr/bin/env python3

import sys

g = 7
p = 20201227

card_pub, door_pub = (int(x.strip()) for x in open(sys.argv[1]))

card_priv = 0
card_cur = 1

while card_cur != card_pub:
  card_cur = (card_cur * g) % p
  card_priv += 1

enc_key = 1

for i in range(card_priv):
  enc_key = (enc_key * door_pub) % p

print("Part 1:", enc_key)
