use sscanf;
use std::collections::HashSet;
use std::collections::HashMap;
use std::io::BufRead;
use std::fmt;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

type CoordSize = i8;

#[derive(Hash, Eq, PartialEq, Debug, Clone)]
struct Point {
    x: CoordSize,
    y: CoordSize,
    z: CoordSize,
}

struct PointNeighbours<'a> {
    point: &'a Point,
    count: u8,
}

impl Iterator for PointNeighbours<'_> {
    type Item = Point;

    fn next(&mut self) -> Option<Self::Item> {
        self.count += 1;
        match self.count {
            1 => Some(Point { x: self.point.x - 1, ..*self.point }),
            2 => Some(Point { x: self.point.x + 1, ..*self.point }),
            3 => Some(Point { y: self.point.y - 1, ..*self.point }),
            4 => Some(Point { y: self.point.y + 1, ..*self.point }),
            5 => Some(Point { z: self.point.z - 1, ..*self.point }),
            6 => Some(Point { z: self.point.z + 1, ..*self.point }),
            _ => None,
        }
    }
}

impl Point {
    fn neighbours(&self) -> PointNeighbours {
        PointNeighbours { point: &self, count: 0 }
    }
}

impl fmt::Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({},{},{})", self.x, self.y, self.z)
    }
}

fn get_input(filename: &str) -> Vec<Point> {
    util::open_input(filename).lines().map(|line| {
        let line = line.unwrap();
        let (x,y,z) = sscanf::sscanf!(line, "{CoordSize},{CoordSize},{CoordSize}").unwrap();
        Point {x, y, z}
    }).collect()
}

fn build_map(input: &Vec<Point>) -> HashMap<Point, u8> {
    let mut grid: HashMap<Point, u8> = Default::default();
    for point in input {
        let mut val = 6;
        for n in point.neighbours() {
            if grid.contains_key(&n) {
                grid.entry(n).and_modify(|v| *v -= 1);
                val -= 1;
            }
        }
        grid.insert(point.clone(), val);
    }

    grid
}

fn solve_for1(grid: &HashMap<Point, u8>) -> u64 {
    grid.values().map(|v| *v as u64).reduce(|a, b| a + b ).unwrap()
}

fn solve_for2(grid: &HashMap<Point, u8>) -> u64 {
    let total_surfacearea = grid.values().map(|v| *v as u64).reduce(|a, b| a + b ).unwrap();
    let x_range = grid.keys().map(|p| (p.x, p.x)).reduce(|a, b| (a.0.min(b.0), a.1.max(b.1))).unwrap();
    let y_range = grid.keys().map(|p| (p.y, p.y)).reduce(|a, b| (a.0.min(b.0), a.1.max(b.1))).unwrap();
    let z_range = grid.keys().map(|p| (p.z, p.z)).reduce(|a, b| (a.0.min(b.0), a.1.max(b.1))).unwrap();

    let x_range = (x_range.0-1)..=(x_range.1+1);
    let y_range = (y_range.0-1)..=(y_range.1+1);
    let z_range = (z_range.0-1)..=(z_range.1+1);

    let mut grid = grid.clone();
    let mut queue: Vec<Point> = vec![ Point{ x: *x_range.start(), y: *y_range.start(), z: *z_range.start() } ];
    let mut fill: HashSet<Point> = Default::default();
    while let Some(point) = queue.pop() {
        if fill.insert(point.clone()) {
            // not yet visited
            for n in point.neighbours() {
                if grid.contains_key(&n) {
                    grid.entry(n).and_modify(|v| *v -= 1);
                } else {
                    if x_range.contains(&n.x) && y_range.contains(&n.y) && z_range.contains(&n.z) {
                        queue.push(n)
                    }
                }
            }
        }
    }
        
    let interior_surfacearea = grid.values().map(|v| *v as u64).reduce(|a, b| a + b ).unwrap();
    total_surfacearea - interior_surfacearea
}


pub fn solve() -> SolutionPair {
    let input = get_input("input/18.txt");
    let grid = build_map(&input);
    let sol1 = solve_for1(&grid);
    let sol2 = solve_for2(&grid);

    (Solution::from(sol1), Solution::from(sol2))
}
