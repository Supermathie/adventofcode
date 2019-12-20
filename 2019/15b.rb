#!/usr/bin/env ruby

require 'matrix'
require_relative 'lib/intcode_thread'

class MapRobot
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

  def map_area
    @comp.execute
    max_depth = 0
    while true
      dir = first_unknown_neighbour
      if dir.nil?
        return max_depth if @steps.empty?
        backtrack
      else
        move dir
      end
      max_depth = [@steps.length, max_depth].max
    end
    @comp.kill
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
          output << { 0 => '#', 1 => '.', 2 => 'O', nil => '#' }[@map[Vector[x, y]]]
        end
      end
      output << "\n" # newline
    end
    output
  end
  
  def map_to_matrix
    map = print_map.split("\n")
    height = map.length
    width = map.first.length
    Matrix.build(height, width) { |y, x| map[y][x] }
  end
  attr_accessor :position, :oxygen_location, :steps
end

class PathRobot
  # uses a coordinate system of y, x to match a matrix i = row, j = column
  def initialize(map, position: nil)
    @map = map.dup
    @visited = Matrix.build(@map.row_count, @map.column_count) { false }
    @position = position || Vector[0, 0]
    @steps = []
  end

  def dirvec(direction)
    Vector[*{
      up:    [-1,  0],
      right: [ 0,  1],
      down:  [ 1,  0],
      left:  [ 0, -1],
    }[direction]]
  end

  def first_unvisited_neighbour
    [:up, :right, :down, :left].each do |dir|
      pos = @position + dirvec(dir)
      return dir unless @map[*pos] == '#' or @visited[*pos]
    end
    nil
  end

  def move(direction)
    delta = dirvec direction

    if @map[*(@position + dirvec(direction))] == '#'
      raise StandardError, 'tried to move into a wall'
    end

    @position += delta
    @steps.push direction
    @visited[*@position] = true
    self
  end

  def backtrack
    raise StandardError, 'cannot backtrack when no steps taken' if @steps.empty?
    direction = @steps.pop
    delta = dirvec direction
    if @map[*(@position - delta)] == '#' # reversed!
      raise StandardError, 'tried to backtrack into a wall'
    end

    @position -= delta # reversed!
    self
  end

  def visited_all?
    @map.count { |e| e != '#' } == @visited.count { |e| e }
  end

  def find_max_depth_dfs
    max_depth = 0
    until visited_all?
      dir = first_unvisited_neighbour
      if dir.nil?
        backtrack
      else
        move dir
      end
      max_depth = [@steps.length, max_depth].max
    end
    max_depth
  end

  attr_accessor :position, :steps
end

maprobot = MapRobot.new('15.input')
maprobot.map_area

pathrobot = PathRobot.new(maprobot.map_to_matrix, position: Vector[33, 3]) # location of oxygen in rebased coordinates
puts pathrobot.find_max_depth_dfs