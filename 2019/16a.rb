#!/usr/bin/env ruby

@pattern = Hash.new do |h, k|
  pat = [].tap do |l|
    (k+1).times { l << 0 }
    (k+1).times { l << 1 }
    (k+1).times { l << 0 }
    (k+1).times { l << -1 }
  end
  pat << pat.shift # 'ignore' the first 0
  h[k] = pat
end

def fft_round(input)
  input.length.times.map do |i|
    pattern = @pattern[i]
    input.map.with_index { |e, j| e * pattern[j % pattern.length] }.sum.abs % 10
  end
end

def fft(input, rounds: 1)
  output = input
  rounds.times { output = fft_round output }
  output
end

tests = [
  ["80871224585914546619083218645595", 24176176],
  ["19617804207202209144916044189917", 73745418],
  ["69317163492948606335995924319873", 52432133],
]

tests.each do |input, output|
  ans = fft(input.each_char.map(&:to_i), rounds: 100).slice(0..7).join.to_i
  puts "#{input} → #{ans} (should be #{output})" unless ans == output
end

input = File.open('16.input').readline.each_char.map(&:to_i)
puts(fft(input, rounds: 100).slice(0..7).join)
