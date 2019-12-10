#!/usr/bin/env ruby

require 'set'

layout = File.open('10.input').readlines.map(&:chomp)
asteroids = [].tap do |a|
  layout.each_with_index do |line, i|
    line.chars.each_with_index do |col, j|
      a << [i, j] if col == '#'
    end
  end
end

best = asteroids.map do |candidate|
  angles = Set.new
  (asteroids-candidate).map do |a|
    slope = Float(candidate[0] - a[0]) / (candidate[1] - a[1])
    angles << [slope, candidate[0] > a[0]]
  end
  angles.size
end.max

puts best
