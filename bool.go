package flags

type BoolFlag struct {
	Value    *bool
	Metadata *Metadata
}

func (f *Flags) Bool(name string, value bool, description string) (*bool, error) {
	result := &value

	m := &Metadata{Name: name, Description: description, Class: "bool"}

	b := &BoolFlag{
		Value:    result,
		Metadata: m,
	}
	err := f.Add(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *BoolFlag) attributes() *Metadata {
	return b.Metadata
}

func (b *BoolFlag) setValue(value interface{}) {
	*b.Value = value.(bool)
}

func (b *BoolFlag) Get() bool {
	return *b.Value
}
