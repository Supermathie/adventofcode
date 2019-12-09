#!/usr/bin/env ruby
require 'set'

def travelled_locations(wirespec)
  pos = [0, 0]
  travelled = Set.new

  wirespec.each do |move|
    dir = move.slice(0)
    steps = move.slice(1..).to_i
    case dir
    when 'L'
      steps.times do |i|
        pos[0] -= 1
	travelled << pos.dup
      end
    when 'R'
      steps.times do |i|
        pos[0] += 1
	travelled << pos.dup
      end
    when 'U'
      steps.times do |i|
        pos[1] += 1
	travelled << pos.dup
      end
    when 'D'
      steps.times do |i|
        pos[1] -= 1
	travelled << pos.dup
      end
    end
  end
  travelled
end



crossed_locations = File.open('3.input') do |f|
  wire1 = travelled_locations(f.readline.chomp.split(','))
  wire2 = travelled_locations(f.readline.chomp.split(','))
  wire1.intersection wire2
end

nearest_location = crossed_locations.sort { |x, y| x[0].abs + x[1].abs <=> y[0].abs + y[1].abs }.first
puts nearest_location.map(&:abs).sum
