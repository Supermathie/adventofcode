use serde::{Deserialize, Serialize};
use std::cmp::Ordering;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

#[derive(Debug, Deserialize, Serialize, PartialEq, Eq, Clone)]
#[serde(untagged)]
enum PacketData {
    Int(u64),
    Packet(Packet),
}

impl PartialOrd for PacketData {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for PacketData {
    fn cmp(&self, other: &Self) -> Ordering {
        match (self, other) {
            (PacketData::Int(s), PacketData::Int(o)) => s.cmp(o),
            (PacketData::Packet(ref s), PacketData::Packet(ref o)) => s.cmp(o),
            (PacketData::Int(s), PacketData::Packet(ref o)) => Packet(vec![PacketData::Int(*s)]).cmp(o),
            (PacketData::Packet(ref s), PacketData::Int(o)) => s.cmp(&Packet(vec![PacketData::Int(*o)])),
        }
    }
}

#[derive(Debug, Deserialize, Serialize, PartialEq, Eq, Clone)]
struct Packet(Vec<PacketData>);

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        for (i, data) in self.0.iter().enumerate() {
            if let Some(odata) = other.0.get(i) {
                let ord = data.cmp(odata);
                if ord != Ordering::Equal {
                    return ord
                }
            } else {
                // other ran out of data
                return Ordering::Greater
            }
        }
        // self ran out of items, so it's either the same or smaller size
        self.0.len().cmp(&other.0.len())
    }
}

fn parse_line(line: &str) -> Packet {
    // println!("parsing: {}", line);
    let result = serde_json::from_str(line).unwrap();
    // println!("result: {:?}", result);
    result
}

fn get_input(filename: &str) -> Vec<Packet> {
    let input = util::open_input(filename).lines();

    let mut packets: Vec<Packet> = vec![];

    for line in input {
        let line = line.unwrap();
        if line.len() > 0 {
            packets.push(parse_line(&line));
        }
    }
    packets
}

fn solve_for1(input: &Vec<Packet>) -> usize {
    let mut sol = 0;

    for i in 0..input.len()/2 {
        let v1 = &input[i*2];
        let v2 = &input[i*2+1];
        let cmp = v1.cmp(v2);
        // print!("{} and {} -> {:?}\n", serde_json::to_string(v1).unwrap(), serde_json::to_string(v2).unwrap(), cmp);
        if cmp == Ordering::Less {
            sol += i + 1;
        }
    }
    sol
}

pub fn solve() -> SolutionPair {
    let mut input = get_input("input/13.txt");

    let sol1 = solve_for1(&input);
    let mut sol2 = 0;

    let div2 = parse_line("[[2]]");
    let div6 = parse_line("[[6]]");
    input.push(div2.clone());
    input.push(div6.clone());
    input.sort();
    for (i, p) in input.drain(..).enumerate() {
        if p == div2 {
            sol2 = i + 1;
        }
        if p == div6 {
            sol2 *= i + 1;
        }
    }

    (Solution::from(sol1), Solution::from(sol2))
}
