use colored::Colorize;

use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////
#[allow(dead_code)]
fn print_forest(forest: &Vec<Vec<u8>>, visible_trees: &Vec<Vec<bool>>) {
    let x_max = forest[0].len();
    let y_max = forest.len();

    for y in 0..y_max {
        for x in 0..x_max {
            let tree = forest[y][x];
            let good = format!("{}", tree).green();
            let bad = format!("{}", tree);

            // let tree = "🌲";
            // let good = format!("{}", tree).green();
            // let bad = format!("{}", "  ").black();
            if visible_trees[y][x] {
                print!("{}", good);
            } else {
                print!("{}", bad);
            }
        }
        print!("\n");
    }
}

fn scenic_score(forest: &Vec<Vec<u8>>, x: usize, y: usize) -> usize {
    let x_max = forest[0].len();
    let y_max = forest.len();

    let my_height = forest[y][x];
    let mut score = 1;
    {
        // look up
        let mut cur_score: usize = 0;
        let mut cur_y = y;

        loop {
            if cur_y == 0 {
                break;
            }
            cur_y -= 1;
            cur_score += 1;
            if forest[cur_y][x] >= my_height {
                break;
            }
        }
        score *= cur_score;
    }
    {
        // look down
        let mut cur_score: usize = 0;
        let mut cur_y = y;

        loop {
            if cur_y == y_max-1 {
                break;
            }
            cur_y += 1;
            cur_score += 1;
            if forest[cur_y][x] >= my_height {
                break;
            }
        }
        score *= cur_score;
    }
    {
        // look left
        let mut cur_score: usize = 0;
        let mut cur_x = x;

        loop {
            if cur_x == 0 {
                break;
            }
            cur_x -= 1;
            cur_score += 1;
            if forest[y][cur_x] >= my_height {
                break;
            }
        }
        score *= cur_score;
    }
    {
        // look right
        let mut cur_score: usize = 0;
        let mut cur_x = x;

        loop {
            if cur_x == x_max-1 {
                break;
            }
            cur_x += 1;
            cur_score += 1;
            if forest[y][cur_x] >= my_height {
                break;
            }
        }
        score *= cur_score;
    }

    score
}

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/8.txt").lines();
    let mut forest: Vec<Vec<u8>> = Default::default();
    
    for line in input {
        let line = line.unwrap();
        let tree_row: Vec<u8> = line.trim().as_bytes().iter().map(|b| b - b'0').collect();
        forest.push(tree_row);
    }

    let x_max = forest[0].len();
    let y_max = forest.len();
    let mut visible_trees: Vec<Vec<bool>> = vec![vec![false; x_max]; y_max];

    for x in 0..x_max {
        // from top
        let start = 0;
        visible_trees[start][x] = true;
        let mut vis_height = forest[start][x];
        for y in 1..y_max-1 {
            let cur_tree_height = forest[y][x];
            if cur_tree_height > vis_height {
                visible_trees[y][x] = true;
                vis_height = cur_tree_height;
                if vis_height == 9 {
                    break;
                }
            }
        }

        // from bottom
        let start = y_max - 1;
        visible_trees[start][x] = true;
        let mut vis_height = forest[start][x];
        for y in (1..start).rev() {
            let cur_tree_height = forest[y][x];
            if cur_tree_height > vis_height {
                visible_trees[y][x] = true;
                vis_height = cur_tree_height;
                if vis_height == 9 {
                    break;
                }
            }
        }
    }

    for y in 0..y_max {
        let tree_row = &forest[y];
        let visible_row = &mut visible_trees[y];

        // from left
        let start = 0;
        visible_row[start] = true;
        let mut vis_height = tree_row[start];
        for x in 1..x_max-1 {
            let cur_tree_height = tree_row[x];
            if cur_tree_height > vis_height {
                visible_row[x] = true;
                vis_height = cur_tree_height;
                if vis_height == 9 {
                    break;
                }
            }
        }

        // from right
        let start = x_max - 1;
        visible_row[start] = true;
        let mut vis_height = tree_row[start];
        for x in (1..start).rev() {
            let cur_tree_height = tree_row[x];
            if cur_tree_height > vis_height {
                visible_row[x] = true;
                vis_height = cur_tree_height;
                if vis_height == 9 {
                    break;
                }
            }
        }
    }

    // print_forest(&forest, &visible_trees);

    let mut sol2  = 0;

    for y in 0..y_max {
        for x in 0..x_max {
            sol2 = sol2.max(scenic_score(&forest, x, y));
        }
    }

    let sol1: usize = visible_trees.iter().map(|row| row.iter().filter(|tree| **tree).count()).sum();

    (Solution::from(sol1), Solution::from(sol2))
}
