package main

import (
	"fmt"

	. "github.com/gagliardetto/cqlgen/jen"
	"github.com/gagliardetto/utilz"
)

func main() {
	f := NewFile()
	f.HeaderDoc("This is a doc.")
	f.HeaderDoc("Second line of doc.")
	f.Import("go")
	f.Import("DataFlow::PathGraph")

	f.Doc("Doc about this module.")
	f.Private().Module().Id("SomeFramework").BlockFunc(func(modGr *Group) {
		modGr.Doc("Doc about class")
		modGr.Private().Class().Id("UntrustedSource").Extends().List(Id("UntrustedFlowSource::Range")).
			BlockFunc(func(classGr *Group) {
				classGr.Id("UntrustedSource").Call().BlockFunc(func(metGr *Group) {
					metGr.Comment("Example block 1: the type is a function.")
					metGr.Comment("The source is either the result (one or more), or a parameter (one or more),")
					metGr.Comment("or a mix of them.")

					metGr.ParensFunc(func(par *Group) {
						par.Comment("Function: github.com/example/package.GetSomething")
						par.ExistsFunc(func(exists *Group) {
							exists.List(
								Id("Function").Id("fn"),
								Id("DataFlow::CallNode").Id("call"),
							)

							exists.Bar()

							exists.Id("fn").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("GetSomething"))

							exists.Bar()
							exists.Comment("The source is the result:")

							exists.Id("call").Eq().Id("fn").Dot("getACall").Call().
								And().
								This().Eq().Id("call").Dot("getResult").Call()

						})

						par.Or()

						par.Comment("Function: github.com/example/package.ParseSomething")
						par.ExistsFunc(func(exists *Group) {
							exists.List(
								Id("Function").Id("fn"),
								Id("DataFlow::CallNode").Id("call"),
							)

							exists.Bar()

							exists.Id("fn").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("ParseSomething"))

							exists.Bar()
							exists.Comment("The source is the 0th parameter:")

							exists.Id("call").Eq().Id("fn").Dot("getACall").Call().
								And().
								This().Eq().Id("FunctionOutput::parameter").Call(Lit(0)).Dot("getExitNode").Call(Id("call"))
						})
					})

					metGr.Or()

					metGr.Comment("Example block 2: the type is a struct.")
					metGr.Comment("The source can be a method call (results or parameters of a call),")
					metGr.Comment("or a field read.")
					metGr.ParensFunc(func(par *Group) {
						par.Comment("Struct: github.com/example/package.Context")
						par.ExistsFunc(func(exists *Group) {
							exists.List(
								String().Id("typeName"),
							)
							exists.Bar()
							exists.Id("typeName").Eq().Lit("Context")
							exists.Bar()

							exists.Comment("Method calls on `Context`:")
							exists.ExistsFunc(func(subExists *Group) {
								subExists.List(
									Id("DataFlow::MethodCallNode").Id("call"),
									String().Id("methodName"),
								)

								subExists.Bar()
								subExists.Id("call").Dot("getTarget").Call().Dot("hasQualifiedName").Call(
									Lit("github.com/example/package"),
									Id("typeName"),
									Id("methodName"),
								)
								subExists.And()
								subExists.ParensFunc(
									func(par *Group) {
										par.Id("methodName").Eq().Lit("FullPath")
										par.And()
										par.Parens(
											Comment("The source is the method call result #0:"),
											This().Eq().Id("call").Dot("getResult").Call(Lit(0)),
											Or(),
											Comment("The source is method call parameter #0:"),
											This().Eq().Id("FunctionOutput::parameter").Call(Lit(0)).Dot("getExitNode").Call(Id("call")),
										)

										par.Or()
										par.Comment("The source is any result of the call?")
										par.Id("methodName").Eq().Lit("GetHeader")
									},
								)

							})

							exists.Or()
							exists.Comment("Field reads on `Context`:")
							exists.ExistsFunc(func(subExists *Group) {
								subExists.List(
									Id("DataFlow::Field").Id("fld"),
									String().Id("fieldName"),
								)

								subExists.Bar()
								subExists.Id("fld").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Id("typeName"), Id("fieldName"))

								subExists.And()
								subExists.Comment("The source is one of these fields reads:")

								subExists.Id("fieldName").In().Set(Lit("Accepted"), Lit("Params"))
								subExists.And()
								subExists.This().Eq().Id("fld").Dot("getARead").Call()
							})
						})
						par.Or()

						par.Comment("Struct: github.com/example/package.SomeStruct")
						par.ExistsFunc(func(exists *Group) {
							exists.List(
								Id("DataFlow::Field").Id("fld"),
								String().Id("fieldName"),
							)

							exists.Bar()
							exists.Id("fld").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("SomeStruct"), Id("fieldName"))

							exists.And()
							exists.Comment("The source is one of these fields reads:")

							exists.Id("fieldName").In().Set(Lit("Hello"), Lit("World"))
							exists.And()
							exists.This().Eq().Id("fld").Dot("getARead").Call()

						})

					})

					metGr.Or()

					metGr.Comment("Example block 3: the type is some custom type.")
					metGr.Comment("The source is a method call on that type (results or parameters of a call).")
					metGr.Comment("Some type: github.com/example/package.Slice")
					metGr.ExistsFunc(func(exists *Group) {
						exists.String().Id("typeName")
						exists.Bar()
						exists.Id("typeName").Eq().Lit("Slice")
						exists.Bar()

						exists.Comment("Method calls:")

						exists.ExistsFunc(func(subExists *Group) {
							subExists.List(
								Id("DataFlow::MethodCallNode").Id("call"),
								String().Id("methodName"),
							)

							subExists.Bar()
							subExists.Id("call").Dot("getTarget").Call().Dot("hasQualifiedName").Call(
								Lit("github.com/example/package"),
								Id("typeName"),
								Id("methodName"),
							)
							subExists.And()
							subExists.ParensFunc(
								func(par *Group) {
									par.Id("methodName").Eq().Lit("GetHeader")
									par.And()
									par.Comment("The source is the method call result #0:")
									par.This().Eq().Id("call").Dot("getResult").Call(Lit(0))
									par.Or()
									par.Id("methodName").Eq().Lit("ParseHeader")
									par.And()
									par.Comment("The source is method call parameter #1:")
									par.This().Eq().Id("FunctionOutput::parameter").Call(Lit(1)).Dot("getExitNode").Call(Id("call"))

								},
							)

						})
					})

					metGr.Or()

					metGr.Comment("Example block 4: the type is an interface.")
					metGr.Comment("The source is a method call on that interface (results or parameters of a call).")
					metGr.Comment("Interface: github.com/example/package.SomeInterface")
					metGr.ExistsFunc(func(exists *Group) {
						exists.String().Id("typeName")
						exists.Bar()
						exists.Id("typeName").Eq().Lit("SomeInterface")
						exists.Bar()

						exists.Comment("Method calls:")

						exists.ExistsFunc(func(subExists *Group) {
							subExists.List(
								Id("DataFlow::MethodCallNode").Id("call"),
								String().Id("methodName"),
							)

							subExists.Bar()
							subExists.Id("call").Dot("getTarget").Call().Dot("implements").Call(
								Lit("github.com/example/package"),
								Id("typeName"),
								Id("methodName"),
							)
							subExists.And()
							subExists.ParensFunc(
								func(par *Group) {
									par.Id("methodName").Eq().Lit("GetSomething")
									par.And()
									par.Comment("The source is the method call result #0:")
									par.This().Eq().Id("call").Dot("getResult").Call(Lit(0))

									par.Or()

									par.Id("methodName").Eq().Lit("UnmarshalSomething")
									par.And()
									par.Comment("The source is method call parameter #2:")
									par.This().Eq().Id("FunctionOutput::parameter").Call(Lit(2)).Dot("getExitNode").Call(Id("call"))

								},
							)

						})
					})

					metGr.Or()

					metGr.Comment("Example block 5: the type is whatever.")
					metGr.Comment("The source is a read of a variable of that type.")
					metGr.Comment("Type: github.com/example/package.Slice")
					metGr.ExistsFunc(func(exists *Group) {
						exists.List(
							Id("DataFlow::ReadNode").Id("read"),
							Id("ValueEntity").Id("v"),
						)
						exists.Bar()

						exists.Id("read").Dot("reads").Call(Id("v"))
						exists.And()
						exists.Id("v").Dot("getType").Call().Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("Slice"))

						exists.Bar()
						exists.This().Eq().Id("read")

					})

				})
			})
	}).Line()
	utilz.Q(f)
	fmt.Printf("\n\n%#v", f)
}
