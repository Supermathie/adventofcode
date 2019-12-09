#!/usr/bin/env ruby

@orbits = Hash.new

File.open('6.input').readlines.each do |l|
  a,b = l.chomp.split(')')
  @orbits[b] = a
end

def distance(d)
  return 0 if d == 'COM'
  distance(@orbits[d]) + 1
end

puts @orbits.keys.map { |x| distance(x) }.sum
