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

// Block renders a statement list enclosed by curly braces. Use for code blocks.
func Block(statements ...Code) *Statement {
	return newStatement().Block(statements...)
}

// Block renders a statement list enclosed by curly braces. Use for code blocks.
func (g *Group) Block(statements ...Code) *Statement {
	s := Block(statements...)
	g.items = append(g.items, s)
	return s
}

// Block renders a statement list enclosed by curly braces. Use for code blocks.
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

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks.
func BlockFunc(f func(*Group)) *Statement {
	return newStatement().BlockFunc(f)
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks.
func (g *Group) BlockFunc(f func(*Group)) *Statement {
	s := BlockFunc(f)
	g.items = append(g.items, s)
	return s
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks.
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

// Cast renders a period followed by a single item enclosed by parenthesis. Use for type casting.
func Cast(typ Code) *Statement {
	return newStatement().Cast(typ)
}

// Cast renders a period followed by a single item enclosed by parenthesis. Use for type casting.
func (g *Group) Cast(typ Code) *Statement {
	s := Cast(typ)
	g.items = append(g.items, s)
	return s
}

// Cast renders a period followed by a single item enclosed by parenthesis. Use for type casting.
func (s *Statement) Cast(typ Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{typ},
		multi:     false,
		name:      "cast",
		open:      ".(",
		separator: "",
	}
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

// Set renders a set literal expression, i.e. a comma separated list enclosed by square brackets.
func Set(items ...Code) *Statement {
	return newStatement().Set(items...)
}

// Set renders a set literal expression, i.e. a comma separated list enclosed by square brackets.
func (g *Group) Set(items ...Code) *Statement {
	s := Set(items...)
	g.items = append(g.items, s)
	return s
}

// Set renders a set literal expression, i.e. a comma separated list enclosed by square brackets.
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

// SetFunc renders a set literal expression, i.e. a comma separated list enclosed by square brackets.
func SetFunc(f func(*Group)) *Statement {
	return newStatement().SetFunc(f)
}

// SetFunc renders a set literal expression, i.e. a comma separated list enclosed by square brackets.
func (g *Group) SetFunc(f func(*Group)) *Statement {
	s := SetFunc(f)
	g.items = append(g.items, s)
	return s
}

// SetFunc renders a set literal expression, i.e. a comma separated list enclosed by square brackets.
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

// Any renders an `any` expression.
func Any(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Any(vars, formula, expression)
}

// Any renders an `any` expression.
func (g *Group) Any(vars Code, formula Code, expression Code) *Statement {
	s := Any(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Any renders an `any` expression.
func (s *Statement) Any(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "any",
		open:      "any(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Exists renders the `exists` quantifier.
func Exists(vars Code, formula1 Code, formula2 Code) *Statement {
	return newStatement().Exists(vars, formula1, formula2)
}

// Exists renders the `exists` quantifier.
func (g *Group) Exists(vars Code, formula1 Code, formula2 Code) *Statement {
	s := Exists(vars, formula1, formula2)
	g.items = append(g.items, s)
	return s
}

// Exists renders the `exists` quantifier.
func (s *Statement) Exists(vars Code, formula1 Code, formula2 Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula1, formula2},
		multi:     true,
		name:      "exists",
		open:      "exists(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// ForAll renders the `forall` quantifier.
func ForAll(vars Code, formula1 Code, formula2 Code) *Statement {
	return newStatement().ForAll(vars, formula1, formula2)
}

// ForAll renders the `forall` quantifier.
func (g *Group) ForAll(vars Code, formula1 Code, formula2 Code) *Statement {
	s := ForAll(vars, formula1, formula2)
	g.items = append(g.items, s)
	return s
}

// ForAll renders the `forall` quantifier.
func (s *Statement) ForAll(vars Code, formula1 Code, formula2 Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula1, formula2},
		multi:     true,
		name:      "forall",
		open:      "forall(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// ForEx renders the `forex` quantifier.
func ForEx(vars Code, formula1 Code, formula2 Code) *Statement {
	return newStatement().ForEx(vars, formula1, formula2)
}

// ForEx renders the `forex` quantifier.
func (g *Group) ForEx(vars Code, formula1 Code, formula2 Code) *Statement {
	s := ForEx(vars, formula1, formula2)
	g.items = append(g.items, s)
	return s
}

// ForEx renders the `forex` quantifier.
func (s *Statement) ForEx(vars Code, formula1 Code, formula2 Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula1, formula2},
		multi:     true,
		name:      "forex",
		open:      "forex(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Count renders a `count` aggregator.
func Count(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Count(vars, formula, expression)
}

// Count renders a `count` aggregator.
func (g *Group) Count(vars Code, formula Code, expression Code) *Statement {
	s := Count(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Count renders a `count` aggregator.
func (s *Statement) Count(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "count",
		open:      "count(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Min renders a `min` aggregator.
func Min(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Min(vars, formula, expression)
}

// Min renders a `min` aggregator.
func (g *Group) Min(vars Code, formula Code, expression Code) *Statement {
	s := Min(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Min renders a `min` aggregator.
func (s *Statement) Min(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "min",
		open:      "min(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Max renders a `max` aggregator.
func Max(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Max(vars, formula, expression)
}

// Max renders a `max` aggregator.
func (g *Group) Max(vars Code, formula Code, expression Code) *Statement {
	s := Max(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Max renders a `max` aggregator.
func (s *Statement) Max(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "max",
		open:      "max(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Avg renders an `avg` aggregator.
func Avg(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Avg(vars, formula, expression)
}

// Avg renders an `avg` aggregator.
func (g *Group) Avg(vars Code, formula Code, expression Code) *Statement {
	s := Avg(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Avg renders an `avg` aggregator.
func (s *Statement) Avg(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "avg",
		open:      "avg(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Sum renders a `sum` aggregator.
func Sum(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Sum(vars, formula, expression)
}

// Sum renders a `sum` aggregator.
func (g *Group) Sum(vars Code, formula Code, expression Code) *Statement {
	s := Sum(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Sum renders a `sum` aggregator.
func (s *Statement) Sum(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "sum",
		open:      "sum(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Concat renders a `concat` aggregator.
func Concat(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Concat(vars, formula, expression)
}

// Concat renders a `concat` aggregator.
func (g *Group) Concat(vars Code, formula Code, expression Code) *Statement {
	s := Concat(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Concat renders a `concat` aggregator.
func (s *Statement) Concat(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "concat",
		open:      "concat(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Unique renders a `unique` aggregator.
func Unique(vars Code, formula Code, expression Code) *Statement {
	return newStatement().Unique(vars, formula, expression)
}

// Unique renders a `unique` aggregator.
func (g *Group) Unique(vars Code, formula Code, expression Code) *Statement {
	s := Unique(vars, formula, expression)
	g.items = append(g.items, s)
	return s
}

// Unique renders a `unique` aggregator.
func (s *Statement) Unique(vars Code, formula Code, expression Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{vars, formula, expression},
		multi:     true,
		name:      "unique",
		open:      "unique(",
		separator: " | ",
	}
	*s = append(*s, g)
	return s
}

// Boolean renders the `boolean` identifier.
func Boolean() *Statement {
	return newStatement().Boolean()
}

// Boolean renders the `boolean` identifier.
func (g *Group) Boolean() *Statement {
	s := Boolean()
	g.items = append(g.items, s)
	return s
}

// Boolean renders the `boolean` identifier.
func (s *Statement) Boolean() *Statement {
	t := token{
		content: "boolean",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// True renders the `true` identifier.
func True() *Statement {
	// notest
	return newStatement().True()
}

// True renders the `true` identifier.
func (g *Group) True() *Statement {
	// notest
	s := True()
	g.items = append(g.items, s)
	return s
}

// True renders the `true` identifier.
func (s *Statement) True() *Statement {
	// notest
	t := token{
		content: "true",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// False renders the `false` identifier.
func False() *Statement {
	// notest
	return newStatement().False()
}

// False renders the `false` identifier.
func (g *Group) False() *Statement {
	// notest
	s := False()
	g.items = append(g.items, s)
	return s
}

// False renders the `false` identifier.
func (s *Statement) False() *Statement {
	// notest
	t := token{
		content: "false",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Float renders the `float` identifier.
func Float() *Statement {
	// notest
	return newStatement().Float()
}

// Float renders the `float` identifier.
func (g *Group) Float() *Statement {
	// notest
	s := Float()
	g.items = append(g.items, s)
	return s
}

// Float renders the `float` identifier.
func (s *Statement) Float() *Statement {
	// notest
	t := token{
		content: "float",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Int renders the `int` identifier.
func Int() *Statement {
	// notest
	return newStatement().Int()
}

// Int renders the `int` identifier.
func (g *Group) Int() *Statement {
	// notest
	s := Int()
	g.items = append(g.items, s)
	return s
}

// Int renders the `int` identifier.
func (s *Statement) Int() *Statement {
	// notest
	t := token{
		content: "int",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// String renders the `string` identifier.
func String() *Statement {
	// notest
	return newStatement().String()
}

// String renders the `string` identifier.
func (g *Group) String() *Statement {
	// notest
	s := String()
	g.items = append(g.items, s)
	return s
}

// String renders the `string` identifier.
func (s *Statement) String() *Statement {
	// notest
	t := token{
		content: "string",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Date renders the `date` identifier.
func Date() *Statement {
	// notest
	return newStatement().Date()
}

// Date renders the `date` identifier.
func (g *Group) Date() *Statement {
	// notest
	s := Date()
	g.items = append(g.items, s)
	return s
}

// Date renders the `date` identifier.
func (s *Statement) Date() *Statement {
	// notest
	t := token{
		content: "date",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Abstract renders the `abstract` keyword.
func Abstract() *Statement {
	// notest
	return newStatement().Abstract()
}

// Abstract renders the `abstract` keyword.
func (g *Group) Abstract() *Statement {
	// notest
	s := Abstract()
	g.items = append(g.items, s)
	return s
}

// Abstract renders the `abstract` keyword.
func (s *Statement) Abstract() *Statement {
	// notest
	t := token{
		content: "abstract",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// And renders the `and` keyword.
func And() *Statement {
	// notest
	return newStatement().And()
}

// And renders the `and` keyword.
func (g *Group) And() *Statement {
	// notest
	s := And()
	g.items = append(g.items, s)
	return s
}

// And renders the `and` keyword.
func (s *Statement) And() *Statement {
	// notest
	t := token{
		content: "and",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// As renders the `as` keyword.
func As() *Statement {
	// notest
	return newStatement().As()
}

// As renders the `as` keyword.
func (g *Group) As() *Statement {
	// notest
	s := As()
	g.items = append(g.items, s)
	return s
}

// As renders the `as` keyword.
func (s *Statement) As() *Statement {
	// notest
	t := token{
		content: "as",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Asc renders the `asc` keyword.
func Asc() *Statement {
	// notest
	return newStatement().Asc()
}

// Asc renders the `asc` keyword.
func (g *Group) Asc() *Statement {
	// notest
	s := Asc()
	g.items = append(g.items, s)
	return s
}

// Asc renders the `asc` keyword.
func (s *Statement) Asc() *Statement {
	// notest
	t := token{
		content: "asc",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Class renders the `class` keyword.
func Class() *Statement {
	// notest
	return newStatement().Class()
}

// Class renders the `class` keyword.
func (g *Group) Class() *Statement {
	// notest
	s := Class()
	g.items = append(g.items, s)
	return s
}

// Class renders the `class` keyword.
func (s *Statement) Class() *Statement {
	// notest
	t := token{
		content: "class",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Desc renders the `desc` keyword.
func Desc() *Statement {
	// notest
	return newStatement().Desc()
}

// Desc renders the `desc` keyword.
func (g *Group) Desc() *Statement {
	// notest
	s := Desc()
	g.items = append(g.items, s)
	return s
}

// Desc renders the `desc` keyword.
func (s *Statement) Desc() *Statement {
	// notest
	t := token{
		content: "desc",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Else renders the `else` keyword.
func Else() *Statement {
	// notest
	return newStatement().Else()
}

// Else renders the `else` keyword.
func (g *Group) Else() *Statement {
	// notest
	s := Else()
	g.items = append(g.items, s)
	return s
}

// Else renders the `else` keyword.
func (s *Statement) Else() *Statement {
	// notest
	t := token{
		content: "else",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Extends renders the `extends` keyword.
func Extends() *Statement {
	// notest
	return newStatement().Extends()
}

// Extends renders the `extends` keyword.
func (g *Group) Extends() *Statement {
	// notest
	s := Extends()
	g.items = append(g.items, s)
	return s
}

// Extends renders the `extends` keyword.
func (s *Statement) Extends() *Statement {
	// notest
	t := token{
		content: "extends",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// If renders the `if` keyword.
func If() *Statement {
	// notest
	return newStatement().If()
}

// If renders the `if` keyword.
func (g *Group) If() *Statement {
	// notest
	s := If()
	g.items = append(g.items, s)
	return s
}

// If renders the `if` keyword.
func (s *Statement) If() *Statement {
	// notest
	t := token{
		content: "if",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Implies renders the `implies` keyword.
func Implies() *Statement {
	// notest
	return newStatement().Implies()
}

// Implies renders the `implies` keyword.
func (g *Group) Implies() *Statement {
	// notest
	s := Implies()
	g.items = append(g.items, s)
	return s
}

// Implies renders the `implies` keyword.
func (s *Statement) Implies() *Statement {
	// notest
	t := token{
		content: "implies",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// In renders the `in` keyword.
func In() *Statement {
	// notest
	return newStatement().In()
}

// In renders the `in` keyword.
func (g *Group) In() *Statement {
	// notest
	s := In()
	g.items = append(g.items, s)
	return s
}

// In renders the `in` keyword.
func (s *Statement) In() *Statement {
	// notest
	t := token{
		content: "in",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Instanceof renders the `instanceof` keyword.
func Instanceof() *Statement {
	// notest
	return newStatement().Instanceof()
}

// Instanceof renders the `instanceof` keyword.
func (g *Group) Instanceof() *Statement {
	// notest
	s := Instanceof()
	g.items = append(g.items, s)
	return s
}

// Instanceof renders the `instanceof` keyword.
func (s *Statement) Instanceof() *Statement {
	// notest
	t := token{
		content: "instanceof",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Module renders the `module` keyword.
func Module() *Statement {
	// notest
	return newStatement().Module()
}

// Module renders the `module` keyword.
func (g *Group) Module() *Statement {
	// notest
	s := Module()
	g.items = append(g.items, s)
	return s
}

// Module renders the `module` keyword.
func (s *Statement) Module() *Statement {
	// notest
	t := token{
		content: "module",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Newtype renders the `newtype` keyword.
func Newtype() *Statement {
	// notest
	return newStatement().Newtype()
}

// Newtype renders the `newtype` keyword.
func (g *Group) Newtype() *Statement {
	// notest
	s := Newtype()
	g.items = append(g.items, s)
	return s
}

// Newtype renders the `newtype` keyword.
func (s *Statement) Newtype() *Statement {
	// notest
	t := token{
		content: "newtype",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Not renders the `not` keyword.
func Not() *Statement {
	// notest
	return newStatement().Not()
}

// Not renders the `not` keyword.
func (g *Group) Not() *Statement {
	// notest
	s := Not()
	g.items = append(g.items, s)
	return s
}

// Not renders the `not` keyword.
func (s *Statement) Not() *Statement {
	// notest
	t := token{
		content: "not",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Or renders the `or` keyword.
func Or() *Statement {
	// notest
	return newStatement().Or()
}

// Or renders the `or` keyword.
func (g *Group) Or() *Statement {
	// notest
	s := Or()
	g.items = append(g.items, s)
	return s
}

// Or renders the `or` keyword.
func (s *Statement) Or() *Statement {
	// notest
	t := token{
		content: "or",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Override renders the `override` keyword.
func Override() *Statement {
	// notest
	return newStatement().Override()
}

// Override renders the `override` keyword.
func (g *Group) Override() *Statement {
	// notest
	s := Override()
	g.items = append(g.items, s)
	return s
}

// Override renders the `override` keyword.
func (s *Statement) Override() *Statement {
	// notest
	t := token{
		content: "override",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Predicate renders the `predicate` keyword.
func Predicate() *Statement {
	// notest
	return newStatement().Predicate()
}

// Predicate renders the `predicate` keyword.
func (g *Group) Predicate() *Statement {
	// notest
	s := Predicate()
	g.items = append(g.items, s)
	return s
}

// Predicate renders the `predicate` keyword.
func (s *Statement) Predicate() *Statement {
	// notest
	t := token{
		content: "predicate",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Private renders the `private` keyword.
func Private() *Statement {
	// notest
	return newStatement().Private()
}

// Private renders the `private` keyword.
func (g *Group) Private() *Statement {
	// notest
	s := Private()
	g.items = append(g.items, s)
	return s
}

// Private renders the `private` keyword.
func (s *Statement) Private() *Statement {
	// notest
	t := token{
		content: "private",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Query renders the `query` keyword.
func Query() *Statement {
	// notest
	return newStatement().Query()
}

// Query renders the `query` keyword.
func (g *Group) Query() *Statement {
	// notest
	s := Query()
	g.items = append(g.items, s)
	return s
}

// Query renders the `query` keyword.
func (s *Statement) Query() *Statement {
	// notest
	t := token{
		content: "query",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Result renders the `result` keyword.
func Result() *Statement {
	// notest
	return newStatement().Result()
}

// Result renders the `result` keyword.
func (g *Group) Result() *Statement {
	// notest
	s := Result()
	g.items = append(g.items, s)
	return s
}

// Result renders the `result` keyword.
func (s *Statement) Result() *Statement {
	// notest
	t := token{
		content: "result",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Then renders the `then` keyword.
func Then() *Statement {
	// notest
	return newStatement().Then()
}

// Then renders the `then` keyword.
func (g *Group) Then() *Statement {
	// notest
	s := Then()
	g.items = append(g.items, s)
	return s
}

// Then renders the `then` keyword.
func (s *Statement) Then() *Statement {
	// notest
	t := token{
		content: "then",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// This renders the `this` keyword.
func This() *Statement {
	// notest
	return newStatement().This()
}

// This renders the `this` keyword.
func (g *Group) This() *Statement {
	// notest
	s := This()
	g.items = append(g.items, s)
	return s
}

// This renders the `this` keyword.
func (s *Statement) This() *Statement {
	// notest
	t := token{
		content: "this",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}
