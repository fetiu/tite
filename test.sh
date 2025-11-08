#!/bin/bash

set -euo pipefail

if ! command -v cargo > /dev/null; then
    echo "cargo is required to run the Pest parser demo" >&2
    exit 1
fi

cargo run --manifest-path pest_parser/Cargo.toml --bin print_tree --quiet src/example.tite

