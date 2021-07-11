// Package flags implements command line flag parsing
//
// Source code and other details for the project are available at Github
//
//   https://github.com/jgrancell/go-flags
//
package flags

type Flags struct {
	Flags map[string]*Flag
}

type Flag struct {
	Default     interface{}
	Name        string
	Description string
	Class       string
	Value       interface{}
}

func Init() *Flags {
	return &Flags{Flags: make(map[string]*Flag)}
}

func (f *Flags) Bool(name string, def bool, description string) {
	f.Add(name, def, description, "bool")
}

func (f *Flags) String(name string, def string, description string) {
	f.Add(name, def, description, "string")
}

func (f *Flags) Int(name string, def int, description string) {
	f.Add(name, def, description, "int")
}

func (f *Flags) Int64(name string, def int64, description string) {
	f.Add(name, def, description, "int64")
}

func (f *Flags) GetValue(name string) (interface{}, error) {
	if f.isSet(name) {
		return f.Flags[name].Value, nil
	}
	return nil, &FlagNotExist{Name: name}
}

func (f *Flags) Add(name string, def interface{}, description string, my_type string) error {
	flag := &Flag{
		Name:        name,
		Description: description,
		Class:       my_type,
	}

	switch my_type {
	case "bool":
		flag.Default = def.(bool)
	case "string":
		flag.Default = def.(string)
	case "int":
		flag.Default = def.(int)
	case "int64":
		flag.Default = def.(int64)
	default:
		flag.Default = def.(string)
	}

	if f.isDupe(flag) {
		return &DuplicateFlag{Name: flag.Name}
	}
	if f.isSet(flag.Name) {
		return &FlagExists{Name: flag.Name}
	}
	f.Flags[flag.Name] = flag
	return nil
}
