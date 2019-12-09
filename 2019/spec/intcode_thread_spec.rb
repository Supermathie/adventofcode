require 'socket'
require_relative '../lib/intcode_thread'

describe IntcodeThread do
  it "stops" do
    comp = IntcodeThread.new [99]
    comp.execute
    comp.join
    expect(comp[0]).to eq(99)
  end

  it "handles basic addition" do
    comp = IntcodeThread.new [1,0,0,0,99]
    comp.execute
    comp.join
    expect(comp[0]).to eq(2)
  end

  it "handles basic addition with immediate first parameter" do
    comp = IntcodeThread.new [101,50,0,0,99]
    comp.execute
    comp.join
    expect(comp[0]).to eq(151)
  end

  it "handles basic addition with immediate second parameter" do
    comp = IntcodeThread.new [1001,0,50,0,99]
    comp.execute
    comp.join
    expect(comp[0]).to eq(1051)
  end

  it "handles basic multiplication" do
    comp = IntcodeThread.new [2,5,0,0,99,3]
    comp.execute
    comp.join
    expect(comp[0]).to eq(6)
  end

  it "handles basic multiplication with immediate first parameter" do
    comp = IntcodeThread.new [102,10,5,0,99,3]
    comp.execute
    comp.join
    expect(comp[0]).to eq(30)
  end

  it "handles basic multiplication with immediate second parameter" do
    comp = IntcodeThread.new [1002,5,20,0,99,3]
    comp.execute
    comp.join
    expect(comp[0]).to eq(60)
  end

  it "handles basic larger multiplication" do
    comp = IntcodeThread.new [2,4,4,0,99]
    comp.execute
    comp.join
    expect(comp[0]).to eq(9801)
  end

  it "2a example 4" do
    comp = IntcodeThread.new [1,1,1,4,99,5,6,0,99]
    comp.execute
    comp.join
    expect(comp[0]).to eq(30)
  end

  it "supports input" do
    comp = IntcodeThread.new [3, 3, 99, 0]
    comp.execute(input: [42])
    comp.join
    expect(comp[3]).to eq(42)
  end

  it "supports negative numbers" do
    comp = IntcodeThread.new [1101,100,-1,4,0]
    comp.execute
    comp.join
    expect(comp[4]).to eq(99)
  end

  it "supports output" do
    comp = IntcodeThread.new [4, 3, 99, 42]
    output = comp.execute
    comp.join
    expect(comp.output.pop).to eq(42)
  end

  it "can use the equal instruction to test input" do
    comp = IntcodeThread.new [3,9,8,9,10,9,4,9,99,-1,8]
    comp1 = comp.dup.execute(input: [8])
    comp1.join
    expect(comp1.output.pop).to eq(1)
    comp2 = comp.dup.execute(input: [4])
    comp2.join
    expect(comp2.output.pop).to eq(0)
  end

  it "can use the less-than instruction to test input" do
    comp = IntcodeThread.new [3,9,7,9,10,9,4,9,99,-1,8]
    comp1 = comp.dup.execute(input: [4])
    expect(comp1.output.pop).to eq(1)
    comp2 = comp.dup.execute(input: [8])
    expect(comp2.output.pop).to eq(0)
  end

  it "can use the immed equal instruction to test input" do
    comp = IntcodeThread.new [3,3,1108,-1,8,3,4,3,99]
    comp1 = comp.dup.execute(input: [8])
    expect(comp1.output.pop).to eq(1)
    comp2 = comp.dup.execute(input: [4])
    expect(comp2.output.pop).to eq(0)
  end

  it "can use the immed less-than instruction to test input" do
    comp = IntcodeThread.new [3,3,1107,-1,8,3,4,3,99]
    comp1 = comp.dup.execute(input: [4])
    expect(comp1.output.pop).to eq(1)
    comp2 = comp.dup.execute(input: [8])
    expect(comp2.output.pop).to eq(0)
  end

  it "can use the jump test" do
    comp = IntcodeThread.new [3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9]
    comp1 = comp.dup.execute(input: [8])
    expect(comp1.output.pop).to eq(1)
    comp2 = comp.dup.execute(input: [0])
    expect(comp2.output.pop).to eq(0)
  end

  it "can use the jump test in immed mode" do
    comp = IntcodeThread.new [3,3,1105,-1,9,1101,0,0,12,4,12,99,1]
    comp1 = comp.dup.execute(input: [8])
    expect(comp1.output.pop).to eq(1)
    comp2 = comp.dup.execute(input: [0])
    expect(comp2.output.pop).to eq(0)
  end

end
