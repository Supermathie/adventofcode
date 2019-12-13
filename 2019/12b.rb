#!/usr/bin/env ruby

require 'matrix'

class Moon
  def initialize(pos = nil, vel = nil)
    @pos = pos || Vector[0, 0, 0]
    @vel = vel || Vector[0, 0, 0]
  end

  def ==(other)
    @pos == other.pos and @vel == other.vel
  end

  def to_s
    "pos=[#{@pos.to_a.join ','}], vel=[#{@vel.to_a.join ','}]"
  end

  def attract(other)
    d_vel = Vector[
      other.pos[0] <=> @pos[0],
      other.pos[1] <=> @pos[1],
      other.pos[2] <=> @pos[2],
    ]
    other.vel -= d_vel
    @vel += d_vel
  end

  def move
    @pos += @vel
  end

  def energy
    @pos.to_a.map(&:abs).sum * @vel.to_a.map(&:abs).sum
  end

  def hash_x
    [@pos[0], @vel[0]].hash
  end

  def hash_y
    [@pos[1], @vel[1]].hash
  end

  def hash_z
    [@pos[2], @vel[2]].hash
  end

  attr_accessor :pos, :vel
end

def run(filename)
  moons = Array.new
  File.open(filename).each_line do |line|
    x, y, z = line.match(/^<x= ?(-?\d+), y= ?(-?\d+), z= ?(-?\d+)>$/).captures.map(&:to_i)
    moons << Moon.new(Vector[x, y, z])
  end

  x_states = { moons.map(&:hash_x) => 0 }
  y_states = { moons.map(&:hash_y) => 0 }
  z_states = { moons.map(&:hash_z) => 0 }
  x_period = y_period = z_period = nil

  (1..).each do |i|
    moons.combination(2).each do |m1, m2|
      m1.attract(m2)
    end
    moons.each(&:move)

    if !x_period
      if x_states.has_key? moons.map(&:hash_x)
        x_period = i - x_states[moons.map(&:hash_x)]
      else
        x_states[moons.map(&:hash_x)] = i
      end
    end
    if !y_period
      if y_states.has_key? moons.map(&:hash_y)
        y_period = i - y_states[moons.map(&:hash_y)]
      else
        y_states[moons.map(&:hash_y)] = i
      end
    end
    if !z_period
      if z_states.has_key? moons.map(&:hash_z)
        z_period = i - z_states[moons.map(&:hash_z)]
      else
        z_states[moons.map(&:hash_z)] = i
      end
    end
    break if x_period and y_period and z_period
  end
  puts "#{x_period}, #{y_period}, #{z_period}"
  puts x_period.lcm(y_period).lcm(z_period)
end

run('12.input')
