#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/2.txt'

def parse_event(event)
  r = event[/\d+ red/].to_i
  g = event[/\d+ green/].to_i
  b = event[/\d+ blue/].to_i
  [r, g, b]
end

def parse_game(line)
  _, num, rest = *line.match(/Game (\d+): (.*)/)
  events = rest.split(';')
  max_r, max_g, max_b = 0, 0, 0
  events.each do |event|
    r, g, b = parse_event(event)
    max_r = [max_r, r].max
    max_g = [max_g, g].max
    max_b = [max_b, b].max
  end
  [num.to_i, max_r, max_g, max_b]
end

input = File.open(inputfile).read.split("\n")
games = input.map{ parse_game _1 }
possible_games = games.filter { _, r, g, b = *_1; r <= 12 && g <= 13 && b <= 14 }
puts possible_games.map{ _1[0] }.sum
puts games.map { _, r, g, b = *_1; r*g*b }.sum