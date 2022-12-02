use std::fs::{self, File};
use std::io::{self, BufReader};

pub fn open_input(filename: &str) -> BufReader<File> {
    io::BufReader::new(fs::File::open(&filename).expect("cannot open file: "))
}