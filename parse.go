package flags

import (
	"os"
	"strings"
)

func (f *Flags) Parse() ([]string, error) {
	unparsed := make([]string, 0)

	args := os.Args[1:]

	var skip bool
	// Looping through all of the arguments provided to the cli
	for i := 0; i < len(args); i++ {
		// Determining if we need to skip the value
		if skip {
			skip = false
		} else {
			// Determining if our current arg is a flag
			if isFlag(args[i]) {
				// It is a flag, so chopping off the dashes
				name := strings.TrimLeft(args[i], "-")
				// Checking to see if it's a set flag
				if f.isSet(name) {
					if f.Flags[name].Class == "bool" {
						f.Flags[name].Value = true
						continue
					} else {
						if isFlag(args[i+1]) {
							return nil, &NoValueDefinedForFlag{Name: name}
						}
						f.Flags[name].Value = args[i+1]
						skip = true
						continue
					}
				}
			}
		}
		unparsed = append(unparsed, args[i])
	}
	return unparsed, nil
}

func isFlag(input string) bool {
	return strings.Index(input, "-") == 0
}
