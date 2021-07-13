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

func (f *Flags) isDupe(metadata *Metadata) bool {
	existing, ok := f.Flags[metadata.Name]
	if !ok {
		return false
	}

	existing_metadata := existing.attributes()
	return existing_metadata.Description == metadata.Description
}
