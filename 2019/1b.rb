#!/usr/bin/env ruby

def fuel_for(weight)
  return 0 if weight == 0
  fuel = [(weight/3).floor - 2, 0].max
  fuel + fuel_for(fuel)
end

puts File.open('1.input').map{|w| fuel_for w.to_i }.sum
