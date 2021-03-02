package main

import (
	"os"

	. "github.com/gagliardetto/cqlgen/jen"
)

// https://github.com/github/codeql-go/blob/main/ql/examples/snippets/calltomethod.ql
func main() {
	file := NewFile()
	file.HeaderDoc("@name Call to method")
	file.HeaderDoc("@description Finds calls to the `Get` method of type `Header` from the `net/http` package.")
	file.HeaderDoc("@id go/examples/calltoheaderget")
	file.HeaderDoc("@tags call")
	file.HeaderDoc("      function")
	file.HeaderDoc("      net/http")
	file.HeaderDoc("      Header")
	file.HeaderDoc("      strings")

	file.Import("go")

	file.From(
		Id("Method").Id("get"),
		Qual("DataFlow", "CallNode").Id("call"),
	)

	file.Where(DoGroup(func(gr *Group) {
		gr.Id("get").Dot("hasQualifiedName").Call(Lit("net/http"), Lit("Header"), Lit("Get"))

		gr.And()

		gr.Id("call").Eq().Id("get").Dot("getACall").Call()
	}))
	file.Select(Id("call"))

	file.Render(os.Stdout)
}
