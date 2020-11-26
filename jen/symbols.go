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
