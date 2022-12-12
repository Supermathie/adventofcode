use std::collections::VecDeque;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

type CoordSize = i8;
#[derive(Hash, Eq, PartialEq, Debug, Clone, Copy)]

struct Point {
    x: CoordSize,
    y: CoordSize,
}

fn can_traverse_forwards(from: &u8, to: &u8) -> bool {
    *to <= from + 1
}

fn can_traverse_backwards(from: &u8, to: &u8) -> bool {
    *from <= to + 1
}

fn get_input(filename: &str) -> (Vec<Vec<u8>>, Point, Point) {
    let input = util::open_input(filename).lines();

    let mut heightmap: Vec<Vec<u8>> = Default::default();
    let mut start_pos = Point {x: 0, y: 0};
    let mut end_pos = Point {x: 0, y: 0};

    for line in input {
        let mut line = line.unwrap();
        match line.find('S') {
            None => (),
            Some(pos) => {
                start_pos.x = pos as i8;
                start_pos.y = heightmap.len() as i8;
                line = line.replace('S', "a");
            }
        }
        match line.find('E') {
            None => (),
            Some(pos) => {
                end_pos.x = pos as i8;
                end_pos.y = heightmap.len() as i8;
                line = line.replace('E', "z");
            }
        }
        let row: Vec<u8> = line.trim().as_bytes().iter().map(|b| b - b'a').collect();
        heightmap.push(row);
    }
    (heightmap, start_pos, end_pos)
}

fn solve_for(heightmap: &Vec<Vec<u8>>, start_pos: Point, end_condition: impl Fn(Point, u8) -> bool, can_traverse: fn(&u8, &u8) -> bool) -> u32 {
    let x_max = heightmap[0].len();
    let y_max = heightmap.len();

    let mut cost: Vec<Vec<Option<u32>>> = vec![vec![None; x_max]; y_max];
    cost[start_pos.y as usize][start_pos.x as usize] = Some(0);

    let mut queue: VecDeque<Point> = Default::default();
    queue.push_front(start_pos);

    let end_pos: Point = loop {
        let pos = queue.pop_back().expect("queue ran dry!");
        let pos_height = heightmap[pos.y as usize][pos.x as usize];
        let pos_cost = cost[pos.y as usize][pos.x as usize].unwrap();

        if end_condition(pos, pos_height) {
            break pos;
        }

        let new = Point { x: pos.x, y: pos.y - 1 }; // up
        if let Some(new_height) = heightmap.get(new.y as usize).and_then(|row| row.get(new.x as usize)) {
            // ok it's on the map
            if can_traverse(&pos_height, new_height) && cost[new.y as usize][new.x as usize] == None {
                cost[new.y as usize][new.x as usize] = Some(pos_cost + 1);
                queue.push_front(new);
            }
        }
        let new = Point { x: pos.x, y: pos.y + 1 }; // down
        if let Some(new_height) = heightmap.get(new.y as usize).and_then(|row| row.get(new.x as usize)) {
            // ok it's on the map
            if can_traverse(&pos_height, new_height) && cost[new.y as usize][new.x as usize] == None {
                cost[new.y as usize][new.x as usize] = Some(pos_cost + 1);
                queue.push_front(new);
            }
        }
        let new = Point { x: pos.x - 1, y: pos.y }; // left
        if let Some(new_height) = heightmap.get(new.y as usize).and_then(|row| row.get(new.x as usize)) {
            // ok it's on the map
            if can_traverse(&pos_height, new_height) && cost[new.y as usize][new.x as usize] == None {
                cost[new.y as usize][new.x as usize] = Some(pos_cost + 1);
                queue.push_front(new);
            }
        }
        let new = Point { x: pos.x + 1, y: pos.y }; // right
        if let Some(new_height) = heightmap.get(new.y as usize).and_then(|row| row.get(new.x as usize)) {
            // ok it's on the map
            if can_traverse(&pos_height, new_height) && cost[new.y as usize][new.x as usize] == None {
                cost[new.y as usize][new.x as usize] = Some(pos_cost + 1);
                queue.push_front(new);
            }
        }
    };
    cost[end_pos.y as usize][end_pos.x as usize].unwrap()
}

pub fn solve() -> SolutionPair {

    let (heightmap, start_pos, end_pos) = get_input("input/12.txt");

    let end_clone = end_pos.clone();
    let sol1 = solve_for(&heightmap, start_pos, |p, _| p == end_clone, can_traverse_forwards);
    let sol2 = solve_for(&heightmap, end_pos, |_, h| h == 0, can_traverse_backwards);

    (Solution::from(sol1), Solution::from(sol2))
}
