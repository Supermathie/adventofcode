#!/usr/bin/env ruby

require_relative 'lib/intcode_thread'

comp = IntcodeThread.new(File.open('13.input').readline.split(',').map(&:to_i))
comp[0] = 2 # play for free
comp.execute
tiles = Hash.new

score = nil
ball_x = 16

while comp.alive? or comp.output.length > 0
  x = comp.output.pop
  y = comp.output.pop
  tile = comp.output.pop

  if tile == 3 # paddle
    paddle_x = x
  elsif tile == 4 # ball
    ball_dir = x <=> ball_x
    ball_x = x
    ball_y = y
  end

  if x == -1 and y == 0
    score = tile
  else
    tiles[[x,y]] = tile
  end

  if comp.output.length == 0 and comp.status == "sleep" # waiting on input
    #puts "#{ball_x},#{ball_y} #{ball_dir} #{paddle_x} #{input}"
    comp.input << (ball_x <=> paddle_x)
  end
end

puts score
