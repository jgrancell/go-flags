package flags

func (f *Flags) isSet(name string) bool {
	if len(f.Flags) >= 1 {
		_, isset := f.Flags[name]
		if isset {
			return true
		}
	}
	return false
}

func (f *Flags) isDupe(flag *Flag) bool {
	existing, ok := f.Flags[flag.Name]
	if !ok {
		return false
	}

	if existing.Description != flag.Description {
		return false
	}

	if existing.Class != flag.Class {
		return false
	}

	if existing.Default != flag.Default {
		return false
	}
	return true
}
