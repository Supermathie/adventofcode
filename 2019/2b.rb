#!/usr/bin/env ruby

require_relative 'lib/intcode'

memory = File.open('2.input').readline.split(',').map(&:to_i)

(0..99).each do |i|
  (0..99).each do |j|
    comp = Intcode.new(memory.dup)
    comp[1] = i
    comp[2] = j
    comp.execute
    if comp[0] == 19690720
      puts "#{i} #{j}"
    end
  end
end
