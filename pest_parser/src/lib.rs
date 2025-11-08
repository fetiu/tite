use pest::iterators::Pairs;
use pest::Parser;

#[derive(Parser)]
#[grammar = "tite.pest"]
pub struct TiteParser;

pub fn parse_program(source: &str) -> Result<Pairs<'_, Rule>, pest::error::Error<Rule>> {
    TiteParser::parse(Rule::program, source)
}

#[cfg(test)]
mod tests {
    use super::{parse_program, Rule};

    #[test]
    fn parses_example_program() {
        let source = include_str!("../../src/example.tite");
        let pairs = parse_program(source).expect("example program should parse");
        let mut iter = pairs.clone();
        let program = iter.next().expect("program rule");
        assert_eq!(program.as_rule(), Rule::program);
        assert!(program.into_inner().any(|pair| pair.as_rule() == Rule::sequence));
    }
}

