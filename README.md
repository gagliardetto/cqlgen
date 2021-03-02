# cqlgen
cqlgen is a code generator for CodeQL, based and converted from the awesome [Jennifer](https://github.com/dave/jennifer) Go code generation library.

## Examples

See https://github.com/gagliardetto/cqlgen/blob/main/example/main.go for a general example.

You can find other examples in [/example/other/go](/example/other/go).

## Requirements

To be able to format the generated codeql, you need a recent version of the [codeql cli](https://github.com/github/codeql-cli-binaries/releases) (otherwise it will not be formatted), and have it available as `codeql` in your PATH.