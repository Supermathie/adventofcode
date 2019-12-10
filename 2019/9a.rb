#!/usr/bin/env ruby

require_relative 'lib/intcode_thread'

memory = File.open('9.input').readline.split(',').map(&:to_i)
comp = IntcodeThread.new(memory).execute(input: [1])
comp.join
while comp.output.length > 0
  puts comp.output.pop
end
