#!/usr/bin/env ruby

puts File.open('1.input').map { |w| (w.to_i/3).floor - 2 }.sum
