#!/usr/bin/env ruby

# reactions = {
#   "output_mat" => [
#     output_qty,
#     [ "input1_mat", input1_qty],
#     ["input2_mat", input2_qty],
#   ]
# }

class ReactionChamber
  def initialize(data)
    @reactions = {}.tap do |r|
      data.each_line do |line|
        input, output = line.chomp.split(' => ')
        output_mat, output_qty = output.split.reverse
        inputs = input.split(', ').map { |i| i.split.reverse }.each { |a| a[1] = a[1].to_i }
        r[output_mat] = [ output_qty.to_i, inputs ]
      end
    end
    clear!
  end

  def clear!
    @materials = {}
    @ore_used = 0
    @fuel_produced = 0
  end

  def consume_qty(mat, qty)
    if mat == 'ORE'
      @ore_used += qty
    else
      immed = [@materials.fetch(mat, 0), qty].min
      if immed > 0
        qty -= immed
        @materials[mat] -= immed
      end
      qty.times { consume mat }
    end
  end

  def consume(mat)
    if @materials.fetch(mat, 0) == 0
      produce(mat)
    end
    @materials[mat] -= 1
  end

  def produce(mat)
    output_qty, inputs = @reactions[mat]
    inputs.each do |input_mat, input_qty|
      consume_qty(input_mat, input_qty)
    end
    @materials[mat] = @materials.fetch(mat, 0) + output_qty
    @fuel_produced += 1 if mat == 'FUEL'
  end

  def mul=(qty)
    @materials.each_key { |k| @materials[k] *= qty }
    @ore_used *= qty
    @fuel_produced *= qty
  end

  def materials_empty?
    @materials.filter { |m, q| q != 0 }.empty?
  end

  def ore_required_for_material(mat, qty)
    return qty if mat == 'ORE'

    immed = [@materials.fetch(mat, 0), qty].min
    if immed > 0
      qty -= immed
      @materials[mat] -= immed
    end

    output_qty, inputs = @reactions[mat]
    scale = (qty.to_f/output_qty).ceil
    @materials[mat] = @materials.fetch(mat, 0) + output_qty * scale - qty
    inputs.map do |input_mat, input_qty|
      ore_required_for_material(input_mat, input_qty * scale)
    end.sum
  end

  attr_reader :ore_used, :materials, :fuel_produced
end

testinput = []
testinput << [<<EOF, 165]
9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL
EOF

testinput << [<<EOF, 13312]
157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT
EOF

testinput << [<<EOF, 180697]
2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF
EOF

testinput << [<<EOF, 2210736]
171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX
EOF

testinput.each_with_index do |data, i|
  rc = ReactionChamber.new(data[0])
  rc.produce("FUEL")
  puts "test #{i} (\#1) should be #{data[1]}: #{rc.ore_used}"
  rc.clear!
  puts "test #{i} (\#2) should be #{data[1]}: #{rc.ore_required_for_material('FUEL', 1)}"
end

rc = ReactionChamber.new(File.open('14.input'))
rc.consume("FUEL")
puts "#{rc.ore_used} for one FUEL"

rc.clear!

f2 = (1..100000000).bsearch { |i| rc.clear!; rc.ore_required_for_material('FUEL', i) > 1_000_000_000_000 } - 1
puts "#{f2} fuel produced with 1T ore"
