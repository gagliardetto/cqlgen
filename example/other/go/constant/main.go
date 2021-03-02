package main

import (
	"os"

	. "github.com/gagliardetto/cqlgen/jen"
)

// https://github.com/github/codeql-go/blob/main/ql/examples/snippets/constant.ql
func main() {
	file := NewFile()
	file.HeaderDoc("@name Compile-time constant")
	file.HeaderDoc("@description Finds compile-time constants with value zero.")
	file.HeaderDoc("@id go/examples/zeroconstant")
	file.HeaderDoc("@tags expression")
	file.HeaderDoc("      numeric value")
	file.HeaderDoc("      constant")

	file.Import("go")

	file.From(
		Qual("DataFlow", "Node").Id("zero"),
	)

	file.Where().Id("zero").Dot("getNumericValue").Call().Eq().Lit(0)
	file.Select(Id("zero"))

	file.Render(os.Stdout)
}
