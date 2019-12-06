#!/usr/bin/env ruby

memory = File.open('2.input').readline.split(',').map(&:to_i)

def execute(memory)
  pc = 0
  while memory[pc] != 99
    if memory[pc] == 1
      memory[memory[pc+3]] = memory[memory[pc+1]] + memory[memory[pc+2]]
      pc += 4
    elsif memory[pc] == 2
      memory[memory[pc+3]] = memory[memory[pc+1]] * memory[memory[pc+2]]
      pc += 4
    else
      raise StandardError
    end
  end
  memory[0]
end

(0..99).each do |i|
  (0..99).each do |j|
    mem = memory.dup
    mem[1] = i
    mem[2] = j
    if execute(mem) == 19690720
      puts i,j
    end
  end
end
