#!/usr/bin/env ruby
require 'set'

def travelled_locations(wirespec)
  pos = [0, 0]
  total_steps = 0
  travelled = Hash.new

  wirespec.each do |move|
    dir = move.slice(0)
    steps = move.slice(1..).to_i
    case dir
    when 'L'
      steps.times do |i|
        pos[0] -= 1
        total_steps += 1
        travelled[pos.dup] = total_steps unless travelled.include? pos
      end
    when 'R'
      steps.times do |i|
        pos[0] += 1
        total_steps += 1
        travelled[pos.dup] = total_steps unless travelled.include? pos
      end
    when 'U'
      steps.times do |i|
        pos[1] += 1
        total_steps += 1
        travelled[pos.dup] = total_steps unless travelled.include? pos
      end
    when 'D'
      steps.times do |i|
        pos[1] -= 1
        total_steps += 1
        travelled[pos.dup] = total_steps unless travelled.include? pos
      end
    end
  end
  travelled
end

wire1 = nil
wire2 = nil

crossed_locations = File.open('3.input') do |f|
  wire1 = travelled_locations(f.readline.chomp.split(','))
  wire2 = travelled_locations(f.readline.chomp.split(','))
  Set.new(wire1.keys).intersection Set.new(wire2.keys)
end

nearest_location = crossed_locations.sort do |x, y|
  wire1[x] + wire2[x] <=> wire1[y] + wire2[y]
end.first

puts wire1[nearest_location] + wire2[nearest_location]
