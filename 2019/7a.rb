#!/usr/bin/env ruby

require_relative 'lib/intcode'

amplifier = Intcode.new(File.open('7.input').readline.split(',').map(&:to_i))

max_signal = (0..4).to_a.permutation.map do |p|
  signal = 0
  5.times do
    signal = amplifier.dup.execute(input: [p.shift, signal]).first
  end
  signal
end.max

puts max_signal
