package jen

import (
	"bytes"
)

// NewFile Creates a new file, with the specified package name.
func NewFile() *File {
	return &File{
		Group: &Group{
			multi: true,
		},
		imports: map[string]importdef{},
	}
}

// File represents a single source file. Package imports are managed
// automaticaly by File.
type File struct {
	*Group
	path     string
	imports  map[string]importdef
	comments []string
	headers  []string
}

// importdef is used to differentiate packages where we know the package name from packages where the
// import is aliased. If alias == false, then name is the actual package name, and the import will be
// rendered without an alias. If used == false, the import has not been used in code yet and should be
// excluded from the import block.
type importdef struct {
	name string
}

// HeaderDoc adds a comment to the top of the file, above any package
// comments. A blank line is rendered below the header comments, ensuring
// header comments are not included in the package doc.
func (f *File) HeaderDoc(comment string) {
	f.headers = append(f.headers, comment)
}

// PackageComment adds a comment to the top of the file, above the package
// keyword.
func (f *File) PackageComment(comment string) {
	f.comments = append(f.comments, comment)
}

// ImportNames allows multiple names to be imported as a map. Use the [gennames](gennames) command to
// automatically generate a go file containing a map of a selection of package names.
func (f *File) ImportNames(names []string) {
	for _, name := range names {
		f.imports[name] = importdef{name: name}
	}
}

// ImportAlias provides the alias for a package path that should be used in the import block. A
// period can be used to force a dot-import.
func (f *File) Import(path string) {
	f.imports[path] = importdef{name: path}
}

func (f *File) isLocal(path string) bool {
	return f.path == path
}

func (f *File) isValidAlias(alias string) bool {
	// multiple dot-imports are ok
	if alias == "." {
		return true
	}
	// the import alias is invalid if it's a reserved word
	if IsReservedWord(alias) {
		return false
	}
	// the import alias is invalid if it's already been registered
	for _, v := range f.imports {
		if alias == v.name {
			return false
		}
	}
	return true
}

func (f *File) register(path string) string {
	// Register the eventual name
	f.imports[path] = importdef{name: path}
	return path
}

// GoString renders the File for testing. Any error will cause a panic.
func (f *File) GoString() string {
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	return buf.String()
}
