use std::collections::HashMap;
use std::collections::HashSet;
use std::io::BufRead;

use crate::util;
use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

fn make_cur_path(parts: &Vec<String>) -> String {
    if parts.is_empty() {
        String::from("/")
    } else {
        String::from("/") + &parts.join("/") + "/"
    }
}

pub fn solve() -> SolutionPair {
    let input = util::open_input("input/7.txt").lines();
    let mut cur_directory_parts: Vec<String> = vec![];
    let mut files: HashMap<String, usize> = HashMap::new();
    let mut dirs: HashSet<String> = HashSet::new();

    for line in input {
        let line = line.unwrap();
        let cmd: Vec<&str> = line.trim().split(' ').collect();
        match cmd[..] {
            ["$", "ls"] => (), // do nothing
            ["dir", _] => (),  // do nothing
            [size, filename] => {
                let filename = make_cur_path(&cur_directory_parts) + filename;
                let size = size.parse::<usize>().unwrap();
                files.insert(filename, size);
            },
            ["$", "cd", "/"] => cur_directory_parts.clear(),
            ["$", "cd", ".."] => { cur_directory_parts.pop(); },
            ["$", "cd", newdir] => {
                let dirname = make_cur_path(&cur_directory_parts) + newdir + "/";
                dirs.insert(dirname);
    
                cur_directory_parts.push(String::from(newdir));
            },
            [..] => panic!("bad command: {:?}", cmd),
        }
    }

    let mut sol1: usize = 0;
    let total_space = 70_000_000;
    let used_space: usize = files.values().sum();
    let needed_space: usize = 30_000_000 - (total_space - used_space);
    let mut sol2: usize = total_space;
    
    for dir in dirs {
        let mut dirsize: usize = 0;
        for (file, size) in &files {
            if file.starts_with(&dir) {
                dirsize += size;
            }
        }
        if dirsize > 0 && dirsize <= 100_000 {
            sol1 += dirsize;
        }        
        if dirsize >= needed_space && dirsize < sol2 as usize {
            sol2 = dirsize;
        }        
    }

    (Solution::from(sol1), Solution::from(sol2))
}
