#!/usr/bin/env ruby

require_relative 'lib/intcode_thread'

comp = IntcodeThread.new(File.open('19.input').readline.split(',').map(&:to_i))
data = Hash.new { |h, k| h[k] = comp.dup.execute(input: [*k]).output.pop }
puts 50.times.map { |y| 50.times.map { |x| data[[x, y]] }.sum }.sum
