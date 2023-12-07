#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/6.txt'

input = File.open(inputfile).read.split("\n")
times = input[0].scan(/\d+/).map(&:to_i)
dists = input[1].scan(/\d+/).map(&:to_i)
races = times.zip(dists)

p1 = races.map do |t, d|
  (0..t).filter { |i| i * (t-i) > d }.count
end.inject(:*)

puts p1

time = input[0].scan(/\d+/).join('').to_i
dist = input[1].scan(/\d+/).join('').to_i

t1 = (0..time).detect { |i| i * (time-i) > dist }
puts time - t1*2+1