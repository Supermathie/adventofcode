#!/usr/bin/env python3

import collections
import itertools
import operator
import os
import sys

def make_image(dots):
  from PIL import Image
  width = max(map(operator.itemgetter(0), dots))+1
  height = max(map(operator.itemgetter(1), dots))+1
  dom = range(width)
  rng = range(height)
  image = Image.new('L', (width+2, height+2), 'black')
  pixels = image.load()
  for y in rng:
    for x in dom:
      if (x,y) in dots:
        pixels[x+1,y+1] = 0xffffff
  return image

def do_ocr(dots):
  if 'GOOGLE_APPLICATION_CREDENTIALS' not in os.environ:
    raise NotImplemented
  from google.cloud import vision
  import io
  image = make_image(dots)
  b = io.BytesIO()
  image.save(b, format='png')
  #with open('13.png', 'wb') as f:
  #  f.write(b.getvalue())
  client = vision.ImageAnnotatorClient()
  response = client.text_detection(image=vision.Image(content=b.getvalue()))
  return response.text_annotations[0].description.strip()

def print_dots(dots):
  dom = range(max(map(operator.itemgetter(0), dots))+1)
  rng = range(max(map(operator.itemgetter(1), dots))+1)
  for y in rng:
    for x in dom:
      if (x,y) in dots:
        print('#', end='')
      else:
        print(' ', end='')
    print()
  print()

def do_fold(dots, axis, pos):
  above = set()
  below = set()
  axis = 0 if axis == 'x' else 1
  for dot in dots:
    if dot[axis] < pos:
      above.add(dot)
    else:
      below.add(dot)
  if axis == 0:
    return above | { (2*pos-x, y) for x,y in below }
  else:
    return above | { (x, 2*pos-y) for x,y in below }
  

def part1(data):
  dots, folds = data
  return len(do_fold(dots, *folds[0]))

def part2(data):
  dots, folds = data
  for fold in folds:
    dots = do_fold(dots, *fold)
  return dots

def read_data(name):
  dots = set()
  folds = []

  with open(name, 'r') as f:
    for line in f:
      if line.strip() == '':
         break
      dots.add(tuple(map(int, line.split(','))))

    for line in f:
      direction, pos = line.split('=')
      folds.append((direction[-1], int(pos)))
  return dots, folds

def main():
   test_data = read_data('13-test.txt')
   data = read_data('13.txt')

   print(f'part1 test: {part1(test_data)}')
   print(f'part2 test:', end='')
   answer = part2(test_data)
   try:
     print(do_ocr(answer))
   except:
     print()
     print_dots(answer)
   
   print(f'part1 real: {part1(data)}')
   print(f'part2 real: ', end='')
   answer = part2(data)
   try:
     print(do_ocr(answer))
   except:
     print()
     print_dots(answer)

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
