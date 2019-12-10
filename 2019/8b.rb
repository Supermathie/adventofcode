#!/usr/bin/env ruby
require 'pry'

width = 25
height = 6

layers = File.open('8.input').read.chomp.scan(/.{#{width*height}}/)
image = '2' * width * height

layers.each do |l|
  (width*height).times do |i|
    image[i] = l[i] if image[i] == '2'
  end
end

image.chars.each_slice(width) do |s|
  puts s.join.tr('01', ' •')
end
