use std::{fs::File, io::Read};

fn main() {
    let buf = {
        let mut buf = String::new();
        File::open("assets/1-1_input.txt")
            .unwrap()
            .read_to_string(&mut buf)
            .unwrap();
        buf
    };

    let pairs: Vec<(_, _)> = buf
        .trim()
        .split('\n')
        .filter_map(|line| line.split_once("   "))
        .collect();

    let (left, right): (Vec<_>, Vec<_>) = pairs.clone().into_iter().unzip();

    for (val1, val2) in pairs {
        println!("{val1}::{val2}")
    }

    println!("{}", left.len());
    println!("{}", right.len());
}
