// Package jen is a code generator for Go
package jen

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
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

	// If codeql is installed, then use it to format the source:
	isInstalled := CodeqlCliIsInstalled()
	if isInstalled {
		// Format the codeql source code:
		formatted, err := FormatCodeQL(source.Bytes())
		if err != nil {
			return err
		}
		// Write formatted source to destination:
		if _, err := w.Write(formatted); err != nil {
			return err
		}
		return nil
	}
	if _, err := w.Write(source.Bytes()); err != nil {
		return err
	}
	return nil
}

func CodeqlCliIsInstalled() bool {
	return commandExists("codeql")
}

var ErrCodeQLCLINotInstalled = errors.New("codeql cli is not installed")

// FormatCodeQL will format the provided codeql code if the
// codeql CLI is installed. If the codeql CLI is NOT installed,
// then the function will return the error ErrCodeQLCLINotInstalled.
func FormatCodeQL(src []byte) ([]byte, error) {
	// If codeql is installed, then use it to format the source:
	isInstalled := CodeqlCliIsInstalled()
	if !isInstalled {
		return nil, ErrCodeQLCLINotInstalled
	}

	// Create a temporary file:
	tmpFile, err := ioutil.TempFile("", "cqlgen.*.ql")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Write source to the temporary file:
	if _, err := tmpFile.Write(src); err != nil {
		return nil, err
	}

	// Get absolute path to temporary file:
	tmpFileAbsPath, err := filepath.Abs(tmpFile.Name())
	if err != nil {
		return nil, err
	}

	// Format temporary file:
	output, err := exec.Command("codeql", "query", "format", "-qq", "-i", tmpFileAbsPath).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

	// Read formatted file:
	formatted, err := ioutil.ReadFile(tmpFileAbsPath)
	if err != nil {
		return nil, err
	}
	return formatted, nil
}

// Check whether a command exists.
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func (f *File) renderImports(source io.Writer) error {
	// Get keys (paths) and sort them:
	keys := func(v map[string]importdef) []string {
		res := make([]string, 0)
		for key := range v {
			res = append(res, key)
		}
		sort.Strings(res)
		return res
	}(f.imports)

	for _, path := range keys {
		imp := f.imports[path]
		if imp.as != "" {
			if _, err := fmt.Fprintf(source, "import %s as %s\n", path, imp.as); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(source, "import %s\n", path); err != nil {
				return err
			}
		}
	}

	return nil
}
