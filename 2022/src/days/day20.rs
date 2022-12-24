use core::fmt;
use std::collections::VecDeque;
use std::hash::Hash;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

#[derive(Hash, Eq, PartialEq, Debug, Clone)]
struct Number {
    initial: usize,
    val: i64,
}

impl fmt::Display for Number {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.val)
    }
}

fn get_input(filename: &str) -> Vec<i64> {
    util::open_input(filename)
        .lines()
        .map(|line| {
            let line = line.unwrap();
            if let Some(result) = line.parse::<i64>().ok() {
                result
            } else {
                eprintln!("parsing:\n{}", line);
                panic!("failed!")
            }
        })
        .collect()
}

fn solve_for1(input: &Vec<i64>) -> i64 {
    let len = input.len();
    let mut q: VecDeque<Number> = input
        .iter()
        .enumerate()
        .map(|(pos, &val)| Number { initial: pos, val })
        .collect();
    for i in 0..len {
        if let Some((pos, _)) = q.iter().enumerate().find(|(_, num)| num.initial == i) {
            if q[pos].val == 0 {
                continue;
            }
            let num = q.remove(pos).unwrap();
            let mut new_pos = pos as i64 + num.val;
            new_pos = new_pos % (len-1) as i64;
            if new_pos < 0 {
                new_pos += len as i64 -1
            }
            q.insert(new_pos as usize, num);
        } else {
            panic!("pos not in vec");
        }
    }
    let (zero_pos, _) = q.iter().enumerate().find(|(_, num)| num.val == 0).unwrap();
    let p1 = (zero_pos + 1000) % len;
    let p2 = (zero_pos + 2000) % len;
    let p3 = (zero_pos + 3000) % len;

    // println!("{}, {}, {}", q[p1].val, q[p2].val, q[p3].val);

    q[p1].val + q[p2].val + q[p3].val
}

fn solve_for2(input: &Vec<i64>) -> i64 {
    let len = input.len();
    let mut q: VecDeque<Number> = input
        .iter()
        .enumerate()
        .map(|(pos, &val)| Number {
            initial: pos,
            val: val * 811589153,
        })
        .collect();
    for _ in 0..10 {
        for i in 0..len {
            if let Some((pos, _)) = q.iter().enumerate().find(|(_, num)| num.initial == i) {
                if q[pos].val == 0 {
                    continue;
                }
                let num = q.remove(pos).unwrap();
                let mut new_pos = pos as i64 + num.val;
                new_pos = new_pos % (len-1) as i64;
                if new_pos < 0 {
                    new_pos += len as i64 -1
                }
                q.insert(new_pos as usize, num);
            } else {
                panic!("pos not in vec");
            }
        }
    }
    let (zero_pos, _) = q.iter().enumerate().find(|(_, num)| num.val == 0).unwrap();
    let p1 = (zero_pos + 1000) % len;
    let p2 = (zero_pos + 2000) % len;
    let p3 = (zero_pos + 3000) % len;

    // println!("{}, {}, {}", q[p1].val, q[p2].val, q[p3].val);

    q[p1].val + q[p2].val + q[p3].val
}

pub fn solve() -> SolutionPair {
    let input = get_input("input/20.txt");
    let sol1 = solve_for1(&input);
    let sol2 = solve_for2(&input);
    (Solution::from(sol1), Solution::from(sol2))
}
