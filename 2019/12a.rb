#!/usr/bin/env ruby

require 'matrix'

class Moon
  def initialize(pos = nil, vel = nil)
    @pos = pos || Vector[0, 0, 0]
    @vel = vel || Vector[0, 0, 0]
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

  attr_accessor :pos, :vel
end

def run(filename)
  moons = Array.new
  File.open(filename).each_line do |line|
    x, y, z = line.match(/^<x= ?(-?\d+), y= ?(-?\d+), z= ?(-?\d+)>$/).captures.map(&:to_i)
    moons << Moon.new(Vector[x, y, z])
  end

  puts "After 0 steps"
  puts moons.join("\n")
  puts

  1000.times do |i|
    moons.combination(2).each do |m1, m2|
      m1.attract(m2)
    end
    moons.each(&:move)

    if (i+1) % 10 == 0
      puts "After #{i+1} steps"
      puts moons.join("\n")
      puts
    end
  end
  puts moons.map(&:energy).sum
end

run('12.input')
