#!/usr/bin/env ruby

require 'matrix'
require 'set'
require_relative 'lib/intcode_thread'

class Robot
  def initialize(filename)
    @comp = IntcodeThread.new(File.open(filename).readline.split(',').map(&:to_i))
    @position = Vector[0, 0]
    @map = {}
    @oxygen_location = nil
    @steps = []
  end

  def dirvec(direction)
    Vector[*{
      up:    [0,  1],
      right: [1,  0],
      down:  [0, -1],
      left:  [-1, 0],
    }[direction]]
  end

  def move(direction)
    @comp.input.push({
      up:    1,
      down:  2,
      left:  3,
      right: 4,
    }[direction])
    result = @comp.output.pop

    delta = dirvec direction

    case result
    when 0 # wall
      @map[@position + delta] = 0
    when 1 # open
      @position += delta
      @map[@position] = 1
      @steps.push direction
    when 2 # oxygen system
      @position += delta
      @map[@position] = 2
      @oxygen_location = @position
      @steps.push direction
    end
    result
  end

  def first_unknown_neighbour
    [:up, :right, :down, :left].each do |dir|
      return dir unless @map.has_key?(@position + dirvec(dir))
    end
    nil
  end
  
  def backtrack
    direction = @steps.pop
    @comp.input.push({ # reversed!
      up:    2,
      down:  1,
      left:  4,
      right: 3,
    }[direction])
    result = @comp.output.pop
    raise StandardError, 'tried to backtrack into a wall' if result == 0
    
    @position -= dirvec direction # reversed!
  end
  
  def find_oxygen_dfs
    @comp.execute
    while @oxygen_location.nil?
      dir = first_unknown_neighbour
      if dir.nil?
        backtrack
      else
        move dir
      end
    end
    @comp.kill
    @oxygen_location
  end

  def find_oxygen_bfs(depth: 1)
    @comp.execute
    while @oxygen_location.nil?
      dir = first_unknown_neighbour
      if dir.nil? || @steps.length >= depth
        if @steps.empty?
          @comp.kill
          return nil
        else
          backtrack
        end
      else
        move dir
      end
      # sleep 0.05
      # puts print_map(clear: true)
    end
    @comp.kill
    @oxygen_location
  end


  def print_map(clear: false)
    range  = Range.new(*@map.keys.map{ |p| p[0] }.minmax)
    domain = Range.new(*@map.keys.map{ |p| p[1] }.minmax)
    output = String.new
    output << "\033[H\033[2J" if clear
    domain.to_a.reverse.each do |y|
      range.each do |x|
        if Vector[x,y] == @position
          output << '@'
        elsif Vector[x,y] == Vector[0, 0]
          output << '<'
        else
          output << { 0 => '#', 1 => '.', 2 => 'O', nil => ' ' }[@map[Vector[x, y]]]
        end
      end
      output << "\n" # newline
    end
    output
  end
  attr_accessor :position, :oxygen_location, :steps
end

robot = Robot.new('15.input')
puts robot.find_oxygen_dfs
maxlen = robot.steps.length
puts robot.print_map

depth = (1..maxlen).bsearch do |i|
  robot = Robot.new('15.input') 
  oxygen = robot.find_oxygen_bfs(depth: i)
  !!oxygen
end

puts
puts robot.print_map
puts "depth: #{depth}"
