from __future__ import annotations

import llvmlite.ir as ir

import subprocess
from pathlib import Path


def generate_llvm_ir(parse_tree: str) -> str:
    module = ir.Module(name="tite_module")
    # Additional IR generation logic here
    return str(module)

def main():
    source_path = Path(__file__).resolve().with_name("example.tite")
    manifest_path = Path(__file__).resolve().parents[1] / "pest_parser" / "Cargo.toml"

    result = subprocess.run(
        [
            "cargo",
            "run",
            "--manifest-path",
            str(manifest_path),
            "--bin",
            "print_tree",
            "--quiet",
            str(source_path),
        ],
        check=False,
        capture_output=True,
        text=True,
    )

    if result.returncode != 0:
        raise RuntimeError(result.stderr.strip() or "failed to run Pest parser")

    llvm_ir = generate_llvm_ir(result.stdout)
    print(llvm_ir)

if __name__ == "__main__":
    main()
