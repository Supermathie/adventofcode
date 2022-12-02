use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    // Your solution here...
    let input = util::open_input("input/1.txt").lines();

    let mut cur:u64 = 0;
    let mut totals: Vec<u64> = Vec::new();

    for line in input {
        let val = line.expect("uhoh: ");
        if val == "" {
            totals.push(cur);
            cur = 0;
        } else {
            cur += val.parse::<u64>().expect("this should be a number!")
        }
    }

    totals.sort();
    let sol1 = totals.pop().expect("not enough elves");
    let sol2 = sol1 + totals.pop().expect("not enough elves") + totals.pop().expect("not enough elves");

    (Solution::U64(sol1), Solution::U64(sol2))
}

