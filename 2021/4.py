#!/usr/bin/env python3

import itertools

def is_winner(board):
   return any([
      all(board[0:5]),
      all(board[5:10]),
      all(board[10:15]),
      all(board[15:20]),
      all(board[20:25]),
      all(board[::5]),
      all(board[1::5]),
      all(board[2::5]),
      all(board[3::5]),
      all(board[4::5]),
      #all(board[0::6]),     # oops, diagonals don't count
      #all(board[4::4][:5]), # oops, diagonals don't count
   ])
   
def part1(numbers, boards):
   board_hits = [ [False]*25 for _ in boards ] 
   for drawn_number in numbers:
      for board_num, board in enumerate(boards):
         try:
            board_hits[board_num][board.index(drawn_number)] = True
         except ValueError:
            pass
         else:
            if is_winner(board_hits[board_num]):
               #print(f'winner on {drawn_number} on board {board_num}:')
               #print(board)
               #print(board_hits[board_num])
               return sum( [ 0 if hit else num for hit, num in zip(board_hits[board_num], board) ] ) * drawn_number
   raise 'No winner?'

def part2(numbers, boards):
   board_hits = [ [False]*25 for _ in boards ] 
   won_boards = set()
   for drawn_number in numbers:
      for board_num, board in enumerate(boards):
         if board_num in won_boards:
            continue
         try:
            board_hits[board_num][board.index(drawn_number)] = True
         except ValueError:
            pass
         else:
            if is_winner(board_hits[board_num]):
               print(f'winner on {drawn_number} on board {board_num}:')
               won_boards.add(board_num)
               score = sum( [ 0 if hit else num for hit, num in zip(board_hits[board_num], board) ] ) * drawn_number
   return score

def read_data(name):
   with open(name, 'r') as f:
      numbers = [int(x) for x in f.readline().split(',')]
      boards = [ [ int(x) for x in board.split() ] for board in f.read().split('\n\n') ]
   return numbers, boards

def main():
   test_numbers, test_boards = read_data('4-test.txt')
   numbers, boards = read_data('4.txt')

   print(f'part1 test: {part1(test_numbers, test_boards)}')
   print(f'part2 test: {part2(test_numbers, test_boards)}')

   print(f'part1 real: {part1(numbers, boards)}')
   print(f'part2 real: {part2(numbers, boards)}')

import sys
if len(sys.argv) > 1 and sys.argv[1] == '--test':
    import doctest
    doctest.testmod()
else:
    main()
