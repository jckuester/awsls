package resource

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/gobwas/glob"
	"github.com/jckuester/awstools-lib/terraform"
	"github.com/jckuester/terradozer/pkg/provider"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
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
		if compiledRegex.Match(rType) ||
			// allow user to provide also resource types without prefix (e.g., aws_iam_role and iam_role)
			compiledRegex.Match(strings.TrimPrefix(rType, "aws_")) {
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

// HasAttributes returns only the attributes that the given Terraform resource type supports out of a given
// list of attributes.
func HasAttributes(attributes []string, terraformType string, provider *provider.TerraformProvider) (map[string]bool, error) {
	schema, err := provider.GetSchemaForResource(terraformType)
	if err != nil {
		return nil, err
	}

	result := map[string]bool{}

	for _, attr := range attributes {
		_, ok := schema.Block.Attributes[attr]
		if ok {
			result[attr] = true
		}
	}

	return result, nil
}

// GetAttribute returns any Terraform attribute of a resource by name.
func GetAttribute(name string, r *terraform.Resource) (string, error) {
	if r.UpdatableResource == nil {
		return "", fmt.Errorf("resource is nil")
	}

	state := r.State()

	if state == nil {
		return "", fmt.Errorf("state is nil")
	}

	if state.IsNull() {
		return "", fmt.Errorf("state is nil value")
	}

	if !state.IsKnown() {
		return "", fmt.Errorf("state is unknown")
	}

	if !state.IsWhollyKnown() {
		return "", fmt.Errorf("state is not wholly known")
	}

	if !state.CanIterateElements() {
		return "", fmt.Errorf("cannot iterate: %s", *state)
	}

	attrValue, ok := state.AsValueMap()[name]
	if !ok {
		return "", fmt.Errorf("attribute not found: %s", name)
	}

	switch attrValue.Type() {
	case cty.Bool:
		var v bool
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		return strconv.FormatBool(v), nil

	case cty.Number:
		var v int64
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		return strconv.FormatInt(v, 10), nil

	case cty.String:
		var v string
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		return v, nil

	case cty.Map(cty.String):
		var v map[string]string
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		if len(v) > 0 {
			var list []string

			for k, v := range v {
				list = append(list, fmt.Sprintf("%s=%s", k, v))
			}

			sort.Strings(list)

			return strings.Join(list, ","), nil
		}

		return "", nil
	default:
		return "", fmt.Errorf("currently unhandled type: %s", attrValue.Type().FriendlyName())
	}
}
