// This file is generated - do not edit.

package jen

// Parens renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func Parens(item ...Code) *Statement {
	return newStatement().Parens(item...)
}

// Parens renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func (g *Group) Parens(item ...Code) *Statement {
	s := Parens(item...)
	g.items = append(g.items, s)
	return s
}

// Parens renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func (s *Statement) Parens(item ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     item,
		multi:     true,
		name:      "parens",
		open:      "(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// ParensFunc renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func ParensFunc(f func(*Group)) *Statement {
	return newStatement().ParensFunc(f)
}

// ParensFunc renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func (g *Group) ParensFunc(f func(*Group)) *Statement {
	s := ParensFunc(f)
	g.items = append(g.items, s)
	return s
}

// ParensFunc renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func (s *Statement) ParensFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     true,
		name:      "parens",
		open:      "(",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// List renders a comma separated list. Use for multiple return functions.
func List(items ...Code) *Statement {
	return newStatement().List(items...)
}

// List renders a comma separated list. Use for multiple return functions.
func (g *Group) List(items ...Code) *Statement {
	s := List(items...)
	g.items = append(g.items, s)
	return s
}

// List renders a comma separated list. Use for multiple return functions.
func (s *Statement) List(items ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     items,
		multi:     false,
		name:      "list",
		open:      "",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// ListFunc renders a comma separated list. Use for multiple return functions.
func ListFunc(f func(*Group)) *Statement {
	return newStatement().ListFunc(f)
}

// ListFunc renders a comma separated list. Use for multiple return functions.
func (g *Group) ListFunc(f func(*Group)) *Statement {
	s := ListFunc(f)
	g.items = append(g.items, s)
	return s
}

// ListFunc renders a comma separated list. Use for multiple return functions.
func (s *Statement) ListFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "list",
		open:      "",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Values renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func Values(values ...Code) *Statement {
	return newStatement().Values(values...)
}

// Values renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (g *Group) Values(values ...Code) *Statement {
	s := Values(values...)
	g.items = append(g.items, s)
	return s
}

// Values renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (s *Statement) Values(values ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     values,
		multi:     false,
		name:      "values",
		open:      "{",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// ValuesFunc renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func ValuesFunc(f func(*Group)) *Statement {
	return newStatement().ValuesFunc(f)
}

// ValuesFunc renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (g *Group) ValuesFunc(f func(*Group)) *Statement {
	s := ValuesFunc(f)
	g.items = append(g.items, s)
	return s
}

// ValuesFunc renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (s *Statement) ValuesFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		multi:     false,
		name:      "values",
		open:      "{",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Index renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func Index(items ...Code) *Statement {
	return newStatement().Index(items...)
}

// Index renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (g *Group) Index(items ...Code) *Statement {
	s := Index(items...)
	g.items = append(g.items, s)
	return s
}

// Index renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (s *Statement) Index(items ...Code) *Statement {
	g := &Group{
		close:     "]",
		items:     items,
		multi:     false,
		name:      "index",
		open:      "[",
		separator: ":",
	}
	*s = append(*s, g)
	return s
}

// IndexFunc renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func IndexFunc(f func(*Group)) *Statement {
	return newStatement().IndexFunc(f)
}

// IndexFunc renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (g *Group) IndexFunc(f func(*Group)) *Statement {
	s := IndexFunc(f)
	g.items = append(g.items, s)
	return s
}

// IndexFunc renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (s *Statement) IndexFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "]",
		multi:     false,
		name:      "index",
		open:      "[",
		separator: ":",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Block renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func Block(statements ...Code) *Statement {
	return newStatement().Block(statements...)
}

// Block renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (g *Group) Block(statements ...Code) *Statement {
	s := Block(statements...)
	g.items = append(g.items, s)
	return s
}

// Block renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (s *Statement) Block(statements ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     statements,
		multi:     true,
		name:      "block",
		open:      "{",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func BlockFunc(f func(*Group)) *Statement {
	return newStatement().BlockFunc(f)
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (g *Group) BlockFunc(f func(*Group)) *Statement {
	s := BlockFunc(f)
	g.items = append(g.items, s)
	return s
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (s *Statement) BlockFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		multi:     true,
		name:      "block",
		open:      "{",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Defs renders a statement list enclosed in parenthesis. Use for definition lists.
func Defs(definitions ...Code) *Statement {
	return newStatement().Defs(definitions...)
}

// Defs renders a statement list enclosed in parenthesis. Use for definition lists.
func (g *Group) Defs(definitions ...Code) *Statement {
	s := Defs(definitions...)
	g.items = append(g.items, s)
	return s
}

// Defs renders a statement list enclosed in parenthesis. Use for definition lists.
func (s *Statement) Defs(definitions ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     definitions,
		multi:     true,
		name:      "defs",
		open:      "(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// DefsFunc renders a statement list enclosed in parenthesis. Use for definition lists.
func DefsFunc(f func(*Group)) *Statement {
	return newStatement().DefsFunc(f)
}

// DefsFunc renders a statement list enclosed in parenthesis. Use for definition lists.
func (g *Group) DefsFunc(f func(*Group)) *Statement {
	s := DefsFunc(f)
	g.items = append(g.items, s)
	return s
}

// DefsFunc renders a statement list enclosed in parenthesis. Use for definition lists.
func (s *Statement) DefsFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     true,
		name:      "defs",
		open:      "(",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Call renders a comma separated list enclosed by parenthesis. Use for function calls.
func Call(params ...Code) *Statement {
	return newStatement().Call(params...)
}

// Call renders a comma separated list enclosed by parenthesis. Use for function calls.
func (g *Group) Call(params ...Code) *Statement {
	s := Call(params...)
	g.items = append(g.items, s)
	return s
}

// Call renders a comma separated list enclosed by parenthesis. Use for function calls.
func (s *Statement) Call(params ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     params,
		multi:     false,
		name:      "call",
		open:      "(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// CallFunc renders a comma separated list enclosed by parenthesis. Use for function calls.
func CallFunc(f func(*Group)) *Statement {
	return newStatement().CallFunc(f)
}

// CallFunc renders a comma separated list enclosed by parenthesis. Use for function calls.
func (g *Group) CallFunc(f func(*Group)) *Statement {
	s := CallFunc(f)
	g.items = append(g.items, s)
	return s
}

// CallFunc renders a comma separated list enclosed by parenthesis. Use for function calls.
func (s *Statement) CallFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "call",
		open:      "(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Params renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func Params(params ...Code) *Statement {
	return newStatement().Params(params...)
}

// Params renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (g *Group) Params(params ...Code) *Statement {
	s := Params(params...)
	g.items = append(g.items, s)
	return s
}

// Params renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (s *Statement) Params(params ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     params,
		multi:     false,
		name:      "params",
		open:      "(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// ParamsFunc renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func ParamsFunc(f func(*Group)) *Statement {
	return newStatement().ParamsFunc(f)
}

// ParamsFunc renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (g *Group) ParamsFunc(f func(*Group)) *Statement {
	s := ParamsFunc(f)
	g.items = append(g.items, s)
	return s
}

// ParamsFunc renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (s *Statement) ParamsFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "params",
		open:      "(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Assert renders a period followed by a single item enclosed by parenthesis. Use for type assertions.
func Assert(typ Code) *Statement {
	return newStatement().Assert(typ)
}

// Assert renders a period followed by a single item enclosed by parenthesis. Use for type assertions.
func (g *Group) Assert(typ Code) *Statement {
	s := Assert(typ)
	g.items = append(g.items, s)
	return s
}

// Assert renders a period followed by a single item enclosed by parenthesis. Use for type assertions.
func (s *Statement) Assert(typ Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{typ},
		multi:     false,
		name:      "assert",
		open:      ".(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// Map renders the keyword followed by a single item enclosed by square brackets. Use for map definitions.
func Map(typ Code) *Statement {
	return newStatement().Map(typ)
}

// Map renders the keyword followed by a single item enclosed by square brackets. Use for map definitions.
func (g *Group) Map(typ Code) *Statement {
	s := Map(typ)
	g.items = append(g.items, s)
	return s
}

// Map renders the keyword followed by a single item enclosed by square brackets. Use for map definitions.
func (s *Statement) Map(typ Code) *Statement {
	g := &Group{
		close:     "]",
		items:     []Code{typ},
		multi:     false,
		name:      "map",
		open:      "map[",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// If renders the keyword followed by a semicolon separated list.
func If(conditions ...Code) *Statement {
	return newStatement().If(conditions...)
}

// If renders the keyword followed by a semicolon separated list.
func (g *Group) If(conditions ...Code) *Statement {
	s := If(conditions...)
	g.items = append(g.items, s)
	return s
}

// If renders the keyword followed by a semicolon separated list.
func (s *Statement) If(conditions ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     conditions,
		multi:     false,
		name:      "if",
		open:      "if ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// IfFunc renders the keyword followed by a semicolon separated list.
func IfFunc(f func(*Group)) *Statement {
	return newStatement().IfFunc(f)
}

// IfFunc renders the keyword followed by a semicolon separated list.
func (g *Group) IfFunc(f func(*Group)) *Statement {
	s := IfFunc(f)
	g.items = append(g.items, s)
	return s
}

// IfFunc renders the keyword followed by a semicolon separated list.
func (s *Statement) IfFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "if",
		open:      "if ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Return renders the keyword followed by a comma separated list.
func Return(results ...Code) *Statement {
	return newStatement().Return(results...)
}

// Return renders the keyword followed by a comma separated list.
func (g *Group) Return(results ...Code) *Statement {
	s := Return(results...)
	g.items = append(g.items, s)
	return s
}

// Return renders the keyword followed by a comma separated list.
func (s *Statement) Return(results ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     results,
		multi:     false,
		name:      "return",
		open:      "return ",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// ReturnFunc renders the keyword followed by a comma separated list.
func ReturnFunc(f func(*Group)) *Statement {
	return newStatement().ReturnFunc(f)
}

// ReturnFunc renders the keyword followed by a comma separated list.
func (g *Group) ReturnFunc(f func(*Group)) *Statement {
	s := ReturnFunc(f)
	g.items = append(g.items, s)
	return s
}

// ReturnFunc renders the keyword followed by a comma separated list.
func (s *Statement) ReturnFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "return",
		open:      "return ",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// For renders the keyword followed by a semicolon separated list.
func For(conditions ...Code) *Statement {
	return newStatement().For(conditions...)
}

// For renders the keyword followed by a semicolon separated list.
func (g *Group) For(conditions ...Code) *Statement {
	s := For(conditions...)
	g.items = append(g.items, s)
	return s
}

// For renders the keyword followed by a semicolon separated list.
func (s *Statement) For(conditions ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     conditions,
		multi:     false,
		name:      "for",
		open:      "for ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// ForFunc renders the keyword followed by a semicolon separated list.
func ForFunc(f func(*Group)) *Statement {
	return newStatement().ForFunc(f)
}

// ForFunc renders the keyword followed by a semicolon separated list.
func (g *Group) ForFunc(f func(*Group)) *Statement {
	s := ForFunc(f)
	g.items = append(g.items, s)
	return s
}

// ForFunc renders the keyword followed by a semicolon separated list.
func (s *Statement) ForFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "for",
		open:      "for ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Switch renders the keyword followed by a semicolon separated list.
func Switch(conditions ...Code) *Statement {
	return newStatement().Switch(conditions...)
}

// Switch renders the keyword followed by a semicolon separated list.
func (g *Group) Switch(conditions ...Code) *Statement {
	s := Switch(conditions...)
	g.items = append(g.items, s)
	return s
}

// Switch renders the keyword followed by a semicolon separated list.
func (s *Statement) Switch(conditions ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     conditions,
		multi:     false,
		name:      "switch",
		open:      "switch ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// SwitchFunc renders the keyword followed by a semicolon separated list.
func SwitchFunc(f func(*Group)) *Statement {
	return newStatement().SwitchFunc(f)
}

// SwitchFunc renders the keyword followed by a semicolon separated list.
func (g *Group) SwitchFunc(f func(*Group)) *Statement {
	s := SwitchFunc(f)
	g.items = append(g.items, s)
	return s
}

// SwitchFunc renders the keyword followed by a semicolon separated list.
func (s *Statement) SwitchFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "switch",
		open:      "switch ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Interface renders the keyword followed by a method list enclosed by curly braces.
func Interface(methods ...Code) *Statement {
	return newStatement().Interface(methods...)
}

// Interface renders the keyword followed by a method list enclosed by curly braces.
func (g *Group) Interface(methods ...Code) *Statement {
	s := Interface(methods...)
	g.items = append(g.items, s)
	return s
}

// Interface renders the keyword followed by a method list enclosed by curly braces.
func (s *Statement) Interface(methods ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     methods,
		multi:     true,
		name:      "interface",
		open:      "interface{",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// InterfaceFunc renders the keyword followed by a method list enclosed by curly braces.
func InterfaceFunc(f func(*Group)) *Statement {
	return newStatement().InterfaceFunc(f)
}

// InterfaceFunc renders the keyword followed by a method list enclosed by curly braces.
func (g *Group) InterfaceFunc(f func(*Group)) *Statement {
	s := InterfaceFunc(f)
	g.items = append(g.items, s)
	return s
}

// InterfaceFunc renders the keyword followed by a method list enclosed by curly braces.
func (s *Statement) InterfaceFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		multi:     true,
		name:      "interface",
		open:      "interface{",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Struct renders the keyword followed by a field list enclosed by curly braces.
func Struct(fields ...Code) *Statement {
	return newStatement().Struct(fields...)
}

// Struct renders the keyword followed by a field list enclosed by curly braces.
func (g *Group) Struct(fields ...Code) *Statement {
	s := Struct(fields...)
	g.items = append(g.items, s)
	return s
}

// Struct renders the keyword followed by a field list enclosed by curly braces.
func (s *Statement) Struct(fields ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     fields,
		multi:     true,
		name:      "struct",
		open:      "struct{",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// StructFunc renders the keyword followed by a field list enclosed by curly braces.
func StructFunc(f func(*Group)) *Statement {
	return newStatement().StructFunc(f)
}

// StructFunc renders the keyword followed by a field list enclosed by curly braces.
func (g *Group) StructFunc(f func(*Group)) *Statement {
	s := StructFunc(f)
	g.items = append(g.items, s)
	return s
}

// StructFunc renders the keyword followed by a field list enclosed by curly braces.
func (s *Statement) StructFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		multi:     true,
		name:      "struct",
		open:      "struct{",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Case renders the keyword followed by a comma separated list.
func Case(cases ...Code) *Statement {
	return newStatement().Case(cases...)
}

// Case renders the keyword followed by a comma separated list.
func (g *Group) Case(cases ...Code) *Statement {
	s := Case(cases...)
	g.items = append(g.items, s)
	return s
}

// Case renders the keyword followed by a comma separated list.
func (s *Statement) Case(cases ...Code) *Statement {
	g := &Group{
		close:     ":",
		items:     cases,
		multi:     false,
		name:      "case",
		open:      "case ",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// CaseFunc renders the keyword followed by a comma separated list.
func CaseFunc(f func(*Group)) *Statement {
	return newStatement().CaseFunc(f)
}

// CaseFunc renders the keyword followed by a comma separated list.
func (g *Group) CaseFunc(f func(*Group)) *Statement {
	s := CaseFunc(f)
	g.items = append(g.items, s)
	return s
}

// CaseFunc renders the keyword followed by a comma separated list.
func (s *Statement) CaseFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ":",
		multi:     false,
		name:      "case",
		open:      "case ",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Append renders the append built-in function.
func Append(args ...Code) *Statement {
	return newStatement().Append(args...)
}

// Append renders the append built-in function.
func (g *Group) Append(args ...Code) *Statement {
	s := Append(args...)
	g.items = append(g.items, s)
	return s
}

// Append renders the append built-in function.
func (s *Statement) Append(args ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     args,
		multi:     false,
		name:      "append",
		open:      "append(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// AppendFunc renders the append built-in function.
func AppendFunc(f func(*Group)) *Statement {
	return newStatement().AppendFunc(f)
}

// AppendFunc renders the append built-in function.
func (g *Group) AppendFunc(f func(*Group)) *Statement {
	s := AppendFunc(f)
	g.items = append(g.items, s)
	return s
}

// AppendFunc renders the append built-in function.
func (s *Statement) AppendFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "append",
		open:      "append(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Cap renders the cap built-in function.
func Cap(v Code) *Statement {
	return newStatement().Cap(v)
}

// Cap renders the cap built-in function.
func (g *Group) Cap(v Code) *Statement {
	s := Cap(v)
	g.items = append(g.items, s)
	return s
}

// Cap renders the cap built-in function.
func (s *Statement) Cap(v Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{v},
		multi:     false,
		name:      "cap",
		open:      "cap(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Close renders the close built-in function.
func Close(c Code) *Statement {
	return newStatement().Close(c)
}

// Close renders the close built-in function.
func (g *Group) Close(c Code) *Statement {
	s := Close(c)
	g.items = append(g.items, s)
	return s
}

// Close renders the close built-in function.
func (s *Statement) Close(c Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{c},
		multi:     false,
		name:      "close",
		open:      "close(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Complex renders the complex built-in function.
func Complex(r Code, i Code) *Statement {
	return newStatement().Complex(r, i)
}

// Complex renders the complex built-in function.
func (g *Group) Complex(r Code, i Code) *Statement {
	s := Complex(r, i)
	g.items = append(g.items, s)
	return s
}

// Complex renders the complex built-in function.
func (s *Statement) Complex(r Code, i Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{r, i},
		multi:     false,
		name:      "complex",
		open:      "complex(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Copy renders the copy built-in function.
func Copy(dst Code, src Code) *Statement {
	return newStatement().Copy(dst, src)
}

// Copy renders the copy built-in function.
func (g *Group) Copy(dst Code, src Code) *Statement {
	s := Copy(dst, src)
	g.items = append(g.items, s)
	return s
}

// Copy renders the copy built-in function.
func (s *Statement) Copy(dst Code, src Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{dst, src},
		multi:     false,
		name:      "copy",
		open:      "copy(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Delete renders the delete built-in function.
func Delete(m Code, key Code) *Statement {
	return newStatement().Delete(m, key)
}

// Delete renders the delete built-in function.
func (g *Group) Delete(m Code, key Code) *Statement {
	s := Delete(m, key)
	g.items = append(g.items, s)
	return s
}

// Delete renders the delete built-in function.
func (s *Statement) Delete(m Code, key Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{m, key},
		multi:     false,
		name:      "delete",
		open:      "delete(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Imag renders the imag built-in function.
func Imag(c Code) *Statement {
	return newStatement().Imag(c)
}

// Imag renders the imag built-in function.
func (g *Group) Imag(c Code) *Statement {
	s := Imag(c)
	g.items = append(g.items, s)
	return s
}

// Imag renders the imag built-in function.
func (s *Statement) Imag(c Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{c},
		multi:     false,
		name:      "imag",
		open:      "imag(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Len renders the len built-in function.
func Len(v Code) *Statement {
	return newStatement().Len(v)
}

// Len renders the len built-in function.
func (g *Group) Len(v Code) *Statement {
	s := Len(v)
	g.items = append(g.items, s)
	return s
}

// Len renders the len built-in function.
func (s *Statement) Len(v Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{v},
		multi:     false,
		name:      "len",
		open:      "len(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Make renders the make built-in function. The final parameter of the make function is optional, so it is represented by a variadic parameter list.
func Make(args ...Code) *Statement {
	return newStatement().Make(args...)
}

// Make renders the make built-in function. The final parameter of the make function is optional, so it is represented by a variadic parameter list.
func (g *Group) Make(args ...Code) *Statement {
	s := Make(args...)
	g.items = append(g.items, s)
	return s
}

// Make renders the make built-in function. The final parameter of the make function is optional, so it is represented by a variadic parameter list.
func (s *Statement) Make(args ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     args,
		multi:     false,
		name:      "make",
		open:      "make(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// New renders the new built-in function.
func New(typ Code) *Statement {
	return newStatement().New(typ)
}

// New renders the new built-in function.
func (g *Group) New(typ Code) *Statement {
	s := New(typ)
	g.items = append(g.items, s)
	return s
}

// New renders the new built-in function.
func (s *Statement) New(typ Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{typ},
		multi:     false,
		name:      "new",
		open:      "new(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Panic renders the panic built-in function.
func Panic(v Code) *Statement {
	return newStatement().Panic(v)
}

// Panic renders the panic built-in function.
func (g *Group) Panic(v Code) *Statement {
	s := Panic(v)
	g.items = append(g.items, s)
	return s
}

// Panic renders the panic built-in function.
func (s *Statement) Panic(v Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{v},
		multi:     false,
		name:      "panic",
		open:      "panic(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Print renders the print built-in function.
func Print(args ...Code) *Statement {
	return newStatement().Print(args...)
}

// Print renders the print built-in function.
func (g *Group) Print(args ...Code) *Statement {
	s := Print(args...)
	g.items = append(g.items, s)
	return s
}

// Print renders the print built-in function.
func (s *Statement) Print(args ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     args,
		multi:     false,
		name:      "print",
		open:      "print(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// PrintFunc renders the print built-in function.
func PrintFunc(f func(*Group)) *Statement {
	return newStatement().PrintFunc(f)
}

// PrintFunc renders the print built-in function.
func (g *Group) PrintFunc(f func(*Group)) *Statement {
	s := PrintFunc(f)
	g.items = append(g.items, s)
	return s
}

// PrintFunc renders the print built-in function.
func (s *Statement) PrintFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "print",
		open:      "print(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Println renders the println built-in function.
func Println(args ...Code) *Statement {
	return newStatement().Println(args...)
}

// Println renders the println built-in function.
func (g *Group) Println(args ...Code) *Statement {
	s := Println(args...)
	g.items = append(g.items, s)
	return s
}

// Println renders the println built-in function.
func (s *Statement) Println(args ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     args,
		multi:     false,
		name:      "println",
		open:      "println(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// PrintlnFunc renders the println built-in function.
func PrintlnFunc(f func(*Group)) *Statement {
	return newStatement().PrintlnFunc(f)
}

// PrintlnFunc renders the println built-in function.
func (g *Group) PrintlnFunc(f func(*Group)) *Statement {
	s := PrintlnFunc(f)
	g.items = append(g.items, s)
	return s
}

// PrintlnFunc renders the println built-in function.
func (s *Statement) PrintlnFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "println",
		open:      "println(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Real renders the real built-in function.
func Real(c Code) *Statement {
	return newStatement().Real(c)
}

// Real renders the real built-in function.
func (g *Group) Real(c Code) *Statement {
	s := Real(c)
	g.items = append(g.items, s)
	return s
}

// Real renders the real built-in function.
func (s *Statement) Real(c Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{c},
		multi:     false,
		name:      "real",
		open:      "real(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Recover renders the recover built-in function.
func Recover() *Statement {
	return newStatement().Recover()
}

// Recover renders the recover built-in function.
func (g *Group) Recover() *Statement {
	s := Recover()
	g.items = append(g.items, s)
	return s
}

// Recover renders the recover built-in function.
func (s *Statement) Recover() *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{},
		multi:     false,
		name:      "recover",
		open:      "recover(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Exists renders the exists quantifier.
func Exists(args ...Code) *Statement {
	return newStatement().Exists(args...)
}

// Exists renders the exists quantifier.
func (g *Group) Exists(args ...Code) *Statement {
	s := Exists(args...)
	g.items = append(g.items, s)
	return s
}

// Exists renders the exists quantifier.
func (s *Statement) Exists(args ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     args,
		multi:     true,
		name:      "exists",
		open:      "exists(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// ExistsFunc renders the exists quantifier.
func ExistsFunc(f func(*Group)) *Statement {
	return newStatement().ExistsFunc(f)
}

// ExistsFunc renders the exists quantifier.
func (g *Group) ExistsFunc(f func(*Group)) *Statement {
	s := ExistsFunc(f)
	g.items = append(g.items, s)
	return s
}

// ExistsFunc renders the exists quantifier.
func (s *Statement) ExistsFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     true,
		name:      "exists",
		open:      "exists(",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Set renders a set literal expressions, i.e. a comma separated list enclosed by square brackets.
func Set(items ...Code) *Statement {
	return newStatement().Set(items...)
}

// Set renders a set literal expressions, i.e. a comma separated list enclosed by square brackets.
func (g *Group) Set(items ...Code) *Statement {
	s := Set(items...)
	g.items = append(g.items, s)
	return s
}

// Set renders a set literal expressions, i.e. a comma separated list enclosed by square brackets.
func (s *Statement) Set(items ...Code) *Statement {
	g := &Group{
		close:     "]",
		items:     items,
		multi:     false,
		name:      "set",
		open:      "[",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// SetFunc renders a set literal expressions, i.e. a comma separated list enclosed by square brackets.
func SetFunc(f func(*Group)) *Statement {
	return newStatement().SetFunc(f)
}

// SetFunc renders a set literal expressions, i.e. a comma separated list enclosed by square brackets.
func (g *Group) SetFunc(f func(*Group)) *Statement {
	s := SetFunc(f)
	g.items = append(g.items, s)
	return s
}

// SetFunc renders a set literal expressions, i.e. a comma separated list enclosed by square brackets.
func (s *Statement) SetFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "]",
		multi:     false,
		name:      "set",
		open:      "[",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Eq renders a =.
func Eq() *Statement {
	return newStatement().Eq()
}

// Eq renders a =.
func (g *Group) Eq() *Statement {
	s := Eq()
	g.items = append(g.items, s)
	return s
}

// Eq renders a =.
func (s *Statement) Eq() *Statement {
	g := &Group{
		close:     "",
		items:     []Code{},
		multi:     false,
		name:      "eq",
		open:      "=",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// Boolean renders the boolean identifier.
func Boolean() *Statement {
	return newStatement().Boolean()
}

// Boolean renders the boolean identifier.
func (g *Group) Boolean() *Statement {
	s := Boolean()
	g.items = append(g.items, s)
	return s
}

// Boolean renders the boolean identifier.
func (s *Statement) Boolean() *Statement {
	t := token{
		content: "boolean",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// True renders the true identifier.
func True() *Statement {
	// notest
	return newStatement().True()
}

// True renders the true identifier.
func (g *Group) True() *Statement {
	// notest
	s := True()
	g.items = append(g.items, s)
	return s
}

// True renders the true identifier.
func (s *Statement) True() *Statement {
	// notest
	t := token{
		content: "true",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// False renders the false identifier.
func False() *Statement {
	// notest
	return newStatement().False()
}

// False renders the false identifier.
func (g *Group) False() *Statement {
	// notest
	s := False()
	g.items = append(g.items, s)
	return s
}

// False renders the false identifier.
func (s *Statement) False() *Statement {
	// notest
	t := token{
		content: "false",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Float renders the float identifier.
func Float() *Statement {
	// notest
	return newStatement().Float()
}

// Float renders the float identifier.
func (g *Group) Float() *Statement {
	// notest
	s := Float()
	g.items = append(g.items, s)
	return s
}

// Float renders the float identifier.
func (s *Statement) Float() *Statement {
	// notest
	t := token{
		content: "float",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Int renders the int identifier.
func Int() *Statement {
	// notest
	return newStatement().Int()
}

// Int renders the int identifier.
func (g *Group) Int() *Statement {
	// notest
	s := Int()
	g.items = append(g.items, s)
	return s
}

// Int renders the int identifier.
func (s *Statement) Int() *Statement {
	// notest
	t := token{
		content: "int",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// String renders the string identifier.
func String() *Statement {
	// notest
	return newStatement().String()
}

// String renders the string identifier.
func (g *Group) String() *Statement {
	// notest
	s := String()
	g.items = append(g.items, s)
	return s
}

// String renders the string identifier.
func (s *Statement) String() *Statement {
	// notest
	t := token{
		content: "string",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Date renders the date identifier.
func Date() *Statement {
	// notest
	return newStatement().Date()
}

// Date renders the date identifier.
func (g *Group) Date() *Statement {
	// notest
	s := Date()
	g.items = append(g.items, s)
	return s
}

// Date renders the date identifier.
func (s *Statement) Date() *Statement {
	// notest
	t := token{
		content: "date",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Private renders the private keyword.
func Private() *Statement {
	// notest
	return newStatement().Private()
}

// Private renders the private keyword.
func (g *Group) Private() *Statement {
	// notest
	s := Private()
	g.items = append(g.items, s)
	return s
}

// Private renders the private keyword.
func (s *Statement) Private() *Statement {
	// notest
	t := token{
		content: "private",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Module renders the module keyword.
func Module() *Statement {
	// notest
	return newStatement().Module()
}

// Module renders the module keyword.
func (g *Group) Module() *Statement {
	// notest
	s := Module()
	g.items = append(g.items, s)
	return s
}

// Module renders the module keyword.
func (s *Statement) Module() *Statement {
	// notest
	t := token{
		content: "module",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Class renders the class keyword.
func Class() *Statement {
	// notest
	return newStatement().Class()
}

// Class renders the class keyword.
func (g *Group) Class() *Statement {
	// notest
	s := Class()
	g.items = append(g.items, s)
	return s
}

// Class renders the class keyword.
func (s *Statement) Class() *Statement {
	// notest
	t := token{
		content: "class",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Extends renders the extends keyword.
func Extends() *Statement {
	// notest
	return newStatement().Extends()
}

// Extends renders the extends keyword.
func (g *Group) Extends() *Statement {
	// notest
	s := Extends()
	g.items = append(g.items, s)
	return s
}

// Extends renders the extends keyword.
func (s *Statement) Extends() *Statement {
	// notest
	t := token{
		content: "extends",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Or renders the or keyword.
func Or() *Statement {
	// notest
	return newStatement().Or()
}

// Or renders the or keyword.
func (g *Group) Or() *Statement {
	// notest
	s := Or()
	g.items = append(g.items, s)
	return s
}

// Or renders the or keyword.
func (s *Statement) Or() *Statement {
	// notest
	t := token{
		content: "or",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// And renders the and keyword.
func And() *Statement {
	// notest
	return newStatement().And()
}

// And renders the and keyword.
func (g *Group) And() *Statement {
	// notest
	s := And()
	g.items = append(g.items, s)
	return s
}

// And renders the and keyword.
func (s *Statement) And() *Statement {
	// notest
	t := token{
		content: "and",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// In renders the in keyword.
func In() *Statement {
	// notest
	return newStatement().In()
}

// In renders the in keyword.
func (g *Group) In() *Statement {
	// notest
	s := In()
	g.items = append(g.items, s)
	return s
}

// In renders the in keyword.
func (s *Statement) In() *Statement {
	// notest
	t := token{
		content: "in",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// This renders the this keyword.
func This() *Statement {
	// notest
	return newStatement().This()
}

// This renders the this keyword.
func (g *Group) This() *Statement {
	// notest
	s := This()
	g.items = append(g.items, s)
	return s
}

// This renders the this keyword.
func (s *Statement) This() *Statement {
	// notest
	t := token{
		content: "this",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}
