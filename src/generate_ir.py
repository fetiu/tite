import llvmlite.ir as ir
import llvmlite.binding as llvm
from antlr4 import FileStream, CommonTokenStream
from parser.TiteLexer import TiteLexer
from parser.TiteParser import TiteParser

def generate_llvm_ir(tree):
    module = ir.Module(name="tite_module")
    # Additional IR generation logic here
    return str(module)

def main():
    input_stream = FileStream("example.tite")
    lexer = TiteLexer(input_stream)
    stream = CommonTokenStream(lexer)
    parser = TiteParser(stream)
    tree = parser.program()
    llvm_ir = generate_llvm_ir(tree)
    print(llvm_ir)

if __name__ == "__main__":
    main()
