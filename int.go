package flags

type IntFlag struct {
	Value    *int
	Metadata *Metadata
}

func (f *Flags) Int(name string, value int, description string) (*int, error) {
	result := &value

	m := &Metadata{Name: name, Description: description, Class: "int"}

	i := &IntFlag{
		Value:    result,
		Metadata: m,
	}
	err := f.Add(i)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *IntFlag) attributes() *Metadata {
	return i.Metadata
}

func (i *IntFlag) setValue(value interface{}) {
	*i.Value = value.(int)
}

func (i *IntFlag) Get() int {
	return *i.Value
}
