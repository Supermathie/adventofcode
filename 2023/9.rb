#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/9.txt'
input = File.open(inputfile).readlines.map { _1.strip.split.map(&:to_i) }

def predict(seq)
  return 0 if seq.all? { _1 == 0 }
  seq.last + predict(seq.each_cons(2).map { _2 - _1 })
end

puts input.map { predict(_1) }.sum
puts input.map { predict(_1.reverse) }.sum