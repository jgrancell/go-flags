// Package flags implements command line flag parsing
//
// Source code and other details for the project are available at Github
//
//   https://github.com/jgrancell/go-flags
//
package flags

type Flags struct {
	Flags map[string]Flag
}

type Metadata struct {
	Name        string
	Description string
	Class       string
}

type Flag interface {
	attributes() *Metadata
	setValue(interface{})
}

func Init() *Flags {
	return &Flags{Flags: make(map[string]Flag)}
}

func (f *Flags) Add(flag Flag) error {
	metadata := flag.attributes()
	if f.isDupe(metadata) {
		return &DuplicateFlag{Name: metadata.Name}
	}
	if f.isSet(metadata.Name) {
		return &FlagExists{Name: metadata.Name}
	}

	f.Flags[metadata.Name] = flag
	return nil
}
