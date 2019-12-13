#!/usr/bin/env ruby

class IntcodeThread < Array
  def resolv_addr(addr, mode)
    case mode
    when 0
      self[addr]
    when 1
      addr
    when 2
      @reladd + self[addr]
    end
  end

  def [](addr)
    # we want unwritten memory to have a value of 0
    super || 0
  end

  def execute(input: nil, output: nil, trace: false)
    @input  = input  || Queue.new
    @output = output || Queue.new
    @reladd = 0

    @thread = Thread.new {
      @pc = 0
      puts "#{self[@pc..@pc+3]} [BEGIN]" if trace
      while self[@pc] != 99 # halt
        opcode, modes = parse_opcode(self[@pc])
        puts "#{self[@pc..@pc+3]} pc=#{@pc} opcode=#{opcode} modes=#{modes}" if trace
        if opcode == 1
          # add s0, s1, d
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          s1 = self[resolv_addr(@pc+2, modes[1].to_i)]
          d = resolv_addr(@pc+3, modes[2].to_i)
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} d=#{d} [ADD]" if trace
          self[d] = s0 + s1
          @pc += 4
        elsif opcode == 2
          # mul s0, s1, d
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          s1 = self[resolv_addr(@pc+2, modes[1].to_i)]
          d = resolv_addr(@pc+3, modes[2].to_i)
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} d=#{d} [MUL]" if trace
          self[d] = s0 * s1
          @pc += 4
        elsif opcode == 3
          # store [input], d
          d = resolv_addr(@pc+1, modes[0].to_i)
          val = @input.pop
          puts "#{self[@pc..@pc+2]} pc=#{@pc} d=#{d} old=#{self[d]} new=#{val} [STO]" if trace
          self[d] = val
          @pc += 2
        elsif opcode == 4
          # output s0
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          puts "#{self[@pc..@pc+2]} pc=#{@pc} s0=#{s0} [OUT]" if trace
          @output << s0
          @pc += 2
        elsif opcode == 5
          # jnz s0, s1
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          s1 = self[resolv_addr(@pc+2, modes[1].to_i)]
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} [JNZ]" if trace
          if s0 != 0
            @pc = s1
          else
            @pc += 3
          end
        elsif opcode == 6
          # jz s0, s1
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          s1 = self[resolv_addr(@pc+2, modes[1].to_i)]
          puts "#{self[@pc..@pc+2]} pc=#{@pc} s0=#{s0} s1=#{s1} [JEZ]" if trace
          if s0 == 0
            @pc = s1
          else
            @pc += 3
          end
        elsif opcode == 7
          # tlt s0, s1, d
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          s1 = self[resolv_addr(@pc+2, modes[1].to_i)]
          d = resolv_addr(@pc+3, modes[2].to_i)
          val = s0 < s1 ? 1 : 0
          self[d] = val
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} d=#{d} [TLT]" if trace
          @pc += 4
        elsif opcode == 8
          # teq s0, s1, d
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          s1 = self[resolv_addr(@pc+2, modes[1].to_i)]
          d = resolv_addr(@pc+3, modes[2].to_i)
          val = s0 == s1 ? 1 : 0
          self[d] = val
          puts "#{self[@pc..@pc+3]} pc=#{@pc} s0=#{s0} s1=#{s1} d=#{d} [TEQ]" if trace
          @pc += 4
        elsif opcode == 9
          # sra s0
          s0 = self[resolv_addr(@pc+1, modes[0].to_i)]
          puts "#{self[@pc..@pc+1]} pc=#{@pc} s0=#{s0} [SRA]" if trace
          @reladd += s0
          @pc += 2
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

  def status
    @thread.status
  end

  attr_accessor :input, :output

  def parse_opcode(op)
    opcode = op % 100
    modes = (op / 100).to_s.each_char.map(&:to_i).reverse
    return opcode, modes
  end
end
