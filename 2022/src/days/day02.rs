use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

//use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

fn lookup(mv: char) -> i64 {
    match mv {
        'A' => 0,
        'B' => 1,
        'C' => 2,
        'X' => 0,
        'Y' => 1,
        'Z' => 2,
        _ => panic!("unknown move")
    }
}

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/2.txt").lines();

    let mut score1:i64 = 0;
    let mut score2:i64 = 0;
    for line in input {
        let val = line.unwrap();
        let m1 = lookup(val.chars().nth(0).unwrap());
        let m2 = lookup(val.chars().nth(2).unwrap());

        match (m2 - m1 + 3) % 3 {
            0 => {
                // draw :/
                score1 += m2 + 1 + 3;
            }
            1 => {
                // win :)
                score1 += m2 + 1 + 6;
            }
            2 => {
                // lose :(
                score1 += m2 + 1 + 0;
            }
            e => panic!("uhoh: {}", e)
        }
        match m2 {
            0 => {
                // lose :(
                score2 += (m1 + 2) % 3 + 1 + 0;
            }
            1 => {
                // draw :/
                score2 += m1 + 1 + 3 ;
            }
            2 => {
                // win :)
                score2 += (m1 + 1) % 3 + 1 + 6 ;
            }
            e => panic!("uhoh: {}", e)
        }
    }

    (Solution::from(score1), Solution::from(score2))
}
