package resource

import "github.com/gobwas/glob"

// IsType returns true if the given string s is a Terraform AWS resource type.
func IsType(s string) bool {
	for _, t := range Types {
		if t == s {
			return true
		}
	}

	return false
}

// MatchResourceTypes return all Terraform AWS resource types that match the given glob pattern.
func MatchTypes(regex string) []string {
	result := []string{}

	compiledRegex, _ := glob.Compile(regex)

	for _, rType := range Types {
		if compiledRegex.Match(rType) {
			result = append(result, rType)
		}
	}

	return result
}
