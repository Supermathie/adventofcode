#!/usr/bin/env ruby

class IntcodeThread < Array
  def execute(input: nil, output: nil, trace: false)
    @input  = input  || Queue.new
    @output = output || Queue.new

    @thread = Thread.new {
      @pc = 0
      puts "#{self[@pc..@pc+3]} [BEGIN]" if trace
      while self[@pc] != 99 # halt
        opcode, modes = parse_opcode(self[@pc])
        puts "#{self[@pc..@pc+3]} pc=#{@pc} opcode=#{opcode} modes=#{modes}" if trace
        if opcode == 1
          # add s0, s1, d
          s0 = self[@pc+1]; s0 = self[s0] if modes[0].to_i == 0
          s1 = self[@pc+2]; s1 = self[s1] if modes[1].to_i == 0
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} [ADD]" if trace
          self[self[@pc+3]] = s0 + s1
          @pc += 4
        elsif opcode == 2
          # mul s0, s1, d
          s0 = self[@pc+1]; s0 = self[s0] if modes[0].to_i == 0
          s1 = self[@pc+2]; s1 = self[s1] if modes[1].to_i == 0
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} [MUL]" if trace
          self[self[@pc+3]] = s0 * s1
          @pc += 4
        elsif opcode == 3
          # store [input], d
          val = @input.pop
          puts "#{self[@pc..@pc+2]} pc=#{@pc} loc=#{self[@pc+1]} old=#{self[self[@pc+1]]} new=#{val} [STO]" if trace
          self[self[@pc+1]] = val
          @pc += 2
        elsif opcode == 4
          # output s0
          s0 = self[@pc+1]; s0 = self[s0] if modes[0].to_i == 0
          puts "#{self[@pc..@pc+2]} pc=#{@pc} s0=#{s0} [OUT]" if trace
          @output << s0
          @pc += 2
        elsif opcode == 5
          # jnz s0, s1
          s0 = self[@pc+1]; s0 = self[s0] if modes[0].to_i == 0
          s1 = self[@pc+2]; s1 = self[s1] if modes[1].to_i == 0
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} [JNZ]" if trace
          if s0 != 0
            @pc = s1
          else
            @pc += 3
          end
        elsif opcode == 6
          # jz s0, s1
          s0 = self[@pc+1]; s0 = self[s0] if modes[0].to_i == 0
          s1 = self[@pc+2]; s1 = self[s1] if modes[1].to_i == 0
          puts "#{self[@pc..@pc+2]} pc=#{@pc} s0=#{s0} s1=#{s1} [JEZ]" if trace
          if s0 == 0
            @pc = s1
          else
            @pc += 3
          end
        elsif opcode == 7
          # tlt s0, s1, d
          s0 = self[@pc+1]; s0 = self[s0] if modes[0].to_i == 0
          s1 = self[@pc+2]; s1 = self[s1] if modes[1].to_i == 0
          val = s0 < s1 ? 1 : 0
          self[self[@pc+3]] = val
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} d=#{val} [TLT]" if trace
          @pc += 4
        elsif opcode == 8
          # teq s0, s1, d
          s0 = self[@pc+1]; s0 = self[s0] if modes[0].to_i == 0
          s1 = self[@pc+2]; s1 = self[s1] if modes[1].to_i == 0
          val = s0 == s1 ? 1 : 0
          self[self[@pc+3]] = val
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} d=#{val} [TEQ]" if trace
          @pc += 4
        else
          raise StandardError, "Unknown opcode #{self[@pc]} at pc: #{@pc}"
        end
      end
      puts "#{self[@pc..@pc+3]} pc=#{@pc} [END]" if trace
    }
    self
  end

  def alive?
    @thread.alive?
  end

  def join
    @thread.join
  end

  attr_accessor :input, :output

  def parse_opcode(op)
    opcode = op % 100
    modes = (op / 100).to_s.each_char.map(&:to_i).reverse
    return opcode, modes
  end
end
