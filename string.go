package flags

type StringFlag struct {
	Value    *string
	Metadata *Metadata
}

func (f *Flags) String(name string, value string, description string) (*string, error) {
	result := &value

	m := &Metadata{Name: name, Description: description, Class: "string"}

	s := &StringFlag{
		Value:    result,
		Metadata: m,
	}
	err := f.Add(s)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *StringFlag) attributes() *Metadata {
	return s.Metadata
}

func (s *StringFlag) setValue(value interface{}) {
	*s.Value = value.(string)
}

func (s *StringFlag) Get() string {
	return *s.Value
}
