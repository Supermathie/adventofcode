#!/usr/bin/env ruby
require 'set'

changes = File.open('1.input').map(&:to_i)

def find_freq(changes)
  freq = 0
  freq_set = Set.new
  
  while true
    changes.each do |c|
	    freq = freq + c
		  return freq if freq_set.include? freq
		  freq_set << freq
    end
  end
end

puts find_freq(changes)
