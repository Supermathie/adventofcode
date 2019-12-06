#!/usr/bin/env ruby

require_relative 'lib/intcode'

memory = File.open('2.input').readline.split(',').map(&:to_i)
memory[1] = 12
memory[2] = 2

comp = Intcode.new(memory)
comp.execute
puts comp[0]
