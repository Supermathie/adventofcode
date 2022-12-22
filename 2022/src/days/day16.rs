use priority_queue::PriorityQueue;
use sscanf;

use std::collections::HashMap;
use std::collections::HashSet;
use std::hash::Hash;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

struct Room {
    flow_rate: u16,
    tunnels: Vec<u16>,
}

#[derive(Hash, Eq, PartialEq, Clone)]
struct State {
    turn: u8,
    room: u16,
    eroom: u16,
    released: u16,
    flow_rate: u16,
    open: Vec<u16>,
}

impl State {
    fn max_potential(&self, total_valve_pressure: u16, max_turns: u8) -> u16 {
        self.released + (max_turns - self.turn) as u16 * total_valve_pressure
    }
}

fn get_input(filename: &str) -> HashMap<u16, Room> {
    let mut rooms: HashMap<u16, Room> = Default::default();
    let input = util::open_input(filename).lines();

    for line in input {
        let line = line.unwrap();
        let (desc, flow_rate, _, _, _, tunnels) = sscanf::sscanf!(line, r"Valve {u16:r36} has flow rate={u16}; tunnel{String:/s?/} lead{String:/s?/} to valve{String:/s?/} {String}").unwrap();
        let tunnels: Vec<u16> = tunnels.split(", ").map(|s| u16::from_str_radix(s, 36).unwrap()).collect();
        // println!("Valve {}, flow rate {}, tunnels: {:?}", desc, flow_rate, tunnels);
        rooms.insert(desc, Room { flow_rate, tunnels });
    }
    rooms
}

fn solve_for1(rooms: &HashMap<u16, Room>) -> u16 {
    let total_valve_pressure: u16 = rooms.values().map(|r| r.flow_rate).sum();
    let total_useful_valves: u16 = rooms.values().map(|r| r.flow_rate.clamp(0, 1)).sum();
    let max_turns = 30;

    let mut queue = PriorityQueue::new();
    let initial_room = u16::from_str_radix("AA", 36).unwrap();
    let initial = State {
        turn: 0,
        room: initial_room,
        eroom: initial_room,
        released: 0,
        flow_rate: 0,
        open: Vec::new(),
    };

    let max_potential = initial.max_potential(total_valve_pressure, max_turns);
    queue.push(initial, max_potential);
    
    let mut sol = 0;
    while let Some((state, prio)) = queue.pop() {
        if prio < sol {
            // this cannot be better than an existing solution
            continue
        }
        if state.open.len() as u16 == total_useful_valves {
            // all valves are open! just wait
            sol = sol.max(prio);
            continue
        }
        if state.turn == max_turns {
            // time's up!
            continue
        }
        let room = &rooms[&state.room];
        if room.flow_rate > 0 && !state.open.contains(&state.room) {
            // consider opening this valve
            let mut open = state.open.clone();
            open.push(state.room);
            open.sort();
            let new_state = State {
                turn: state.turn+1,
                released: state.released + state.flow_rate,
                flow_rate: state.flow_rate + room.flow_rate,
                open,
                ..state
            };
            let max_potential = new_state.max_potential(total_valve_pressure, max_turns);
            sol = sol.max(new_state.released);
            queue.push(new_state, max_potential);
        }
        for new_room in &room.tunnels {
            // consider moving to a new room
            let new_state = State {
                turn: state.turn+1,
                released: state.released + state.flow_rate,
                room: *new_room,
                open: state.open.clone(),
                ..state
            };
            let max_potential = new_state.max_potential(total_valve_pressure, max_turns);
            sol = sol.max(new_state.released);
            queue.push(new_state, max_potential);
        }
    }
    sol
}

fn solve_for2(rooms: &HashMap<u16, Room>, sol_min: u16) -> u16 {
    let total_valve_pressure: u16 = rooms.values().map(|r| r.flow_rate).sum();
    let total_useful_valves: u16 = rooms.values().map(|r| r.flow_rate.clamp(0, 1)).sum();
    let mut visited_states: HashSet<State> = Default::default();

    let max_turns = 26;

    let mut queue = PriorityQueue::new();
    let initial_room = u16::from_str_radix("AA", 36).unwrap();
    let initial = State {
        turn: 0,
        room: initial_room,
        eroom: initial_room,
        released: 0,
        flow_rate: 0,
        open: Vec::new(),
    };

    let max_potential = initial.max_potential(total_valve_pressure, max_turns);
    queue.push(initial, max_potential);
    
    let mut sol = sol_min;
    while let Some((state, prio)) = queue.pop() {
        if prio < sol || visited_states.contains(&state) {
            // this cannot be better than an existing solution
            continue
        }
        if state.open.len() as u16 == total_useful_valves {
            // all valves are open! just wait
            return sol.max(prio);
        }
        visited_states.insert(state.clone());
        if state.turn == max_turns {
            // time's up!
            continue
        }
        let orig_flow_rate = state.flow_rate;
        // elephant actions first
        let mut estates: Vec<State> = vec![];
        {
            let eroom = &rooms[&state.eroom];
            if eroom.flow_rate > 0 && !state.open.contains(&state.eroom) {
                // consider opening this valve
                let mut open = state.open.clone();
                open.push(state.eroom);
                open.sort();
                let new_state = State {
                    open,
                    flow_rate: state.flow_rate + eroom.flow_rate,
                    ..state
                };
                estates.push(new_state);
            }
            for new_room in &eroom.tunnels {
                // consider moving to a new room
                let new_state = State {
                    eroom: *new_room,
                    open: state.open.clone(),
                    ..state
                };
                estates.push(new_state);
            }
        }

        for state in estates {
            let room = &rooms[&state.room];
            if room.flow_rate > 0 && !state.open.contains(&state.room) {
                // consider opening this valve
                let mut open = state.open.clone();
                open.push(state.room);
                open.sort();
                let new_state = State {
                    turn: state.turn+1,
                    released: state.released + orig_flow_rate,
                    flow_rate: state.flow_rate + room.flow_rate,
                    open,
                    ..state
                };
                let max_potential = new_state.max_potential(total_valve_pressure, max_turns);
                sol = sol.max(new_state.released);
                if max_potential > sol && !visited_states.contains(&state){
                    queue.push_increase(new_state, max_potential);
                }
            }
            for new_room in &room.tunnels {
                // consider moving to a new room
                let new_state = State {
                    turn: state.turn+1,
                    released: state.released + orig_flow_rate,
                    room: *new_room,
                    open: state.open.clone(),
                    ..state
                };
                let max_potential = new_state.max_potential(total_valve_pressure, max_turns);
                sol = sol.max(new_state.released);
                if max_potential > sol && !visited_states.contains(&state) {
                    queue.push_increase(new_state, max_potential);
                }
            }
        }
    }
    sol
}

pub fn solve() -> SolutionPair {
    // let input = get_input("input/16-test.txt");
    // let sol1 = solve_for1(&input);
    // let sol2 = solve_for2(&input, sol1);

    let input = get_input("input/16.txt");
    let sol1 = solve_for1(&input);
    // let sol1 = 0;
    let sol2 = solve_for2(&input, sol1);

    (Solution::from(sol1 as u32), Solution::from(sol2 as u32))
}
