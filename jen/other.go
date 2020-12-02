package jen

// Bar adds a vertical bar.
func Bar() *Statement {
	return newStatement().Op("|")
}

func (g *Group) Bar() *Statement {
	return g.Add(Bar())
}

func (s *Statement) Bar() *Statement {
	return s.Add(Bar())
}

// Eq adds a `=` comparison operator.
func Eq() *Statement {
	return newStatement().Op("=")
}

func (g *Group) Eq() *Statement {
	return g.Add(Eq())
}

func (s *Statement) Eq() *Statement {
	return s.Add(Eq())
}

// Neq adds a `!=` comparison operator.
func Neq() *Statement {
	return newStatement().Op("!=")
}

func (g *Group) Neq() *Statement {
	return g.Add(Neq())
}

func (s *Statement) Neq() *Statement {
	return s.Add(Neq())
}

// Lt adds a `<` comparison operator.
func Lt() *Statement {
	return newStatement().Op("<")
}

func (g *Group) Lt() *Statement {
	return g.Add(Lt())
}

func (s *Statement) Lt() *Statement {
	return s.Add(Lt())
}

// Gt adds a `>` comparison operator.
func Gt() *Statement {
	return newStatement().Op(">")
}

func (g *Group) Gt() *Statement {
	return g.Add(Gt())
}

func (s *Statement) Gt() *Statement {
	return s.Add(Gt())
}

// Lte adds a `<=` comparison operator.
func Lte() *Statement {
	return newStatement().Op("<=")
}

func (g *Group) Lte() *Statement {
	return g.Add(Lte())
}

func (s *Statement) Lte() *Statement {
	return s.Add(Lte())
}

// Gte adds a `>=` comparison operator.
func Gte() *Statement {
	return newStatement().Op(">=")
}

func (g *Group) Gte() *Statement {
	return g.Add(Gte())
}

func (s *Statement) Gte() *Statement {
	return s.Add(Gte())
}

// None adds none().
func None() *Statement {
	return newStatement().Id("none").Call()
}

func (g *Group) None() *Statement {
	return g.Add(None())
}

func (s *Statement) None() *Statement {
	return s.Add(None())
}

// From adds a `from` clause.
func From(vars ...Code) *Statement {
	return newStatement().Line().Id("from").List(vars...).Line()
}

func (g *Group) From(vars ...Code) *Statement {
	return g.Add(From(vars...))
}

func (s *Statement) From(vars ...Code) *Statement {
	return s.Add(From(vars...))
}

// Where adds a `where` clause.
func Where(formula ...Code) *Statement {
	return newStatement().Line().Id("where").Add(formula...).Line()
}

func (g *Group) Where(formula ...Code) *Statement {
	return g.Add(Where(formula...))
}

func (s *Statement) Where(formula ...Code) *Statement {
	return s.Add(Where(formula...))
}

// Select adds a `select` clause.
func Select(exprs ...Code) *Statement {
	return newStatement().Line().Id("select").List(exprs...).Line()
}

func (g *Group) Select(exprs ...Code) *Statement {
	return g.Add(Select(exprs...))
}

func (s *Statement) Select(exprs ...Code) *Statement {
	return s.Add(Select(exprs...))
}

// BindingSet adds a `bindingset` annotation.
func BindingSet(names ...string) *Statement {

	code := make([]Code, 0)
	for _, name := range names {
		code = append(code, Id(name))
	}

	return newStatement().Id("bindingset").Id("bindingset").Set(code...)
}

func (g *Group) BindingSet(set string) *Statement {
	return g.Add(BindingSet(set))
}

func (s *Statement) BindingSet(set string) *Statement {
	return s.Add(BindingSet(set))
}

// Semicolon adds a semicolon.
func Semicolon() *Statement {
	return newStatement().Op(";")
}

func (g *Group) Semicolon() *Statement {
	return g.Add(Semicolon())
}

func (s *Statement) Semicolon() *Statement {
	return s.Add(Semicolon())
}

// OrderBy adds an `order by`.
func OrderBy() *Statement {
	return newStatement().Id("order by")
}

func (g *Group) OrderBy() *Statement {
	return g.Add(OrderBy())
}

func (s *Statement) OrderBy() *Statement {
	return s.Add(OrderBy())
}

// DontCare adds a dontcare expression, represented by a  `_`.
func DontCare() *Statement {
	return newStatement().Op("_")
}

func (g *Group) DontCare() *Statement {
	return g.Add(DontCare())
}

func (s *Statement) DontCare() *Statement {
	return s.Add(DontCare())
}

// Range adds a range expression; e.g. [-3 .. 10].
// The parameters must be literals of int, float, or date type.
func Range(from Code, to Code) *Statement {
	return newStatement().Op("[").Add(from).Op("..").Add(to).Op("]")
}

func (g *Group) Range(from Code, to Code) *Statement {
	return g.Add(Range(from, to))
}

func (s *Statement) Range(from Code, to Code) *Statement {
	return s.Add(Range(from, to))
}

// Join joins code with a separator.
// Before joining, it removes null elements.
func Join(separator Code, elems ...Code) *Statement {
	st := newStatement()

	notNullElems := make([]Code, 0)
	for _, elem := range elems {
		if elem.isNull(nil) {
			continue
		}
		notNullElems = append(notNullElems, elem)
	}

	for index, elem := range notNullElems {
		if index > 0 {
			st.Add(separator).Add(elem)
		} else {
			st.Add(elem)
		}
	}
	return st
}

func (g *Group) Join(separator Code, elems ...Code) *Statement {
	return g.Add(Join(separator, elems...))
}

func (s *Statement) Join(separator Code, elems ...Code) *Statement {
	return s.Add(Join(separator, elems...))
}

func IntsToSet(elems ...int) *Statement {
	lits := make([]Code, 0)
	for _, elem := range elems {
		lits = append(lits, Lit(elem))
	}

	return Set(lits...)
}
func StringsToSet(elems ...string) *Statement {
	lits := make([]Code, 0)
	for _, elem := range elems {
		lits = append(lits, Lit(elem))
	}

	return Set(lits...)
}
