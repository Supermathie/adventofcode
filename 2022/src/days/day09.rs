use colored::Colorize;

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

struct Rope {
    desc: char,
    head: Point,
    tail: Point,
    next: Option<Box<Rope>>,
    // on_tail_move: fn(Point),
}

fn print_state(rope: &Rope) {
    let mut grid: HashMap<Point, char> = HashMap::new();

    grid.insert(Point { x: 0, y: 0 }, 's');
    for d in 0..9 {
        let seg = rope.get(d);
        grid.insert(seg.head.clone(), seg.desc);
    }
    grid.insert(rope.get(9).tail.clone(), '9');

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

impl Rope {
    fn move_tail(&mut self) {
        let dx = self.head.x - self.tail.x;
        let dy = self.head.y - self.tail.y;
        if dx.abs() > 1 || dy.abs() > 1 {
            let dx = dx.clamp(-1, 1);
            let dy = dy.clamp(-1, 1);
            self.tail.x += dx;
            self.tail.y += dy;
            // print!("move_tail: seg {} moving {}, {}\n", self.desc, dx, dy);
            match self.next.as_deref_mut() {
                None => (),
                Some(n) => {
                    n.move_head(dx, dy);
                }
            }
            // (self.on_tail_move)(self.tail);
        }
    }
    fn move_head(&mut self, dx: CoordSize, dy: CoordSize) {
        // print!("move_head: seg {} moving {}, {}\n", self.desc, dx, dy);
        self.head.x += dx;
        self.head.y += dy;
        self.move_tail();
    }
    pub fn up(&mut self) {
        // print!("move_up\n");
        self.move_head(0, -1);
    }
    pub fn down(&mut self) {
        // print!("move_down\n");
        self.move_head(0, 1);
    }
    pub fn left(&mut self) {
        // print!("move_left\n");
        self.move_head(-1, 0);
    }
    pub fn right(&mut self) {
        // print!("move_right\n");
        self.move_head(1, 0);
    }

    pub fn get(&self, depth: u8) -> &Rope {
        let mut rope = self;
        for _ in 0..depth {
            match rope.next.as_deref() {
                None => panic!("end of the rope!"),
                Some(n) => rope = n,
            }
        }
        rope
    }        
}

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/9.txt").lines();

    let mut visited_locations1: HashSet<Point> = HashSet::new();
    let mut visited_locations2: HashSet<Point> = HashSet::new();

    // let cbk = |point: Point| visited_locations1.insert(point);
    let mut rope = Rope {
        desc: '9',
        head: Point { x: 0, y: 0 },
        tail: Point { x: 0, y: 0 },
        next: None,
        // on_tail_move: cbk,
    };
    visited_locations1.insert(Point { x: 0, y: 0 });
    visited_locations2.insert(Point { x: 0, y: 0 });

    for d in (0..8).rev() {
        rope = Rope {
            desc: char::from_digit(d, 10).unwrap(),
            head: Point { x: 0, y: 0 },
            tail: Point { x: 0, y: 0 },
            next: Some(Box::from(rope)),
            // on_tail_move: cbk,
        };
    }

    for line in input {
        let line = line.unwrap();
        let cmd: Vec<&str> = line.trim().split(' ').collect();
        let end_depth = 8;
        match cmd[..] {
            ["U", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope.up();
                    visited_locations1.insert(rope.get(0).tail.clone());
                    visited_locations2.insert(rope.get(end_depth).tail.clone());
                }
            }
            ["D", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope.down();
                    visited_locations1.insert(rope.get(0).tail.clone());
                    visited_locations2.insert(rope.get(end_depth).tail.clone());
                }
            }
            ["L", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope.left();
                    visited_locations1.insert(rope.get(0).tail.clone());
                    visited_locations2.insert(rope.get(end_depth).tail.clone());
                }
            }
            ["R", count] => {
                for _ in 0..count.parse().unwrap() {
                    rope.right();
                    visited_locations1.insert(rope.get(0).tail.clone());
                    visited_locations2.insert(rope.get(end_depth).tail.clone());
                }
            }
            [..] => panic!("bad command: {:?}", cmd),
        }
        // print!("{:?}\n", cmd);
        // print_state(&rope);

        // print_grid(&visited_locations2);
    }

    let sol1 = visited_locations1.len();
    let sol2 = visited_locations2.len();
    (Solution::from(sol1), Solution::from(sol2))
}
