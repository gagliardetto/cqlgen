// Package jen is a code generator for Go
package jen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
)

// Code represents an item of code that can be rendered.
type Code interface {
	render(f *File, w io.Writer, s *Statement) error
	isNull(f *File) bool
}

// Save renders the file and saves to the filename provided.
func (f *File) Save(filename string) error {
	// notest
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

// Render renders the file to the provided writer.
func (f *File) Render(w io.Writer) error {
	body := &bytes.Buffer{}
	if err := f.render(f, body, nil); err != nil {
		return err
	}
	source := &bytes.Buffer{}
	if len(f.headers) > 0 {
		if err := Doc(f.headers...).render(f, source, nil); err != nil {
			return err
		}
		if _, err := fmt.Fprint(source, "\n"); err != nil {
			return err
		}
		// Append an extra newline so that header comments don't get lumped in
		// with package comments.
		if _, err := fmt.Fprint(source, "\n"); err != nil {
			return err
		}
	}
	for _, c := range f.comments {
		if err := Comment(c).render(f, source, nil); err != nil {
			return err
		}
		if _, err := fmt.Fprint(source, "\n"); err != nil {
			return err
		}
	}
	if err := f.renderImports(source); err != nil {
		return err
	}
	if _, err := source.Write(body.Bytes()); err != nil {
		return err
	}
	// TODO: format using codeql cli.
	if _, err := w.Write(source.Bytes()); err != nil {
		return err
	}
	return nil

	formatted, err := format.Source(source.Bytes())
	if err != nil {
		return fmt.Errorf("Error %s while formatting source:\n%s", err, source.String())
	}
	if _, err := w.Write(formatted); err != nil {
		return err
	}
	return nil
}

func (f *File) renderImports(source io.Writer) error {

	for path := range f.imports {
		if _, err := fmt.Fprintf(source, "import %s\n", path); err != nil {
			return err
		}
	}

	return nil
}
