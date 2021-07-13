package flags

type Int64Flag struct {
	Value    *int64
	Metadata *Metadata
}

func (f *Flags) Int64(name string, value int64, description string) (*int64, error) {
	result := &value

	m := &Metadata{Name: name, Description: description, Class: "int64"}

	i := &Int64Flag{
		Value:    result,
		Metadata: m,
	}
	err := f.Add(i)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Int64Flag) attributes() *Metadata {
	return s.Metadata
}

func (i *Int64Flag) setValue(value interface{}) {
	*i.Value = value.(int64)
}

func (i *Int64Flag) Get() int64 {
	return *i.Value
}
