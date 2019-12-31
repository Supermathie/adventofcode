#!/usr/bin/env ruby

require_relative 'lib/priority_queue'
require 'matrix'
require 'set'

class Map < Matrix
  def initialize(*args)
    @walls = ['#']
    @floors = ['.']
    super(*args)
  end

  def can_enter?(coord)
    !@walls.include? self[*coord]
  end

  def self.build_from_data(data)
    map = data.split("\n")
    height = map.length
    width = map.first.length
    build(height, width) { |y, x| map[y][x] }
  end

  def coord_of(char)
    each_with_index.filter { |e, _i, _j| e == char }.first[1..] rescue nil
  end

  def pois
    points = each_with_index.filter { |e, i, j| !(@walls + @floors).include? e }
    {}.tap { |h| points.each { |e, i, j| h[e] = [i, j] } }
  end

  def path_len(src, dst)
    queue = []
    queue << src
    distances = Map.build(row_count, column_count) { nil }
    distances[*src] = 0
    until queue.empty?
      cur = queue.shift
      neighbours(cur).each do |i2, j2|
        next unless can_enter? [i2, j2]

        if distances[i2, j2].nil?
          return distances[*cur] + 1 if [i2, j2] == dst
          distances[i2, j2] = distances[*cur] + 1
          queue << [i2, j2]
        end
      end
    end
    nil # no path found
  end

  def backtrack_path(distances, start)
    cur = start
    [].tap do |path|
      until distances[*cur].zero?
        path << cur
        cur = neighbours(cur).find { |i, j| distances[i, j] == distances[*cur] - 1 }
      end
    end.reverse
  end

  def path_to(src, dst, backtrack_func = self.method(:backtrack_path))
    queue = []
    queue << src
    distances = Map.build(row_count, column_count) { nil }
    distances[*src] = 0
    until queue.empty?
      cur = queue.shift
      neighbours(cur).each do |i2, j2|
        next unless can_enter? [i2, j2]

        if distances[i2, j2].nil?
          distances[i2, j2] = distances[*cur] + 1
          queue << [i2, j2]
          return backtrack_func.call(distances, dst) if [i2, j2] == dst
        end
      end
    end
    nil # no path found
  end

  private

  def neighbours(coord)
    i, j = coord
    [
      [i-1, j],
      [i+1, j],
      [i, j-1],
      [i, j+1],
    ]
  end

  attr_accessor :walls, :floors
end

class MapWithDoors < Map
  def initialize(*args)
    @keys = []
    super(*args)
  end

  def can_enter?(coord)
    return false if self[*coord].match?(/[A-Z]/) and !@keys.include? self[*coord].downcase
    super(coord)
  end

  def backtrack_keys(distances, start)
    cur = start
    [].tap do |keys|
      until distances[*cur].zero?
        keys << self[*cur].downcase if self[*cur].match?(/[A-Z]/)
        cur = neighbours(cur).find { |i, j| distances[i, j] == distances[*cur] - 1 }
      end
    end.reverse
  end

  def keys_to(src, dst)
    path_to(src, dst, backtrack_func = method(:backtrack_keys))
  end

  attr_accessor :keys
end

@map = MapWithDoors.build_from_data File.open('18.input').read
@all_keys_coords = @map.pois.filter { |k, v| k.match? /[a-z]/ }
@all_keys = @all_keys_coords.keys
@start = @map.coord_of '@'

# cache the path lengths between keys
# this assumes that the path length does not change with key possession
@map.keys = @all_keys
@all_path_lens = {}.tap do |h|
  @all_keys.combination 2 do |k1, k2|
    h[[k1, k2].sort] = @map.path_len(@all_keys_coords[k1], @all_keys_coords[k2])
  end
end

# calculate the list of keys required to reach each door
@required_keys = {}.tap do |h|
  @all_keys.each do |k|
    h[k] = @map.keys_to(@start, @all_keys_coords[k]).to_set
  end
end

# calculate the list of reachable keys given a keyring
def reachable_keys(keys)
  keyring = keys.to_set
  (@all_keys-keys).filter { |k| @required_keys[k] <= keyring } # subset
end

def search_for_keys
  visited_keyrings = {}
  pq = PriorityQueue.new
  pq.push([[], 0], 0)
  (1..).each do |i|
    held_keys, cost = pq.pop
    #puts "holding keys: #{held_keys.join ','} cost: #{cost}" if i % 1000 == 0
    sorted_keys = held_keys[0..-2].sort + (held_keys[-1..] || [])
    next if visited_keyrings.key? sorted_keys
    visited_keyrings[sorted_keys] = cost
    return held_keys, cost if held_keys.sort == @all_keys.sort

    pos = @all_keys_coords[held_keys.last] || @map.coord_of('@') # no keys? we're just starting
    @map.keys = held_keys
    #puts "  avail keys: #{avail_keys.join ','}"
    reachable_keys(held_keys).each do |k|
      new_cost = cost + (@all_path_lens[[held_keys.last, k].sort] rescue @map.path_len(pos, @all_keys_coords[k])) # super lazy here
      #puts "    new cost for #{k}: #{new_cost}"
      pq.push(
        [held_keys + [k], new_cost],
        new_cost
      )
    end
  end
end

order, cost = search_for_keys
puts "order: #{order.join ''}, cost: #{cost}"
