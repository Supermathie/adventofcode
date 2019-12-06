require_relative '../lib/intcode'

describe Intcode do
  it "stops" do
    comp = Intcode.new [99]
    expect(comp.execute).to eq(99)
  end

  it "handles basic addition" do
    comp = Intcode.new [1,0,0,0,99]
    expect(comp.execute).to eq(2)
  end

  it "handles basic multiplication" do
    comp = Intcode.new [2,5,0,0,99,3]
    expect(comp.execute).to eq(6)

    comp = Intcode.new [2,4,4,0,99]
    expect(comp.execute).to eq(9801)
  end

  it "2a example 4" do
    comp = Intcode.new [1,1,1,4,99,5,6,0,99]
    expect(comp.execute).to eq(30)
  end
end
