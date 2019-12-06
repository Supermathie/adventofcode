#!/usr/bin/env ruby

require_relative 'lib/intcode'

memory = File.open('5.input').readline.split(',').map(&:to_i)

puts Intcode.new(memory).execute(input: [5])
