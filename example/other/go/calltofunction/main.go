package main

import (
	"os"

	. "github.com/gagliardetto/cqlgen/jen"
)

// https://github.com/github/codeql-go/blob/main/ql/examples/snippets/calltofunction.ql
func main() {
	file := NewFile()
	file.HeaderDoc("@name Call to library function")
	file.HeaderDoc(`@description Finds calls to "fmt.Println".`)
	file.HeaderDoc("@id go/examples/calltoprintln")
	file.HeaderDoc("@tags call")
	file.HeaderDoc("      function")
	file.HeaderDoc("      println")

	file.Import("go")

	file.From(
		Id("Function").Id("println"),
		Qual("DataFlow", "CallNode").Id("call"),
	)

	file.Where(DoGroup(func(gr *Group) {
		gr.Id("println").Dot("hasQualifiedName").Call(Lit("fmt"), Lit("Println"))

		gr.And()

		gr.Id("call").Eq().Id("println").Dot("getACall").Call()
	}))
	file.Select(Id("call"))

	file.Render(os.Stdout)
}
