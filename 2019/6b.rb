#!/usr/bin/env ruby

@orbits = Hash.new

File.open('6.input').readlines.each do |l|
  a,b = l.chomp.split(')')
  @orbits[b] = a
end

def distance_to(d, dest: 'COM')
  return 0 if d == dest
  distance(@orbits[d]) + 1
end

def ancestors(d)
  return [] if d == 'COM'
  a = @orbits[d]
  ancestors(a) << a
end

you_ancestors = ancestors('YOU')
san_ancestors = ancestors('SAN')

while you_ancestors.first == san_ancestors.first
  you_ancestors.shift
  san_ancestors.shift
end

puts you_ancestors.count + san_ancestors.count
