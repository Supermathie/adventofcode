use regex::Regex;

use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

enum InspectionOp {
    Mul,
    Add,
    Pow,
}

struct Monkey {
    // ord: u8,
    num_inspected: u64,
    items: Vec<u64>,
    inspection_op: InspectionOp,
    inspection_val: u64,
    test_divisor: u64,
    true_monkey: usize,
    false_monkey: usize,
}

impl Monkey {
    #[allow(dead_code)]
    fn inspect_and_throw_items(&mut self) -> (u64, usize) {
        (0, 0)
    }
}

const MONKEY_PARSER: &str = r"Monkey (\d+):
  Starting items: (\d+(?:, \d+)*)
  Operation: new = old (.*)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)";

fn parse_monkey(descr: &str) -> Monkey {
    let re = Regex::new(MONKEY_PARSER).unwrap();
    // print!("parsing:\n{:?}\n---\n", descr);
    let groups = match re.captures(descr) {
        None => panic!("cannot parse: \n{}\n---\n", descr),
        Some(groups) => groups,
    };
    let (_, items, inspection, test_divisor, true_monkey, false_monkey) = (
        &groups[1], &groups[2], &groups[3], &groups[4], &groups[5], &groups[6],
    );

    // let ord: u8 = ord.parse().unwrap();
    let items: Vec<u64> = items
        .split(", ")
        .map(|i| i.parse::<u64>().unwrap())
        .collect();
    // let items = Box::from(items);
    let inspection: Vec<&str> = inspection.split(" ").collect();
    let (inspection_op, inspection_val) = match inspection[..] {
        ["+", num] => (InspectionOp::Add, num.parse::<u64>().unwrap()),
        ["*", "old"] => (InspectionOp::Pow, 2),
        ["*", num] => (InspectionOp::Mul, num.parse::<u64>().unwrap()),
        [..] => panic!("cannot parse inspection operation"),
    };
    let test_divisor = test_divisor.parse().unwrap();
    let true_monkey = true_monkey.parse().unwrap();
    let false_monkey = false_monkey.parse().unwrap();
    Monkey {
        // ord: ord,
        num_inspected: 0,
        items: items,
        inspection_op: inspection_op,
        inspection_val: inspection_val,
        test_divisor: test_divisor,
        true_monkey: true_monkey,
        false_monkey: false_monkey,
    }
}

pub fn solve() -> SolutionPair {
    (Solution::from(solve1()), Solution::from(solve2()))
}

pub fn solve1() -> u64 {
    let input = std::fs::read_to_string("input/11.txt").unwrap();
    let mut monkeys: Vec<Monkey> = input.split("\n\n").map(|spec| parse_monkey(spec)).collect();

    for _ in 0..20 {
        for m in 0..monkeys.len() {
            let items = monkeys[m].items.clone();
            monkeys[m].items.clear();

            for item in items {
                // inspect the item
                monkeys[m].num_inspected += 1;
                let item = match monkeys[m].inspection_op {
                    InspectionOp::Add => item + monkeys[m].inspection_val,
                    InspectionOp::Mul => item * monkeys[m].inspection_val,
                    InspectionOp::Pow => item * item,
                };
                // done inspecting
                let item = item / 3;
                // throw item
                let dest = if item % monkeys[m].test_divisor == 0 {
                    monkeys[m].true_monkey
                } else {
                    monkeys[m].false_monkey
                };
                monkeys[dest].items.push(item);
            }
        }
    }

    let mut scores: Vec<u64> = monkeys.iter().map(|m| m.num_inspected).collect();
    scores.sort();

    scores[scores.len() - 2] * scores[scores.len() - 1]
}

pub fn solve2() -> u64 {
    let input = std::fs::read_to_string("input/11.txt").unwrap();
    let mut monkeys: Vec<Monkey> = input.split("\n\n").map(|spec| parse_monkey(spec)).collect();
    let mut divisor = 1;
    for div in monkeys.iter().map(|m| m.test_divisor) {
        divisor *= div;
    }

    for _ in 0..10000 {
        for m in 0..monkeys.len() {
            let items = monkeys[m].items.clone();
            monkeys[m].items.clear();

            for item in items {
                // inspect the item
                monkeys[m].num_inspected += 1;
                let item = match monkeys[m].inspection_op {
                    InspectionOp::Add => item + monkeys[m].inspection_val,
                    InspectionOp::Mul => item * monkeys[m].inspection_val,
                    InspectionOp::Pow => item * item,
                };
                // done inspecting
                let item = item % divisor;
                // throw item
                let dest = if item % monkeys[m].test_divisor == 0 {
                    monkeys[m].true_monkey
                } else {
                    monkeys[m].false_monkey
                };
                monkeys[dest].items.push(item);
            }
        }
    }

    let mut scores: Vec<u64> = monkeys.iter().map(|m| m.num_inspected).collect();
    scores.sort();

    scores[scores.len() - 2] * scores[scores.len() - 1]
}
