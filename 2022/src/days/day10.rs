use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

fn get_px(reg_x: i64, cycle: i64) -> &'static str {
    let col = (cycle - 1) % 40;
    if (reg_x - col).abs() < 2 {
        if cycle % 40 == 0 {
            "#\n"
        } else {
            "#"
        }
    } else {
        if cycle % 40 == 0 {
            " \n"
        } else {
            " "
        }
    }
}

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/10.txt").lines();

    let mut reg_x: i64 = 1;
    let mut cycle = 1;
    let mut sol1: i64 = 0;
    let mut sol2 = String::from("\n");

    for line in input {
        let line = line.unwrap();
        let cmd: Vec<&str> = line.trim().split(' ').collect();
        match cmd[..] {
            ["noop"] => {
                sol2 += get_px(reg_x, cycle);
                if (cycle + 20) % 40 == 0 {
                    sol1 += cycle * reg_x;
                }
                cycle += 1;
            },
            ["addx", count] => {
                sol2 += get_px(reg_x, cycle);
                if (cycle + 20) % 40 == 0 {
                    sol1 += cycle * reg_x;
                }
                cycle += 1;
                sol2 += get_px(reg_x, cycle);
                if (cycle + 20) % 40 == 0 {
                    sol1 += cycle * reg_x;
                }
                cycle += 1;
                reg_x += count.parse::<i64>().unwrap();
            }
            [..] => panic!("bad command: {:?}", cmd),
        }
    }

    (Solution::from(sol1), Solution::from(sol2))
}
