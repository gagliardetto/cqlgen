package main

import (
	"os"

	. "github.com/gagliardetto/cqlgen/jen"
)

// https://github.com/github/codeql-go/blob/main/ql/examples/snippets/calltobuiltin.ql
func main() {
	file := NewFile()
	file.HeaderDoc("@name Call to built-in function")
	file.HeaderDoc("@description Finds calls to the built-in `len` function.")
	file.HeaderDoc("@id go/examples/calltolen")
	file.HeaderDoc("@tags call")
	file.HeaderDoc("      function")
	file.HeaderDoc("      len")
	file.HeaderDoc("      built-in")

	file.Import("go")

	file.From(
		Qual("DataFlow", "CallNode").Id("call"),
	)

	file.Where(DoGroup(func(gr *Group) {
		gr.Id("call").Eq().Qual("Builtin", "len").Call().Dot("getACall").Call()
	}))
	file.Select(Id("call"))

	file.Render(os.Stdout)
}
