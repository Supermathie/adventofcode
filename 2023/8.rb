#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/8.txt'

input = File.open(inputfile)
directions = input.readline.strip.tr('LR', '01').split('').map(&:to_i)
input.readline # blank line

map = {}
input.read.split("\n").each do |line|
  start, l, r = line.match(/([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)/).captures.map(&:to_sym)
  map[start] = [l, r]
end


# part 1
steps = 0
position = :AAA
instructions = directions.cycle

while position != :ZZZ
  position = map[position][instructions.next]
  steps += 1
end

puts steps

def steps_to_finish(start, directions, map)
  position = start
  steps = 0
  while !position.end_with? 'Z'
    position = map[position][directions.next]
    steps += 1
  end
  steps
end

starts = map.keys.filter { _1.end_with? 'A' }
puts starts.map { |pos| steps_to_finish(pos, directions.cycle, map) }.reduce(&:lcm)