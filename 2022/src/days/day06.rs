use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

fn sop_marker(data: &[u8]) -> bool {
    return
    data[0] != data[1] && data[0] != data[2] && data[0] != data[3] &&
    data[1] != data[2] && data[1] != data[3] &&
    data[2] != data[3]
}

fn som_marker(data: &[u8]) -> bool {
    let mut v = Vec::from(data);
    v.sort();
    for seq in v.windows(2) {
        if seq[0] == seq[1] {
            return false
        }
    }
    true
}

pub fn solve() -> SolutionPair {
    let mut input = util::open_input("input/6.txt").lines();

    let mut sol1: u64 = 0;
    let mut sol2: u64 = 0;

    let data = input.next().unwrap().unwrap();
    let data = data.as_bytes();
    for (pos, seq) in data.windows(4).enumerate() {
        if sop_marker(seq) {
            sol1 = pos as u64 + 4; // pos is start of window, +3 is end, +1 for total # of chars
            break;
        }
    }
    for (pos, seq) in data.windows(14).enumerate() {
        if som_marker(seq) {
            sol2 = pos as u64 + 14; // pos is start of window, +13 is end, +1 for total # of chars
            break;
        }
    }

    (Solution::from(sol1), Solution::from(sol2))
}
