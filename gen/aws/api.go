// +build codegen

package aws

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/private/model/api"
)

// APIs returns the AWS API models loaded by API package name.
func APIs(awsSdkRepoPath string) (api.APIs, error) {
	globs := []string{fmt.Sprintf("%s/models/apis/*/*/api-2.json", awsSdkRepoPath)}

	modelPaths, err := api.ExpandModelGlobPath(globs...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to glob file pattern")
	}
	modelPaths, _ = api.TrimModelServiceVersions(modelPaths)

	loader := api.Loader{
		//BaseImport:          outputPath,
		KeepUnsupportedAPIs: false,
	}

	apis, err := loader.Load(modelPaths)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load API models")
	}
	if len(apis) == 0 {
		return nil, fmt.Errorf("no API models found")
	}

	delete(apis, "github.com/aws/aws-sdk-go-v2/service/importexport")

	return apis, nil
}
