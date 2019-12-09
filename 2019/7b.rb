#!/usr/bin/env ruby

require 'pry'
require_relative 'lib/intcode_thread'

amplifier = IntcodeThread.new(File.open('7.input').readline.split(',').map(&:to_i))

max_signal = (5..9).to_a.permutation.map do |p|
  inputs = 5.times.map { Queue.new }
  5.times do |i|
    inputs[i] << p[i]
  end
  amp_a = amplifier.dup.execute(input: inputs[0], output: inputs[1])
  amp_b = amplifier.dup.execute(input: inputs[1], output: inputs[2])
  amp_c = amplifier.dup.execute(input: inputs[2], output: inputs[3])
  amp_d = amplifier.dup.execute(input: inputs[3], output: inputs[4])
  amp_e = amplifier.dup.execute(input: inputs[4])
  inputs[0] << 0 # first signal
  output = amp_e.output
  while amp_e.alive?
    signal = output.pop
    break unless amp_e.alive?
    amp_a.input << signal
  end
  signal
end.max

puts max_signal
