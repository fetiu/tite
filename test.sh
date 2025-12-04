#!/bin/bash

if ! which antlr4-parse > /dev/null; then
    pip install antlr4-tools
fi

cd src
antlr4-parse parser/tyfn.g4 program -gui "example.ty"