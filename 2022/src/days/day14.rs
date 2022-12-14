use std::collections::HashSet;
use std::collections::HashMap;
use std::fmt;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

type CoordSize = u32;

#[derive(Hash, Eq, PartialEq, Debug, Clone)]
struct Point {
    x: CoordSize,
    y: CoordSize,
}

impl fmt::Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({},{})", self.x, self.y)
    }
}
impl Point {
    fn from_str(s: &str) -> Point {
        let (x, y) = s.split_once(',').unwrap();
        Point { x: x.parse().unwrap(), y: y.parse().unwrap() }
    }
}

fn fall_from1(map: &HashSet<Point>, start: &Point, y_max: &CoordSize) -> Option<Point> {
    let mut coord = start.clone();
    loop {
        if coord.y > *y_max {
            return None
        }
        if map.contains(&Point { y: coord.y + 1, ..coord }) {
            // blocked from falling down
            if map.contains(&Point { x: coord.x - 1, y: coord.y + 1 }) {
                // blocked from falling left
                if map.contains(&Point { x: coord.x + 1, y: coord.y + 1 }) {
                    // blocked from falling right
                    return Some(coord)
                } else {
                    coord.x += 1;
                    coord.y += 1;
                }
            } else {
                coord.x -= 1;
                coord.y += 1;
            }
        } else {
            coord.y += 1;
        }
    }
}

fn fall_from2(map: &HashSet<Point>, start: &Point, y_max: &CoordSize) -> Point {
    let mut coord = start.clone();
    loop {
        if coord.y == *y_max + 1 {
            return coord
        }
        if map.contains(&Point { y: coord.y + 1, ..coord }) {
            // blocked from falling down
            if map.contains(&Point { x: coord.x - 1, y: coord.y + 1 }) {
                // blocked from falling left
                if map.contains(&Point { x: coord.x + 1, y: coord.y + 1 }) {
                    // blocked from falling right
                    return coord
                } else {
                    coord.x += 1;
                    coord.y += 1;
                }
            } else {
                coord.x -= 1;
                coord.y += 1;
            }
        } else {
            coord.y += 1;
        }
    }
}

fn solve_for1(input: &HashSet<Point>) -> usize {
    let mut map = input.clone();
    let y_max = map.iter().map(|p| p.y).reduce(|a, b| a.max(b)).unwrap();
    let start = Point { x: 500, y: 0 };

    while let Some(new_grain) = fall_from1(&map, &start, &y_max) {
        map.insert(new_grain);
    }

    map.len() - input.len()
}

fn solve_for2(input: &HashSet<Point>) -> usize {
    let mut map = input.clone();
    let y_max = map.iter().map(|p| p.y).reduce(|a, b| a.max(b)).unwrap();
    let start = Point { x: 500, y: 0 };

    loop {
        let new_grain = fall_from2(&map, &start, &y_max);
        if new_grain == start {
            break;
        }
        map.insert(new_grain);
    }
    map.len() - input.len() + 1
}

fn get_input(filename: &str) -> HashSet<Point> {
    let input = util::open_input(filename).lines();

    let mut map: HashSet<Point> = Default::default();

    for line in input {
        let line = line.unwrap();
        let line: Vec<&str> = line.split(" -> ").collect();
        for win in line.windows(2) {
            let start = Point::from_str(win[0]);
            let end = Point::from_str(win[1]);
            for x in (start.x..=end.x).chain(end.x..=start.x) {
                for y in (start.y..=end.y).chain(end.y..=start.y) {
                    map.insert(Point { x: x, y: y });
                }
            }
        }
    }
    map
}


pub fn solve() -> SolutionPair {
    let mut input = get_input("input/14.txt");

    let sol1 = solve_for1(&input);
    let sol2 = solve_for2(&input);


    (Solution::from(sol1), Solution::from(sol2))
}
