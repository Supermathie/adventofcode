use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/4.txt").lines();

    let mut sol1: u64 = 0;
    let mut sol2: u64 = 0;

    for line in input {
        let line = line.unwrap();
        let pair: [[u64; 2]; 2] = line.split(',').map(
            |p| p.split('-').map(|i|
                i.parse::<u64>().unwrap()
            ).collect::<Vec<u64>>().try_into().unwrap()
        ).collect::<Vec<[u64; 2]>>().try_into().unwrap();

        if (pair[0][0] <= pair[1][0] && pair[0][1] >= pair[1][1]) ||
           (pair[1][0] <= pair[0][0] && pair[1][1] >= pair[0][1]) {
            sol1 += 1;
        }

        if  (pair[1][0] <= pair[0][0] && pair[0][0] <= pair[1][1]) ||
            (pair[0][0] <= pair[1][0] && pair[1][0] <= pair[0][1]) {
            sol2 += 1;
        }
    }

    (Solution::U64(sol1), Solution::U64(sol2))
}
