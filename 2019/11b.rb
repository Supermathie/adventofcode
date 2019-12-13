#!/usr/bin/env ruby

require 'pry'
require 'matrix'
require 'set'
require_relative 'lib/intcode_thread'

class Vector
  def last
    self[self.size-1]
  end
end

class Robot
  def initialize
    @direction = :up
    @position = Vector[0, 0]
  end

  def turn_left
    @direction = {
      :up    => :left,
      :left  => :down,
      :down  => :right,
      :right => :up,
    }[@direction]
  end

  def turn_right
    @direction = {
      :up    => :right,
      :right => :down,
      :down  => :left,
      :left  => :up,
    }[@direction]
  end

  def forward
    @position += Vector[*{
      :up    => [0,  1],
      :right => [1,  0],
      :down  => [0, -1],
      :left  => [-1, 0],
    }[@direction]]
  end

  def position
    @position.to_a
  end
end

comp = IntcodeThread.new(File.open('11.input').readline.split(',').map(&:to_i))
comp.execute
panels = Hash.new
robot = Robot.new

input = comp.input
output = comp.output
panels[robot.position] = :white

while comp.alive? or comp.output.length > 0
  colour = panels.fetch(robot.position, :black)
  comp.input << (colour == :black ? 0 : 1)
  panels[robot.position] = comp.output.pop == 0 ? :black : :white
  if comp.output.pop == 0
    robot.turn_left
  else
    robot.turn_right
  end
  robot.forward
end

min_x, max_x = panels.keys.map(&:first).minmax
min_y, max_y = panels.keys.map(&:last).minmax


(min_y..max_y).to_a.reverse.each do |y|
  (min_x..max_x).each do |x|
    colour = panels.fetch([x, y], :black)
    print({ :black => ' ', :white => '×' }[colour])
  end
  puts
end
