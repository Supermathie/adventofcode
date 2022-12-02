use std::env;
use std::io::BufRead;
use std::cmp;

mod util;

fn main() {
    let args: Vec<String> = env::args().collect();
    let year: u32 = args[1].parse().expect("need an uint year");
    let day:  u32 = args[2].parse().expect("need an uint day");
    let input = util::open_input("input/1.txt").lines();

    let mut max:u32 = 0;
    let mut cur:u32 = 0;
    let mut totals: Vec<u32> = Vec::new();

    for line in input {
        let val = line.expect("uhoh: ");
        if val == "" {
            max = cmp::max(cur, max);
            totals.push(cur);
            cur = 0;
        } else {
            cur += val.parse::<u32>().expect("this should be a number!")
        }
    }
    print!("{}\n", max);
    totals.sort();
    // print!("{}\n", totals[totals.len()-1]);
    let top3 = totals.pop().expect("not enough elves") + totals.pop().expect("not enough elves") + totals.pop().expect("not enough elves");
    print!("{}\n", top3);
}
