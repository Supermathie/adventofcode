#!/usr/bin/env ruby

layout = File.open('10.input').readlines.map(&:chomp)
asteroids = [].tap do |a|
  layout.each_with_index do |line, i|
    line.chars.each_with_index do |col, j|
      a << [j, i] if col == '#'
    end
  end
end

def vector(src, dst)
  # rotate to get 0 degrees at "up"
  angle = Math.atan2(dst[1] - src[1], dst[0] - src[0]) * 180 / Math::PI - 270
  dist = (dst[0]-src[0]).abs + (dst[1]-src[1]).abs
  while angle < 0
    angle += 360
  end
  [angle, dist] # not the actual distance but it's good enough
end

source = [11,13]
asteroids.delete(source)

# radix sorted by angle then distance
vectors = asteroids
  .map { |a| [*vector(source, a), a] }
  .sort_by { |angle, dist, a| dist }
  .sort_by { |angle, dist, a| angle }

targets = [].tap do |targets|
  while !vectors.empty?
    last_angle = -1
    new_vectors = vectors.dup
    vectors.each do |a|
      if a[0] != last_angle
        last_angle = a[0]
        new_vectors.delete a
        targets << a[2]
        puts "Asteroid at #{a[2].join ','} destroyed at angle #{a[0]}" if targets.length < 20
      end
    end
    vectors = new_vectors
  end
end
puts targets[199].join(',')
