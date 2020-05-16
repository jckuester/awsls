package resource

import (
	"github.com/apex/log"
	"github.com/gobwas/glob"
)

// IsType returns true if the given string is a Terraform AWS resource type.
func IsType(s string) bool {
	for _, t := range Types {
		if t == s {
			return true
		}
	}

	return false
}

// IsSupportedType returns true if the given string is an awsls supported resource type.
func IsSupportedType(s string) bool {
	for _, t := range SupportedTypes {
		if t == s {
			return true
		}
	}

	return false
}

// MatchResourceTypes returns all by awsls supported resource types that match the given glob pattern.
func MatchSupportedTypes(globPattern string) ([]string, error) {
	var result []string

	compiledRegex, err := glob.Compile(globPattern)
	if err != nil {
		return nil, err
	}

	for _, rType := range Types {
		if compiledRegex.Match(rType) {
			if !IsSupportedType(rType) {
				log.Debugf("resource type not (yet) supported: %s", rType)

				continue
			}

			result = append(result, rType)
		}
	}

	return result, nil
}

// SupportsTags returns true if the given resource type supports tags.
func SupportsTags(s string) bool {
	for _, t := range TypesWithTags {
		if t == s {
			return true
		}
	}

	return false
}

