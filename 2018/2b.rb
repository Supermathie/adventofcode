#!/usr/bin/env ruby

def match(id1, id2)
  found_difference = false
  (0..id1.length).each do |i|
    if id1[i] != id2[i]
      return false if found_difference
      found_difference = true
    end
  end
  return true
end

def findpair(boxids)
  boxids.each do |id1|
    boxids.each do |id2|
      next if id1 == id2
      return [id1, id2] if match(id1, id2)
    end
  end
end
  
boxids = File.open('2.input').map(&:chomp)
id1, id2 = findpair(boxids)
puts id1,id2

l = [].tap do |a|
  (0..id1.length).each do |i|
    a << id1[i] if id1[i] == id2[i]
  end
end
puts l.join('')
