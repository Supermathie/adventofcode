#!/usr/bin/env ruby

# It is a six-digit number.
# The value is within the range given in your puzzle input.
# Two adjacent digits are the same (like 22 in 122345).
# Going from left to right, the digits never decrease; they only ever increase
# or stay the same (like 111123 or 135679).
#
# Other than the range rule, the following are true:
#
# 111111 meets these criteria (double 11, never decreases).
# 223450 does not meet these criteria (decreasing pair of digits 50).
# 123789 does not meet these criteria (no double).
#
# An Elf just remembered one more important detail: the two adjacent matching
# digits are not part of a larger group of matching digits.
# 
# Given this additional criterion, but still ignoring the range rule, the following are now true:
# 
# 112233 meets these criteria because the digits never decrease and all
# repeated digits are exactly two digits long.
# 123444 no longer meets the criteria (the repeated 44 is part of a larger
# group of 444).
# 111122 meets the criteria (even though 1 is repeated more than twice, it
# still contains a double 22).


def is_valid(n)
  digits = n.to_s.each_char.map(&:to_i)
  found_double = false
  digits.each_with_index do |d, i|
    found_double = true if digits[i-1] == d and
                           digits[i-2] != d and
                           digits[i+1] != d
    return false if i > 0 and digits[i-1] > d
  end
  return found_double
end

puts "X: 111111 is valid" if is_valid(111111)
puts "X: 111122 is not valid" unless is_valid(111122)
puts "X: 112233 is not valid" unless is_valid(112233)
puts "X: 223450 is valid" if is_valid(223450)
puts "X: 123789 is valid" if is_valid(123789)
puts "X: 123444 is valid" if is_valid(123444)

puts (284639..748759).map { |n| is_valid n }.count { |n| n }
