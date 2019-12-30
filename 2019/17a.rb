#!/usr/bin/env ruby

require 'matrix'
require_relative 'lib/intcode_thread'

mapcomp = IntcodeThread.new(File.open('17.input').readline.strip.split(',').map(&:to_i))
mapcomp.execute.join
mapdata = "".tap { |s| s << mapcomp.output.pop until mapcomp.output.empty? }
map = map_to_matrix mapdata
intersections = map.each_with_index.filter do |e, i, j|
  [
    e != '.',
    map[i-1, j] != '.',
    map[i+1, j] != '.',
    map[i, j-1] != '.',
    map[i, j+1] != '.',
  ].all?
end

puts intersections.sum { |e, i, j| i * j }
