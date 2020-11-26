package jen

import (
	"io"
)

// Doc adds a doc. If the provided string contains a newline, the
// doc is formatted in multiline style. If the doc string starts
// with "//" or "/*", the automatic formatting is disabled and the string is
// rendered directly.
func Doc(str ...string) *Statement {
	return newStatement().Doc(str...)
}

// Doc adds a doc. If the provided string contains a newline, the
// doc is formatted in multiline style. If the doc string starts
// with "//" or "/*", the automatic formatting is disabled and the string is
// rendered directly.
func (g *Group) Doc(str ...string) *Statement {
	s := Doc(str...)
	g.items = append(g.items, s)
	return s
}

// Doc adds a doc. If the provided string contains a newline, the
// doc is formatted in multiline style. If the doc string starts
// with "//" or "/*", the automatic formatting is disabled and the string is
// rendered directly.
func (s *Statement) Doc(str ...string) *Statement {
	c := doc{
		doc: str,
	}
	*s = append(*s, c)
	return s
}

type doc struct {
	doc []string
}

func (c doc) isNull(f *File) bool {
	return false
}

func (c doc) render(f *File, w io.Writer, s *Statement) error {
	if _, err := w.Write([]byte("/**\n")); err != nil {
		return err
	}
	for _, line := range c.doc {
		if _, err := w.Write([]byte(" * " + line + "\n")); err != nil {
			return err
		}
	}
	if _, err := w.Write([]byte(" */")); err != nil {
		return err
	}
	return nil
}
