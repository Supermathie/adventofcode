use sscanf;
use std::collections::HashSet;
use std::io::BufRead;
use std::fmt;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

type CoordSize = i64;

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
    fn manhattan(&self, other: &Self) -> u64 {
        (self.x.checked_sub(other.x).unwrap()).abs() as u64 + (self.y.checked_sub(other.y).unwrap()).abs() as u64
    }

}

#[derive(Hash, Eq, PartialEq, Debug, Clone)]
struct Sensor {
    loc: Point,
    beacon: Point,
    empty_dist: u64,
}

fn get_input(filename: &str) -> Vec<Sensor> {
    let input = util::open_input(filename).lines();

    let mut sensors: Vec<Sensor> = Default::default();

    for line in input {
        let line = line.unwrap();
        let (x, y, bx, by) = sscanf::sscanf!(line, "Sensor at x={CoordSize}, y={CoordSize}: closest beacon is at x={CoordSize}, y={CoordSize}").unwrap();
        let loc = Point { x, y };
        let beacon = Point { x: bx, y: by};
        let empty_dist = loc.manhattan(&beacon);
        sensors.push(Sensor { loc, beacon, empty_dist });
    }
    sensors
}

fn solve_for1(input: &Vec<Sensor>, y: i64) -> u64 {
    let beacons: HashSet<&Point> = input.iter().map(|s| &s.beacon).collect();
    let min_empty = input.iter().map(|s| s.loc.x.checked_sub(s.empty_dist.try_into().unwrap()).unwrap()).reduce(|a, b|a.min(b)).unwrap();
    let max_empty = input.iter().map(|s| s.loc.x.checked_add(s.empty_dist.try_into().unwrap()).unwrap()).reduce(|a, b|a.max(b)).unwrap();
    // println!("{}, {}", min_empty, max_empty);
    let mut sol1: u64 = 0;
    'x_loop: for x in min_empty..=max_empty {
        let trial_point = Point { x, y };
        if beacons.contains(&trial_point) {
            // print!("B");
            continue;
        }
        for s in input {
            if s.loc.manhattan(&trial_point) <= s.empty_dist {
                sol1 += 1;
                // print!("#");
                continue 'x_loop;
            }
        }
        // print!(".");
    }
    // print!("\n");
    sol1
}

fn solve_for2(input: &Vec<Sensor>, _x_min: i64, x_max: i64, _y_min: i64, y_max: i64) -> i64 {
    let mut x = 0;
    let mut y = 0;
    while y <= y_max {
        'x_loop: while x <= x_max {
            let trial_point = Point { x, y };
            // println!("testing: {}", trial_point);
            for s in input {
                let dist = s.loc.manhattan(&trial_point);
                if dist <= s.empty_dist {
                    // trial is closer to beacon
                    // println!("close to beacon: {}", s.loc);
                    x = x.max(s.loc.x) + (s.empty_dist - dist) as i64 + 1;
                    continue 'x_loop;
                }
            }
            return 4000000 * x + y
        }
        // hit the end of the line
        x = 0;
        y += 1;
    }
    panic!("no sensor locations found")
}

pub fn solve() -> SolutionPair {
    // let input = get_input("input/15-test.txt");
    // let sol1 = solve_for1(&input, 10);
    // let sol2 = solve_for2(&input, 0, 20, 0, 20);

    let input = get_input("input/15.txt");
    let sol1 = solve_for1(&input, 2000000);
    let sol2 = solve_for2(&input, 0, 4000000, 0, 4000000);

    (Solution::from(sol1), Solution::from(sol2))
}
