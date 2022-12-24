use sscanf;
use std::collections::HashMap;
use std::hash::Hash;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

#[derive(Hash, Eq, PartialEq, Debug, Clone)]
struct Blueprint {
    id: u8,
    ore_ore_cost: u32,
    clay_ore_cost: u32,
    obsidian_ore_cost: u32,
    obsidian_clay_cost: u32,
    geode_ore_cost: u32,
    geode_obsidian_cost: u32,
    ore_cap: u32,
    clay_cap: u32,
    obsidian_cap: u32,
}

#[derive(Hash, Eq, PartialEq, Debug, Clone, Default)]
struct State {
    time_left: u32,
    ore_robots: u32,
    clay_robots: u32,
    obsidian_robots: u32,
    geode_robots: u32,
    ore: u32,
    clay: u32,
    obsidian: u32,
    geode: u32,
}

fn get_input(filename: &str) -> Vec<Blueprint> {
    util::open_input(filename).lines().map(|line| {
        let line = line.unwrap();
        if let Some(result) = sscanf::sscanf!(line, "Blueprint {u8}: Each ore robot costs {u32} ore. Each clay robot costs {u32} ore. Each obsidian robot costs {u32} ore and {u32} clay. Each geode robot costs {u32} ore and {u32} obsidian.").ok() {
            let (id,ore_ore_cost,clay_ore_cost,obsidian_ore_cost,obsidian_clay_cost,geode_ore_cost,geode_obsidian_cost) = result;
            let ore_cap =
                    ore_ore_cost.
                max(clay_ore_cost).
                max(obsidian_ore_cost).
                max(geode_ore_cost);
            let clay_cap =
                obsidian_clay_cost;
            let obsidian_cap =
                obsidian_clay_cost;
        
            Blueprint { id, ore_ore_cost, clay_ore_cost, obsidian_ore_cost, obsidian_clay_cost, geode_ore_cost, geode_obsidian_cost, clay_cap, obsidian_cap, ore_cap }
        } else {
            println!("while on:\n{}", line);
            panic!("match failed!")
        }
    }).collect()
}

fn step(bp: &Blueprint, state: &State, memo: &mut HashMap<State, u32>) -> u32 {
    if state.time_left == 0 {
        return state.geode;
    }
    if let Some(val) = memo.get(state) {
        return *val;
    }
    let mut possible_next_states: Vec<State> = Default::default();
    if state.ore >= bp.ore_ore_cost && state.ore_robots < bp.ore_cap {
        possible_next_states.push(State {
            time_left: state.time_left - 1,
            ore_robots: state.ore_robots + 1,
            ore: state.ore + state.ore_robots - bp.ore_ore_cost,
            clay: state.clay + state.clay_robots,
            obsidian: state.obsidian + state.obsidian_robots,
            geode: state.geode + state.geode_robots,
            ..*state
        });
    }
    if state.ore >= bp.clay_ore_cost && state.clay_robots < bp.clay_cap {
        possible_next_states.push(State {
            time_left: state.time_left - 1,
            clay_robots: state.clay_robots + 1,
            ore: state.ore + state.ore_robots - bp.clay_ore_cost,
            clay: state.clay + state.clay_robots,
            obsidian: state.obsidian + state.obsidian_robots,
            geode: state.geode + state.geode_robots,
            ..*state
        });
    }
    if state.ore >= bp.obsidian_ore_cost && state.clay >= bp.obsidian_clay_cost && state.obsidian_robots < bp.obsidian_cap {
        possible_next_states.push(State {
            time_left: state.time_left - 1,
            obsidian_robots: state.obsidian_robots + 1,
            ore: state.ore + state.ore_robots - bp.obsidian_ore_cost,
            clay: state.clay + state.clay_robots - bp.obsidian_clay_cost,
            obsidian: state.obsidian + state.obsidian_robots,
            geode: state.geode + state.geode_robots,
            ..*state
        });
    }
    if state.ore >= bp.geode_ore_cost && state.obsidian >= bp.geode_obsidian_cost {
        possible_next_states.push(State {
            time_left: state.time_left - 1,
            geode_robots: state.geode_robots + 1,
            ore: state.ore + state.ore_robots - bp.geode_ore_cost,
            clay: state.clay + state.clay_robots,
            obsidian: state.obsidian + state.obsidian_robots - bp.geode_obsidian_cost,
            geode: state.geode + state.geode_robots,
            ..*state
        });
    }
    
    possible_next_states.push(State {
        time_left: state.time_left - 1,
        ore: state.ore + state.ore_robots,
        clay: state.clay + state.clay_robots,
        obsidian: state.obsidian + state.obsidian_robots,
        geode: state.geode + state.geode_robots,
        ..*state
    });

    let best = possible_next_states
        .iter()
        .map(|s| step(bp, s, memo))
        .reduce(|a, b| a.max(b))
        .unwrap();
    // println!("{:?}, best: {}", state, best);
    memo.insert(state.clone(), best);
    best
}

fn solve_for1(input: &Vec<Blueprint>) -> u32 {
    let initial = State {
        time_left: 24,
        ore_robots: 1,
        ..Default::default()
    };
    let memo: HashMap<State, u32> = Default::default();
    input.iter().map(|bp| {
        step(bp, &initial, &mut memo.clone()) * bp.id as u32
    }).reduce(|a, b| a + b).unwrap()
}

fn solve_for2(input: &Vec<Blueprint>) -> u32 {
    let initial = State {
        time_left: 32,
        ore_robots: 1,
        ..Default::default()
    };
    let memo: HashMap<State, u32> = Default::default();
    input[0..3].iter().map(|bp| {
        step(bp, &initial, &mut memo.clone()) as u32
    }).reduce(|a, b| a * b).unwrap()
}

pub fn solve() -> SolutionPair {
    let input = get_input("input/19.txt");
    let sol1 = solve_for1(&input);
    let sol2 = solve_for2(&input);
    (Solution::from(sol1), Solution::from(sol2))
}
