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

// Eq adds =.
func Eq() *Statement {
	return newStatement().Op("=")
}

func (g *Group) Eq() *Statement {
	return g.Add(Eq())
}

func (s *Statement) Eq() *Statement {
	return s.Add(Eq())
}

// Neq adds !=.
func Neq() *Statement {
	return newStatement().Op("!=")
}

func (g *Group) Neq() *Statement {
	return g.Add(Neq())
}

func (s *Statement) Neq() *Statement {
	return s.Add(Neq())
}

// Lt adds <.
func Lt() *Statement {
	return newStatement().Op("<")
}

func (g *Group) Lt() *Statement {
	return g.Add(Lt())
}

func (s *Statement) Lt() *Statement {
	return s.Add(Lt())
}

// Gt adds >.
func Gt() *Statement {
	return newStatement().Op(">")
}

func (g *Group) Gt() *Statement {
	return g.Add(Gt())
}

func (s *Statement) Gt() *Statement {
	return s.Add(Gt())
}

// Lte adds <=.
func Lte() *Statement {
	return newStatement().Op("<=")
}

func (g *Group) Lte() *Statement {
	return g.Add(Lte())
}

func (s *Statement) Lte() *Statement {
	return s.Add(Lte())
}

// Gte adds >=.
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
