#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/4.txt'

class Card
  def initialize(line)
    _, num, numbers, winning = *line.match(/Card +(\d+): (.*)\|(.*)/)
    @num = num.to_i
    @numbers = numbers.strip.split(/ +/).compact.map(&:to_i)
    @winning = winning.strip.split(/ +/).compact.map(&:to_i)
  end

  attr_reader :num

  def matches
    @matches ||= @numbers.filter { @winning.include? _1 }.count
  end

  def score
    @score ||= matches > 0 ? 2 ** (matches-1) : 0
  end
end

input = File.open(inputfile).read.split("\n")
cards = input.map{ Card.new _1 }
puts cards.map(&:score).sum

hand = {}
cards.each { hand[_1.num] = 1 }

cards.each do |card|
  ((card.num+1)..[card.num+card.matches, cards.last.num].min).each do |n|
    hand[n] += hand[card.num]
  end
end
puts hand.values.sum