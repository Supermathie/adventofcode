#!/usr/bin/env ruby

inputfile = ARGV[0] || 'input/7.txt'

class Hand < Object
  def initialize(line)
    cards, bid = line.match(/([2-9TJQKA]{5}) (\d+)/).captures
    @cards = cards.split('')
    @bid = bid.to_i
  end

  attr_reader :cards
  attr_reader :bid

  def rankmap
    '23456789ABCDE'
  end

  def card_ranks
    @card_ranks ||= @cards.map { |c| c.tr('23456789TJQKA', rankmap) }
  end

  def type
    card_groups = @cards.each_with_object(Hash.new(0)) { |e, h| h[e] += 1 }
    @type ||=
      case card_groups.values.sort
      when [1, 1, 1, 1, 1]
        0
      when [1, 1, 1, 2]
        1
      when [1, 2, 2]
        2
      when [1, 1, 3]
        3
      when [2, 3]
        4
      when [1, 4]
        5
      when [5]
        6
      else
        raise 'unknown pattern'
      end
  end

  def <=>(other)
    [type, *card_ranks] <=> [other.type, *other.card_ranks]
  end

  def to_s
    "#{@cards.join('')} #{card_ranks.join('')} #{type} #{@bid}"
  end
end

class JokerHand < Hand
  def rankmap
    '23456789A0CDE'
  end

  def type
    card_groups = @cards.each_with_object(Hash.new(0)) { |e, h| h[e] += 1 }
    jokers = card_groups['J']
    @type ||=
      case card_groups.values.sort
      when [1, 1, 1, 1, 1]
        0 + jokers
      when [1, 1, 1, 2]
        case jokers
        when 0
          1
        else
          3
        end
      when [1, 2, 2]
        case jokers
        when 0
          2
        when 1
          4
        when 2
          5
        end
      when [1, 1, 3]
        case jokers
        when 0
          3
        else
          5
        end
      when [2, 3]
        case jokers
        when 0
          4
        else
          6
        end
      when [1, 4]
        case jokers
        when 0
          5
        else
          6
        end
      when [5]
        6
      else
        raise 'unknown pattern'
      end
  end
end

input = File.open(inputfile).read.split("\n")

hands = input.map { Hand.new(_1) }.sort
#puts hands
puts hands.each_with_index.map { |hand, i| hand.bid * (i+1) }.sum

jhands = input.map { JokerHand.new(_1) }.sort
#puts jhands
puts jhands.each_with_index.map { |hand, i| hand.bid * (i+1) }.sum