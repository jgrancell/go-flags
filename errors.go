package flags

type FlagExists struct {
	Name string
}

func (f *FlagExists) Error() string {
	return "flag " + f.Name + " is already defined"
}

type DuplicateFlag struct {
	Name string
}

func (d *DuplicateFlag) Error() string {
	return "flag " + d.Name + " is an exact duplicate of an already-defined flag"
}

type NoValueDefinedForFlag struct {
	Name string
}

func (n *NoValueDefinedForFlag) Error() string {
	return "you must provide a value for argument " + n.Name
}

type FlagNotExist struct {
	Name string
}

func (n *FlagNotExist) Error() string {
	return "flag " + n.Name + " is not defined"
}
