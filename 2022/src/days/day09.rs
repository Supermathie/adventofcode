use std::collections::HashSet;
use std::collections::HashMap;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

type CoordSize = i16;
#[derive(Hash, Eq, PartialEq, Debug, Clone)]

struct Point {
    x: CoordSize,
    y: CoordSize,
}

struct RopeSegment {
    desc: char,
    loc: Point,
    next: Option<Box<RopeSegment>>,
}

#[allow(dead_code)]
fn print_state(rope_segment: &RopeSegment) {
    let mut grid: HashMap<Point, char> = HashMap::new();

    grid.insert(Point { x: 0, y: 0 }, 's');
    for d in 0..ROPE_LEN {
        let seg = rope_segment.get(d);
        grid.insert(seg.loc.clone(), seg.desc);
    }

    let x_min = grid.keys().map(|p| p.x).min().unwrap();
    let x_max = grid.keys().map(|p| p.x).max().unwrap();
    let y_min = grid.keys().map(|p| p.y).min().unwrap();
    let y_max = grid.keys().map(|p| p.y).max().unwrap();

    for y in y_min..=y_max {
        for x in x_min..=x_max {
            let p = Point { x: x, y: y };
            if grid.contains_key(&p) {
                print!("{}", grid[&p]);
            } else {
                print!("{}", '.');
            }
        }
        print!("\n");
    }
    print!("\n");
}

#[allow(dead_code)]
fn print_grid(grid: &HashSet<Point>) {
    let x_min = grid.iter().map(|p| p.x).min().unwrap();
    let x_max = grid.iter().map(|p| p.x).max().unwrap();
    let y_min = grid.iter().map(|p| p.y).min().unwrap();
    let y_max = grid.iter().map(|p| p.y).max().unwrap();

    for y in y_min..=y_max {
        for x in x_min..=x_max {
            let p = Point { x: x, y: y };
            if grid.contains(&p) {
                print!("{}", '#');
            } else {
                print!("{}", '.');
            }
        }
        print!("\n");
    }
    print!("\n");
}

impl RopeSegment {
    fn move_next(&mut self) {
        if let Some(next) = self.next.as_deref_mut() {
            let dx = self.loc.x - next.loc.x;
            let dy = self.loc.y - next.loc.y;

            if dx.abs() > 1 || dy.abs() > 1 {
                let dx = dx.clamp(-1, 1);
                let dy = dy.clamp(-1, 1);
                next.do_move(dx, dy);           
            }
        }
    }
    fn do_move(&mut self, dx: CoordSize, dy: CoordSize) {
        // print!("do_move: seg {} moving {}, {}\n", self.desc, dx, dy);
        self.loc.x += dx;
        self.loc.y += dy;
        self.move_next();
    }
    pub fn up(&mut self) {
        // print!("move_up\n");
        self.do_move(0, -1);
    }
    pub fn down(&mut self) {
        // print!("move_down\n");
        self.do_move(0, 1);
    }
    pub fn left(&mut self) {
        // print!("move_left\n");
        self.do_move(-1, 0);
    }
    pub fn right(&mut self) {
        // print!("move_right\n");
        self.do_move(1, 0);
    }
    pub fn get(&self, depth: u8) -> &RopeSegment {
        let mut rope_segment = self;
        for _ in 0..depth {
            match rope_segment.next.as_deref() {
                None => panic!("I'm at the end of my rope_segment!"),
                Some(n) => rope_segment = n,
            }
        }
        rope_segment
    }        
    pub fn get_loc(&self, depth: u8) -> Point {
        self.get(depth).loc.clone()
    }
}

const ROPE_LEN: u8 = 10;

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/9.txt").lines();

    let mut visited_locations1: HashSet<Point> = HashSet::new();
    visited_locations1.insert(Point { x: 0, y: 0 });

    let mut visited_locations2: HashSet<Point> = HashSet::new();
    visited_locations2.insert(Point { x: 0, y: 0 });

    let mut rope_segment = RopeSegment {
        desc: char::from_digit(ROPE_LEN as u32, 36).unwrap(),
        loc: Point { x: 0, y: 0 },
        next: None,
    };

    for d in (0..ROPE_LEN).rev() {
        rope_segment = RopeSegment {
            desc: char::from_digit(d as u32, 36).unwrap(),
            loc: Point { x: 0, y: 0 },
            next: Some(Box::from(rope_segment)),
        };
    }

    for line in input {
        let line = line.unwrap();
        let cmd: Vec<&str> = line.trim().split(' ').collect();
        match cmd[..] {
            ["U", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope_segment.up();
                    visited_locations1.insert(rope_segment.get_loc(1));
                    visited_locations2.insert(rope_segment.get_loc(ROPE_LEN - 1));
                }
            }
            ["D", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope_segment.down();
                    visited_locations1.insert(rope_segment.get_loc(1));
                    visited_locations2.insert(rope_segment.get_loc(ROPE_LEN - 1));
                }
            }
            ["L", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope_segment.left();
                    visited_locations1.insert(rope_segment.get_loc(1));
                    visited_locations2.insert(rope_segment.get_loc(ROPE_LEN - 1));
                }
            }
            ["R", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope_segment.right();
                    visited_locations1.insert(rope_segment.get_loc(1));
                    visited_locations2.insert(rope_segment.get_loc(ROPE_LEN - 1));
                }
            }
            [..] => panic!("bad command: {:?}", cmd),
        }
        // print!("{:?}\n", cmd);
        // print_state(&rope_segment);

        // print_grid(&visited_locations1);
        // print_grid(&visited_locations2);
    }

    let sol1 = visited_locations1.len();
    let sol2 = visited_locations2.len();
    (Solution::from(sol1), Solution::from(sol2))
}
