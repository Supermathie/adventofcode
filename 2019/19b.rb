#!/usr/bin/env ruby

require_relative 'lib/intcode_thread'

comp = IntcodeThread.new(File.open('19.input').readline.split(',').map(&:to_i))
beamdata = Hash.new { |h, k| h[k] = comp.dup.execute(input: [*k.reverse]).output.pop }

def valid_upper_left_coord(x, y, beamdata)
  [
    beamdata[[x,    y   ]] == 1,
    beamdata[[x+99, y   ]] == 1,
    beamdata[[x,    y+99]] == 1,
    beamdata[[x+99, y+99]] == 1,
  ].all?
end

x = (0..10000).bsearch do |x|
  min_y = (0..10000).drop_while { |y| beamdata[[x,y]] == 0 }.first
  if min_y
    max_y = (min_y..10000).drop_while { |y| beamdata[[x,y]] == 1 }.first
    max_y = max_y ? max_y - 1 : 10000
    !!(min_y..max_y).drop_while do |y|
      !valid_upper_left_coord(x, y, beamdata)
    end.first
  end
end

min_y = (0..10000).drop_while { |y| beamdata[[x,y]] == 0 }.first
max_y = (min_y..10000).drop_while { |y| beamdata[[x,y]] == 1 }.first - 1
y = (min_y..max_y).drop_while do |y|
  !valid_upper_left_coord(x, y, beamdata)
end.first

puts "(#{x},#{y}) #{x * 10000 + y}"