package main

import (
	"os"

	. "github.com/gagliardetto/cqlgen/jen"
)

// https://github.com/github/codeql-go/blob/main/ql/examples/snippets/zerocheck.ql
func main() {
	file := NewFile()
	file.HeaderDoc("@name Comparison with zero")
	file.HeaderDoc("@description Finds comparisons between an unsigned value and zero.")
	file.HeaderDoc("@id go/examples/unsignedgez")
	file.HeaderDoc("@tags comparison")
	file.HeaderDoc("      unsigned")

	file.Import("go")

	file.From(
		Qual("DataFlow", "RelationalComparisonNode").Id("cmp"),
		Qual("DataFlow", "Node").Id("unsigned"),
		Qual("DataFlow", "Node").Id("zero"),
	)

	file.Where(DoGroup(func(gr *Group) {
		gr.Id("zero").Dot("getNumericValue").Call().Eq().Lit(0)

		gr.And()

		gr.Id("unsigned").Dot("getType").Call().Dot("getUnderlyingType").Call().Instanceof().Id("UnsignedIntegerType")

		gr.And()

		gr.Id("cmp").Dot("leq").Call(DontCare(), Id("zero"), Id("unsigned"), Lit(0))
	}))
	file.Select(Id("cmp"), Id("unsigned"))

	file.Render(os.Stdout)
}
