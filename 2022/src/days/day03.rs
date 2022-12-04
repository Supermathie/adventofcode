use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};
//use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

fn to_bitmap(items: &[u8]) -> u64 {
    let mut score: u64 = 0;
    for item in items {
        let val = if (b'a'..=b'z').contains(item) {
            item - b'a'
        } else if (b'A'..=b'Z').contains(item) {
            (item - b'A') + 26
        } else {
            panic!("unknown item");
        };

        score |= 1 << val;
    }
    score
}

fn find_bit(val: u64) -> u8 {
    let mut bit: u8 = 0;

    loop {
        if bit == 64 {
            panic!("no bits found")
        }
        if val & (1 << bit) != 0 {
            break;
        }
        bit += 1
    }
    bit
}

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/3.txt").lines();

    let mut sol1: u64 = 0;
    let mut sol2: u64 = 0;

    let mut groupscores: [u64; 3] = [0, 0, 0];
    for (lineno, line) in input.enumerate() {
        let line = line.unwrap();
        let items = line.trim().as_bytes();

        let lscore: u64 = to_bitmap(&items[0..(items.len()/2)]);
        let rscore: u64 = to_bitmap(&items[(items.len()/2)..]);
        sol1 += find_bit(lscore & rscore) as u64 + 1;

        let group_mem = lineno % 3;
        groupscores[group_mem] = lscore | rscore;

        if group_mem == 2 {
            sol2 += find_bit(groupscores[0] & groupscores[1] & groupscores[2]) as u64 + 1
        }
    }

    (Solution::U64(sol1), Solution::U64(sol2))
}
