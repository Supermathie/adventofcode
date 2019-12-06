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

class Integer
  def each_digit
    self.to_s.each_char { |c| yield c.to_i }
  end
end

def is_valid(n)
  found_double = false
  prev_digit = -1
  n.each_digit do |d|
    found_double = true if prev_digit == d
    return false if prev_digit > d
    prev_digit = d
  end
  return found_double
end

puts "X: 111111 is not valid" unless is_valid(111111)
puts "X: 223450 is valid" if is_valid(223450)
puts "X: 123789 is valid" if is_valid(123789)

puts (284639..748759).map { |n| is_valid n }.count { |n| n }
