// +build codegen

package aws

import (
	"sort"

	"github.com/aws/aws-sdk-go-v2/private/model/api"
)

// ServicePkgNames returns the package name of all AWS services.
func ServicePkgNames(apis api.APIs) []string {
	var result []string
	for _, a := range apis {
		result = append(result, a.PackageName())
	}

	return result
}

// ServicesCoveredByTerraform returns the package name of all AWS services that are
// (at least partially) covered by Terraform resources.
func ServicesCoveredByTerraform(serviceMap map[string]string) []string {
	serviceSet := make(map[string]bool)

	for _, v := range serviceMap {
		serviceSet[v] = true
	}

	var result []string
	for k, _ := range serviceSet {
		result = append(result, k)
	}

	sort.Strings(result)

	return result
}
