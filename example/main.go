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

	/*
		f.Any(
			Add(Int(), Id("i")),
			Add(Id("i").Gte().Lit(10)),
			DoGroup(func(f *Group) {
				f.Id("i").Eq().Lit(10)
			}),
		)
	*/

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
						par.Exists(
							List(
								Id("Function").Id("fn"),
								Id("DataFlow::CallNode").Id("call"),
							),
							DoGroup(func(st *Group) {
								st.Id("fn").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("GetSomething"))
							}),
							DoGroup(func(st *Group) {
								st.Comment("The source is the result:")

								st.Id("call").Eq().Id("fn").Dot("getACall").Call().
									And().
									This().Eq().Id("call").Dot("getResult").Call()
							}),
						)

						par.Or()

						par.Comment("Function: github.com/example/package.ParseSomething")
						par.Exists(
							List(
								Id("Function").Id("fn"),
								Id("DataFlow::CallNode").Id("call"),
							),
							DoGroup(func(st *Group) {
								st.Id("fn").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("ParseSomething"))
							}),
							DoGroup(func(st *Group) {
								st.Comment("The source is the 0th parameter:")

								st.Id("call").Eq().Id("fn").Dot("getACall").Call().
									And().
									This().Eq().Id("FunctionOutput::parameter").Call(Lit(0)).Dot("getExitNode").Call(Id("call"))
							}),
						)
					})

					metGr.Or()

					metGr.Comment("Example block 2: the type is a struct.")
					metGr.Comment("The source can be a method call (results or parameters of a call),")
					metGr.Comment("or a field read.")
					metGr.ParensFunc(func(par *Group) {
						par.Comment("Struct: github.com/example/package.Context")
						par.Exists(
							List(
								String().Id("typeName"),
							),
							DoGroup(func(st *Group) {
								st.Id("typeName").Eq().Lit("Context")
							}),
							DoGroup(func(st *Group) {
								st.Comment("Method calls on `Context`:")
								st.Exists(
									List(
										Id("DataFlow::MethodCallNode").Id("call"),
										String().Id("methodName"),
									),
									DoGroup(func(st *Group) {
										st.Id("call").Dot("getTarget").Call().Dot("hasQualifiedName").Call(
											Lit("github.com/example/package"),
											Id("typeName"),
											Id("methodName"),
										)
										st.And()
										st.ParensFunc(
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
									}),
									nil,
								)

								st.Or()
								st.Comment("Field reads on `Context`:")
								st.Exists(
									List(
										Id("DataFlow::Field").Id("fld"),
										String().Id("fieldName"),
									),
									DoGroup(func(st *Group) {
										st.Id("fld").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Id("typeName"), Id("fieldName"))

										st.And()
										st.Comment("The source is one of these fields reads:")

										st.Id("fieldName").In().Set(Lit("Accepted"), Lit("Params"))
										st.And()
										st.This().Eq().Id("fld").Dot("getARead").Call()
									}),
									nil,
								)
							}),
						)
						par.Or()

						par.Comment("Struct: github.com/example/package.SomeStruct")
						par.Exists(
							List(
								Id("DataFlow::Field").Id("fld"),
								String().Id("fieldName"),
							),
							DoGroup(func(st *Group) {
								st.Id("fld").Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("SomeStruct"), Id("fieldName"))

								st.And()
								st.Comment("The source is one of these fields reads:")

								st.Id("fieldName").In().Set(Lit("Hello"), Lit("World"))
								st.And()
								st.This().Eq().Id("fld").Dot("getARead").Call()
							}),
							nil,
						)

					})

					metGr.Or()

					metGr.Comment("Example block 3: the type is some custom type.")
					metGr.Comment("The source is a method call on that type (results or parameters of a call).")
					metGr.Comment("Some type: github.com/example/package.Slice")
					metGr.Exists(
						String().Id("typeName"),
						Id("typeName").Eq().Lit("Slice"),
						DoGroup(func(st *Group) {

							st.Comment("Method calls:")

							st.Exists(
								List(
									Id("DataFlow::MethodCallNode").Id("call"),
									String().Id("methodName"),
								),
								DoGroup(func(st *Group) {
									st.Id("call").Dot("getTarget").Call().Dot("hasQualifiedName").Call(
										Lit("github.com/example/package"),
										Id("typeName"),
										Id("methodName"),
									)
									st.And()
									st.ParensFunc(
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
								}),
								nil,
							)
						}),
					)

					metGr.Or()

					metGr.Comment("Example block 4: the type is an interface.")
					metGr.Comment("The source is a method call on that interface (results or parameters of a call).")
					metGr.Comment("Interface: github.com/example/package.SomeInterface")
					metGr.Exists(
						String().Id("typeName"),
						Id("typeName").Eq().Lit("SomeInterface"),
						DoGroup(func(st *Group) {
							st.Comment("Method calls:")

							st.Exists(
								List(
									Id("DataFlow::MethodCallNode").Id("call"),
									String().Id("methodName"),
								),
								DoGroup(func(st *Group) {
									st.Id("call").Dot("getTarget").Call().Dot("implements").Call(
										Lit("github.com/example/package"),
										Id("typeName"),
										Id("methodName"),
									)
									st.And()
									st.ParensFunc(
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
								}),
								nil,
							)
						}),
					)

					metGr.Or()

					metGr.Comment("Example block 5: the type is whatever.")
					metGr.Comment("The source is a read of a variable of that type.")
					metGr.Comment("Type: github.com/example/package.Slice")
					metGr.Exists(
						List(
							Id("DataFlow::ReadNode").Id("read"),
							Id("ValueEntity").Id("v"),
						),
						DoGroup(func(st *Group) {
							st.Id("read").Dot("reads").Call(Id("v"))
							st.And()
							st.Id("v").Dot("getType").Call().Dot("hasQualifiedName").Call(Lit("github.com/example/package"), Lit("Slice"))
						}),
						DoGroup(func(st *Group) {
							st.This().Eq().Id("read")
						}),
					)

				})
			})
	}).Line()
	utilz.Q(f)
	fmt.Printf("\n\n%#v", f)
}
