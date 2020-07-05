package internal

import (
	"fmt"
	"strings"
)

type Attributes []string

func (a *Attributes) String() string {
	return fmt.Sprint(*a)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (a *Attributes) Set(attributes string) error {
	if len(*a) > 0 {
		return fmt.Errorf("attributes flag already set")
	}

	for _, attr := range strings.Split(attributes, ",") {
		*a = append(*a, attr)
	}

	return nil
}

func (a *Attributes) Type() string {
	return "stringSlice"
}
