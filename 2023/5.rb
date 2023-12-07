#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/5.txt'

class RangeMap
  def initialize()
    @ranges = {}
  end
  attr_reader :ranges

  def add(src, dst, length)
    @ranges[src..src+length-1] = dst-src
  end

  def find_next_range(n)
    @ranges.to_a.sort_by { _1[0].begin }.detect { _1[0].begin >= n }
  end

  def find_range(n)
    @ranges.each do |r, d|
      return r, d if r.include?(n)
    end
    return nil, nil
  end

  def [](n)
    r, d = find_range(n)
    r ? n + d : n
  end

  def map_range(r)
    r2, d2 = find_range(r.begin)
    ans = if r2
      if r2.cover? r
        [(r.begin + d2)..(r.end + d2)]
      else
        [(r.begin + d2)..(r2.end + d2)] + map_range(r2.end+1..r.end)
      end
    else
      # start does not overlap
      next_range = find_next_range(r.begin)&.[]0
      if next_range && next_range.begin <= r.end
        [(r.begin..next_range.begin-1)] + map_range(next_range.begin..r.end)
      else
        # nothing overlaps
        [r]
      end
    end
    #puts "map_range(#{r.inspect}) with #{@ranges.inspect}"
    #puts "  #{ans.inspect}"
    ans
  end

  def map_ranges(ranges)
    ranges.map { map_range(_1) }.flatten
  end
end

def seed_to_location(seed, maps)
  n = seed
  %w[
    seed-to-soil
    soil-to-fertilizer
    fertilizer-to-water
    water-to-light
    light-to-temperature
    temperature-to-humidity
    humidity-to-location
  ].each { n = maps[_1][n] }
  n
end

def seed_range_to_location(seed_range, maps)
  seed_ranges = [seed_range]
  %w[
    seed-to-soil
    soil-to-fertilizer
    fertilizer-to-water
    water-to-light
    light-to-temperature
    temperature-to-humidity
    humidity-to-location
  ].each { seed_ranges = maps[_1].map_ranges(seed_ranges) }
  seed_ranges
end

input = File.open(inputfile)
seeds = input.readline().strip().split(': ')[1].split.map(&:to_i)
maps = {}

curmap = nil
input.each_line do |line|
  case line.strip
  when ''
    next
  when /(.*) map:/
    curmap = RangeMap.new()
    maps[$~[1]] = curmap
  when /^(\d+) (\d+) (\d+)$/
    dst, src, len = $~.captures.map(&:to_i)
    curmap.add(src, dst, len)
  else
    raise "wut #{line}"
  end
end

puts seeds.map { seed_to_location(_1, maps) }.min

seed_ranges = seeds.each_slice(2).map { _1.._1+_2-1 }
puts seed_ranges.map { seed_range_to_location(_1, maps) }.flatten.map(&:begin).min