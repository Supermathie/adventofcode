use sscanf;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////
const NUM_STACKS: usize = 9;

pub fn solve() -> SolutionPair {
    let mut stacks1: [Vec<u8>; NUM_STACKS] = Default::default();

    let mut input = util::open_input("input/5.txt").lines();

    loop {
        let line = input.next().unwrap().unwrap();
        let crates = line.as_bytes();
        if crates[1] == b'1' {
            input.next();
            break;
        }

        for i in 0..NUM_STACKS {
            let crate_id = crates[i * 4 + 1];
            if crate_id != b' ' {
                stacks1[i].push(crate_id);
            }
        }
    }

    for i in 0..NUM_STACKS {
        stacks1[i].reverse();
    }

    let mut stacks2 = stacks1.clone();

    for line in input {
        let line = line.unwrap();
        let (count, src, dst) = sscanf::sscanf!(line, "move {usize} from {usize} to {usize}").unwrap();

        for _ in 0..count {
            // 9000
            let gripper = stacks1[src-1].pop().unwrap(); 
            stacks1[dst-1].push(gripper);
        }

        // 9001
        let mut gripper = stacks2[src-1].split_off(stacks2[src-1].len() - count);
        stacks2[dst-1].append(&mut gripper);
    }

    let sol1_bytes = &stacks1.map(|s| s[s.len()-1]);
    let sol2_bytes = &stacks2.map(|s| s[s.len()-1]);

    let sol1 = std::str::from_utf8(sol1_bytes).unwrap();
    let sol2 = std::str::from_utf8(sol2_bytes).unwrap();

    (Solution::from(sol1), Solution::from(sol2))
}
