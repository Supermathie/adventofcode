#!/usr/bin/env ruby

def check(boxid)
  h = {}.tap do |h|
    boxid.split('').each do |c|
      h[c] = h.fetch(c, 0) + 1
    end
  end
  return (h.has_value? 2), (h.has_value? 3)
end
  

count_2 = 0
count_3 = 0
File.open('2.input').map(&:chomp).each do |boxid|
  has_2, has_3 = check(boxid)
  count_2 += 1 if has_2
  count_3 += 1 if has_3
end

puts count_2 * count_3
  
