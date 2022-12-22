use std::fmt;
use std::fmt::Write;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

#[derive(Default)]
struct Grid([Vec<bool>; 7]);

impl fmt::Display for Grid {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for y in (0..self.max_height()).rev() {
            f.write_char('|')?;
            for x in 0..7 {
                if *self.get(x, y) {
                    f.write_char('#')?
                } else {
                    f.write_char('.')?
                }
            }
            f.write_char('|')?;
            f.write_char('\n')?;
        }
        f.write_str("+-------+\n")?;
        Ok(())
    }
}

impl Grid {
    fn max_height(&self) -> usize {
        self.0.iter().map(|v| v.len()).reduce(|a, b| a.max(b)).unwrap()
    }
    fn min_height(&self) -> usize {
        self.0.iter().map(|v| v.len()).reduce(|a, b| a.min(b)).unwrap()
    }
    fn get(&self, x: usize, y: usize) -> &bool {
        self.0[x].get(y).or(Some(&false)).unwrap()
    }
    fn piece_valid_at(&self, piece: &Piece, loc: &Point) -> bool {
        piece.shape().iter().all(|point| loc.x + point.x < 7 && !self.get(loc.x + point.x, loc.y + point.y))
    }
    fn add_piece_at(&mut self, piece: &Piece, loc: &Point) {
        for point in piece.shape() {
            let x = point.x + loc.x;
            let y = point.y + loc.y;
            let col = &mut self.0[x];
            if col.len() <= y {
                col.resize(y+1, false);
            }
            col[y] = true;
        }
    }
    fn skyline(&self) -> usize {
        let min_height = self.min_height();
        let mut skyline = 0;
        for h in self.0.iter().map(|v| v.len()) {
            skyline *= 100;
            skyline += h - min_height;
        }
        skyline
    }
}

struct Point {
    x: usize,
    y: usize,
}

enum Piece {
    A, // _
    B, // +
    C, // backwards L
    D, // |
    E, // square
}

impl Piece {
    fn shape(&self) -> &'static [Point] {
        match self {
            Self::A => &[ Point { x: 0, y: 0}, Point { x: 1, y: 0}, Point { x: 2, y: 0}, Point { x: 3, y: 0}, ],
            Self::B => &[ Point { x: 0, y: 1}, Point { x: 1, y: 0}, Point { x: 1, y: 1}, Point { x: 1, y: 2}, Point { x: 2, y: 1}, ],
            Self::C => &[ Point { x: 0, y: 0}, Point { x: 1, y: 0}, Point { x: 2, y: 0}, Point { x: 2, y: 1}, Point { x: 2, y: 2}, ],
            Self::D => &[ Point { x: 0, y: 0}, Point { x: 0, y: 1}, Point { x: 0, y: 2}, Point { x: 0, y: 3}, ],
            Self::E => &[ Point { x: 0, y: 0}, Point { x: 1, y: 0}, Point { x: 0, y: 1}, Point { x: 1, y: 1}, ],
        }
    }
}

fn get_input(filename: &str) -> String {
    let mut buf = String::new();
    util::open_input(filename).read_line(&mut buf).expect("read error");
    buf
}

fn drop_pieces<'a>(grid: &mut Grid, jetstream: &mut impl Iterator<Item = char>, pieces: &mut impl Iterator<Item = &'a Piece>, max_pieces: usize) {
    let mut num_pieces:usize  = 0;
    while num_pieces < max_pieces {
        let piece = pieces.next().expect("impossible!");
        let mut pos = Point { x: 2, y: grid.max_height() + 3};

        loop {
            match jetstream.next() { // jetstream fires
                Some('<') => {
                    if pos.x > 0 {
                        if grid.piece_valid_at(piece, &Point { x: pos.x - 1, ..pos }) {
                            pos.x -= 1;
                        }
                    }
                }
                Some('>') => {
                    if grid.piece_valid_at(piece, &Point { x: pos.x + 1, ..pos }) {
                        pos.x += 1;
                    }
                }
                _ => panic!("Impossible!")
            }
            // piece tries to drop
            if pos.y > 0 && grid.piece_valid_at(piece, &Point { y: pos.y - 1, ..pos}) {
                pos.y -= 1;
            } else {
                // landed!
                num_pieces += 1;
                grid.add_piece_at(piece, &pos);
                break
            }
        }
    }

}

fn solve_for1(input: &String, max_pieces: usize) -> usize {
    let mut jetstream = input.trim().chars().cycle();
    let mut pieces = [Piece::A, Piece::B, Piece::C, Piece::D, Piece::E].iter().cycle();

    let mut grid: Grid = Default::default();
    drop_pieces(&mut grid, &mut jetstream, &mut pieces, max_pieces);

    grid.max_height()
}

fn solve_for2(input: &String, max_pieces: usize) -> usize {
    let mut jetstream = input.trim().chars().cycle();
    let mut pieces = [Piece::A, Piece::B, Piece::C, Piece::D, Piece::E].iter().cycle();

    let mut grid: Grid = Default::default();

    let cycle_size = input.trim().chars().count() * 5 * 171 * 2;

    let init_num_pieces = max_pieces % cycle_size;
    let mut num_cycles_remaining = max_pieces / cycle_size;

    drop_pieces(&mut grid, &mut jetstream, &mut pieces, init_num_pieces);

    let mut prev_increase = 0;

    println!("cycle size: {}, init_num: {}", cycle_size, init_num_pieces);

    while num_cycles_remaining > 0 {
        // println!("cycles remaining: {}, skyline: {}", num_cycles_remaining, grid.skyline());
        let prev_height = grid.max_height();
        let prev_skyline = grid.skyline();
        drop_pieces(&mut grid, &mut jetstream, &mut pieces, cycle_size);
        num_cycles_remaining -= 1;

        let cur_increase = grid.max_height() - prev_height;
        let cur_skyline = grid.skyline();
        println!("prev_increase: {}, cur_increase: {}, skyline: {}", prev_increase, cur_increase, grid.skyline());
        if prev_increase == cur_increase && cur_skyline == prev_skyline {
            return grid.max_height() + num_cycles_remaining * cur_increase
        } else {
            prev_increase = cur_increase;
        }
    }
    grid.max_height()
}


pub fn solve() -> SolutionPair {
    let input = get_input("input/17.txt");

    let sol1 = solve_for1(&input, 2022);
    let sol2 = solve_for2(&input, 1000000000000);

    (Solution::from(sol1), Solution::from(sol2))
}
