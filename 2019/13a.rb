#!/usr/bin/env ruby

require_relative 'lib/intcode_thread'

comp = IntcodeThread.new(File.open('13.input').readline.split(',').map(&:to_i))
comp.execute
tiles = Hash.new

while comp.alive? or comp.output.length > 0
  x = comp.output.pop
  y = comp.output.pop
  tile = comp.output.pop

  tiles[[x,y]] = tile
end

puts tiles.values.count(2)
