#!/usr/bin/env ruby

require_relative 'lib/intcode_thread'

comp = IntcodeThread.new(File.open('17.input').readline.strip.split(',').map(&:to_i))
comp[0] = 2 # override robot boot sequence
comp.execute

# L,12,R,8,L,6,R,8,L,6,R,8,L,12,L,12,R,8,L,12,R,8,L,6,R,8,L,6,L,12,R,8,L,6,R,8,L,6,R,8,L,12,L,12,R,8,L,6,R,6,L,12,R,8,L,12,L,12,R,8,L,6,R,6,L,12,L,6,R,6,L,12,R,8,L,12,L,12,R,8
# A,B,A,A,B,C,B,C,C,B
# A: L,12,R,8,L,6,R,8,L,6
# B: R,8,L,12,L,12,R,8
# C: L,6,R,6,L,12

display = false
[
  'A,B,A,A,B,C,B,C,C,B',
  'L,12,R,8,L,6,R,8,L,6',
  'R,8,L,12,L,12,R,8',
  'L,6,R,6,L,12',
  display ? 'y' : 'n',
].each do |s|
  s.each_byte { |b| comp.input.push b }
  comp.input.push 10 # \n
end

while comp.alive? or comp.output.length > 1
  putc comp.output.pop
end

puts comp.output.pop