package internal

import (
	"fmt"
	"runtime"
)

//nolint:gochecknoglobals
var (
	version = "dev"
	commit  = "?"
	date    = "?"
)

func BuildVersionString() string {
	var result = fmt.Sprintf("version: %s", version)

	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}

	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}

	result = fmt.Sprintf("%s\nusing: %s", result, runtime.Version())

	return result
}
