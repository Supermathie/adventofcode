#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/3.txt'

def find_symbol_near(symbols, x, xe, y)
  symbols.each do |symbol|
    _s_n, s_x, s_y = symbol
    return symbol if x-1 <= s_x && s_x <= xe + 1 && y-1 <= s_y && s_y <= y+1
  end
  nil
end

def find_parts_near(parts, x, y)
  parts.filter do |part|
    _p_n, p_x, p_xe, p_y = part
    x-1 <= p_xe && p_x <= x + 1 && y-1 <= p_y && p_y <= y+1
  end
end

def build_map(input)
  possible_parts = []
  parts = []
  nonparts = []
  symbols = []
  gears = []

  input.each_with_index do |line, y|
    line.scan(/\d+/) { x, xe = $~.offset(0); possible_parts << [_1.to_i, x, xe-1, y] }
    line.scan(/[^\d.]/) { x, = $~.offset(0); symbols << [_1, x, y] }
  end

  possible_parts.each do |part|
    _p_n, p_x, p_xe, p_y = part
    if find_symbol_near(symbols, p_x, p_xe, p_y)
      parts << part
    else
      nonparts << part
    end
  end

  symbols.each do |symbol|
    s_n, s_x, s_y = symbol
    next unless s_n == '*'

    nearby_parts = find_parts_near(parts, s_x, s_y)
    next unless nearby_parts.length == 2

    gears << [*symbol, nearby_parts[0][0] * nearby_parts[1][0]]
  end

  [parts, nonparts, symbols, gears]
end

input = File.open(inputfile).read.split("\n")
parts, _nonparts, _symbols, gears = build_map(input)
puts parts.map { _1[0] }.sum
puts gears.map { _1[3] }.sum
