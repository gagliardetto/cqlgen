package main

import (
	"os"

	. "github.com/gagliardetto/cqlgen/jen"
)

// https://github.com/github/codeql-go/blob/main/ql/examples/snippets/updateinloop.ql
func main() {
	file := NewFile()
	file.HeaderDoc("@name Increment statements in loops")
	file.HeaderDoc("@description Finds increment statements that are nested in a loop")
	file.HeaderDoc("@id go/examples/updateinloop")
	file.HeaderDoc("@tags nesting")
	file.HeaderDoc("      increment")

	file.Import("go")

	file.From(
		Id("IncStmt").Id("s"),
		Id("LoopStmt").Id("l"),
	)

	file.Where().Id("s").Dot("getParent").Op("+").Call().Eq().Id("l")
	file.Select(Id("s"), Id("l"))

	file.Render(os.Stdout)
}
