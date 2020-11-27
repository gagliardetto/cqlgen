package jen

import (
	"io"
)

// Doc adds a QLDoc comment.
func Doc(str ...string) *Statement {
	return newStatement().Doc(str...)
}

// Doc adds a QLDoc comment.
func (g *Group) Doc(str ...string) *Statement {
	s := Doc(str...)
	g.items = append(g.items, s)
	return s
}

// Doc adds a QLDoc comment.
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
	//if len(c.doc) == 1 {
	//	if _, err := w.Write([]byte("/** " + c.doc[0] + " */")); err != nil {
	//		return err
	//	}
	//	return nil
	//}

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
