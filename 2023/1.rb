#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/1.txt'

def unnumber(s)
  s
  .gsub('one', '1')
  .gsub('two', '2')
  .gsub('three', '3')
  .gsub('four', '4')
  .gsub('five', '5')
  .gsub('six', '6')
  .gsub('seven', '7')
  .gsub('eight', '8')
  .gsub('nine', '9')
end

def score(s)
  a = s.scan(/(?=(one|two|three|four|five|six|seven|eight|nine|[0-9]))/).flatten
  score = unnumber(a[0]).to_i * 10 + unnumber(a[-1]).to_i
  #puts [s, a, score].join(' ')
  score
end

input = File.open(inputfile).read.split("\n")
puts input.map { n = _1.scan(/[0-9]/).map(&:to_i); n[0] * 10 + n[-1] }.sum
puts input.map { score(_1) }.sum
