#!/usr/bin/env ruby

class Intcode < Array
  def execute(input: nil, trace: false)
    puts "#{self.to_a} [BEGIN]" if trace
    pc = 0
    output = []
    while self[pc] != 99 # halt
      opcode, modes = parse_opcode(self[pc])
      puts "#{self.to_a} pc=#{pc} opcode=#{opcode} modes=#{modes}" if trace
      if opcode == 1
        # add s0, s1, d
        s0 = self[pc+1]; s0 = self[s0] if modes[0].to_i == 0
        s1 = self[pc+2]; s1 = self[s1] if modes[1].to_i == 0
        puts "#{self.to_a} pc=#{pc} s0=#{s0} s1=#{s1} [ADD]" if trace
        self[self[pc+3]] = s0 + s1
        pc += 4
      elsif opcode == 2
        # mul s0, s1, d
        s0 = self[pc+1]; s0 = self[s0] if modes[0].to_i == 0
        s1 = self[pc+2]; s1 = self[s1] if modes[1].to_i == 0
        puts "#{self.to_a} pc=#{pc} s0=#{s0} s1=#{s1} [MUL]" if trace
        self[self[pc+3]] = s0 * s1
        pc += 4
      elsif opcode == 3
        # store [input], d
        self[self[pc+1]] = input.shift
        pc += 2
      elsif opcode == 4
        # output s0
        s0 = self[pc+1]; s0 = self[s0] if modes[0].to_i == 0
        output << s0
        pc += 2
      elsif opcode == 5
        # jnz s0, s1
        s0 = self[pc+1]; s0 = self[s0] if modes[0].to_i == 0
        s1 = self[pc+2]; s1 = self[s1] if modes[1].to_i == 0
        if s0 != 0
          pc = s1
        else
          pc += 3
        end
      elsif opcode == 6
        # jz s0, s1
        s0 = self[pc+1]; s0 = self[s0] if modes[0].to_i == 0
        s1 = self[pc+2]; s1 = self[s1] if modes[1].to_i == 0
        if s0 == 0
          pc = s1
        else
          pc += 3
        end
      elsif opcode == 7
        # tlt s0, s1, d
        s0 = self[pc+1]; s0 = self[s0] if modes[0].to_i == 0
        s1 = self[pc+2]; s1 = self[s1] if modes[1].to_i == 0
        self[self[pc+3]] = s0 < s1 ? 1 : 0
        pc += 4
      elsif opcode == 8
        # teq s0, s1, d
        s0 = self[pc+1]; s0 = self[s0] if modes[0].to_i == 0
        s1 = self[pc+2]; s1 = self[s1] if modes[1].to_i == 0
        self[self[pc+3]] = s0 == s1 ? 1 : 0
        pc += 4
      else
        raise StandardError, "Unknown opcode #{self[pc]} at pc: #{pc}"
      end
    end
    puts "#{self.to_a} pc=#{pc} [END]" if trace
    output
  end

  def parse_opcode(op)
    opcode = op % 100
    modes = (op / 100).to_s.each_char.map(&:to_i).reverse
    return opcode, modes
  end
end
