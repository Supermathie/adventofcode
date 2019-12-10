#!/usr/bin/env ruby
require 'pry'

width = 25
height = 6

layers = File.open('8.input').read.chomp.scan(/.{#{width*height}}/)
most_0 = layers.sort_by { |l| l.count('0') }.first
puts most_0.count('1') * most_0.count('2')
