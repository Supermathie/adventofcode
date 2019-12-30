#!/usr/bin/env ruby

def fft_round!(input, offset: 0)
  # each digit ONLY depends on itself & successive digits so we can modify in-place
  if offset >= input.length
    (input.length-1).downto(0) do |i|
      input[i] = (input.fetch(i+1, 0) + input[i]) % 10
    end
  else
    input.map!.with_index do |e, i|
      # 0: 0, -2, 4, -6, 8, -10…
      # 1: 1, 2, -5, -6, 9, 10…
      # 2: 2, 3, 4, -8, -9, -10…
      # 3: 3, 4, 5, 6, -11, -12, -13, -14…
      #
      # 0: +0 len 1, -+2 len 1, +4 len 1, -+6 len 1…
      # 1: +0 len 2, -+4 len 2, +8 len 2,
      # 2: +0 len 3, -+6 len 3, +12 len 3,
      sum = 0
      runlen = i + offset + 1
      (i..input.length).step(runlen*2).each_with_index do |j, parity|
        sum += input.slice(j..(j+runlen-1)).sum * (parity % 2 == 0 ? 1 : -1)
      end
      sum.abs % 10
    end
  end
end

def fft(input, rounds: 1)
  output = input.dup
  rounds.times { fft_round! output }
  output
end

def fft_with_offset_header(input, rounds: 1)
  offset = input.slice(0..6).join.to_i
  output = (input * 10000).slice(offset..)
  rounds.times { fft_round!(output, offset: offset) }
  output
end

# part a
tests = [
  ["80871224585914546619083218645595", 24176176],
  ["19617804207202209144916044189917", 73745418],
  ["69317163492948606335995924319873", 52432133],
  [File.open('16.input').readline.chomp, 89576828]
]

tests.each do |input, output|
  ans = fft(input.each_char.map(&:to_i), rounds: 100).slice(0..7).join.to_i
  puts "#{input} → #{ans} (should be #{output})" unless ans == output
end

# part b
tests2 = [
  ["03036732577212944063491565474664", 84462026],
  ["02935109699940807407585447034323", 78725270],
  ["03081770884921959731165446850517", 53553731],
]

tests2.each do |input, output|
  ans = fft_with_offset_header(input.each_char.map(&:to_i), rounds: 100).slice(0..7).join.to_i
  puts "#{input} → #{ans} (should be #{output})" unless ans == output
end

# 93242278 too high
input = File.open('16.input').readline.chomp.each_char.map(&:to_i)
output = fft_with_offset_header(input, rounds: 100)
puts(output.slice(0..7).join.to_i)
