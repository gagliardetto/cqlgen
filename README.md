# cqlgen
cqlgen is a code generator for CodeQL, based and converted from the awesome [Jennifer](https://github.com/dave/jennifer) Go code generation library.

## Examples

See https://github.com/gagliardetto/cqlgen/blob/main/examples/main.go for a general example.

You can find other examples in [/examples/other/go](/examples/other/go).

## Requirements

To allow cqlgen to format the generated codeql, you need a recent version of the [codeql cli](https://github.com/github/codeql-cli-binaries/releases) (otherwise it will not be formatted), and have it available as `codeql` in your PATH.

## Basic example

Here is an example:

```go
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
```

The above example will generate the below codeql code:

```codeql
/**
 * @name Call to built-in function
 * @description Finds calls to the built-in `len` function.
 * @id go/examples/calltolen
 * @tags call
 *       function
 *       len
 *       built-in
 */

import go

from DataFlow::CallNode call
where call = Builtin::len().getACall()
select call
```
