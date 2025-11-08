use std::env;
use std::fs;
use std::process;

use pest::iterators::Pair;

use tite_parser::{parse_program, Rule};

fn main() {
    let path = env::args().nth(1).unwrap_or_else(|| {
        eprintln!("usage: print_tree <file>");
        process::exit(1);
    });

    let source = match fs::read_to_string(&path) {
        Ok(src) => src,
        Err(err) => {
            eprintln!("failed to read {path}: {err}");
            process::exit(1);
        }
    };

    let pairs = match parse_program(&source) {
        Ok(pairs) => pairs,
        Err(err) => {
            eprintln!("{err}");
            process::exit(2);
        }
    };

    for pair in pairs {
        print_pair(&pair, 0);
    }
}

fn print_pair(pair: &Pair<'_, Rule>, depth: usize) {
    let indent = "  ".repeat(depth);
    let span_text = pair.as_str().lines().next().unwrap_or("").trim();
    if span_text.is_empty() {
        println!("{indent}{:?}", pair.as_rule());
    } else {
        println!("{indent}{:?}: {}", pair.as_rule(), span_text);
    }

    for inner in pair.clone().into_inner() {
        print_pair(&inner, depth + 1);
    }
}

