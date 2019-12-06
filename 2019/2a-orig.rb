#!/usr/bin/env ruby

memory = File.open('2.input').readline.split(',').map(&:to_i)
memory[1] = 12
memory[2] = 2

def execute(memory)
  pc = 0
  while memory[pc] != 99
    if memory[pc] == 1
      memory[memory[pc+3]] = memory[memory[pc+1]] + memory[memory[pc+2]]
    elsif memory[pc] == 2
      memory[memory[pc+3]] = memory[memory[pc+1]] * memory[memory[pc+2]]
    else
      raise StandardError
    end
    pc += 4
  end
end

execute(memory)
puts memory[0]
