package internal

import (
	"fmt"
	"strings"
)

type Attributes []string

func (attrs *Attributes) String() string {
	return fmt.Sprint(*attrs)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (attrs *Attributes) Set(attributes string) error {
	if len(*attrs) > 0 {
		return fmt.Errorf("attributes flag already set")
	}

	for _, attr := range strings.Split(attributes, ",") {
		*attrs = append(*attrs, attr)
	}

	return nil
}
