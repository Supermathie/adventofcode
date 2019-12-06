#!/usr/bin/env ruby

class Intcode < Array

  def execute
    pc = 0
    while self[pc] != 99 # halt
      if self[pc] == 1
        # add s1, s2, d
        self[self[pc+3]] = self[self[pc+1]] + self[self[pc+2]]
        pc += 4
      elsif self[pc] == 2
        # mul s1, s2, d
        self[self[pc+3]] = self[self[pc+1]] * self[self[pc+2]]
        pc += 4
      else
        raise StandardError, "Unknown opcode #{self[pc]} at pc: #{pc}"
      end
    end
    self[0]
  end
end
