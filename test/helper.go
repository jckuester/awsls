// Package test contains acceptance tests.
package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	// tfstateBucket is the S3 bucket that stores all Terraform states of acceptance tests
	tfstateBucket = "awsls-testacc-tfstate-492043"
	// profile1 is used as default profile throughout the tests if not overwritten by envs
	profile1 = "myaccount1"
	// region1 is used as default region throughout the tests if not overwritten by envs
	region1 = "us-west-2"

	profile2 = "myaccount2"
	region2  = "us-east-1"
)

// EnvVars contains environment variables for that must be set for tests.
type EnvVars struct {
	AWSRegion  string
	AWSProfile string
}

// InitEnv sets environment variables for acceptance tests.
func InitEnv(t *testing.T) EnvVars {
	t.Helper()

	profile := os.Getenv("AWS_PROFILE")
	if profile == "" {
		profile = profile1

		t.Logf("env variable AWS_PROFILE not set, using the following profile for tests: %s", profile1)
	}

	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		region = region1

		t.Logf("env variable AWS_DEFAULT_REGION not set, using the following region for tests: %s", region1)
	}

	return EnvVars{
		AWSProfile: profile,
		AWSRegion:  region,
	}
}

func GetTerraformOptions(terraformDir string, env EnvVars, overrideVars ...map[string]interface{}) *terraform.Options {
	name := "awsls-testacc-" + strings.ToLower(random.UniqueId())

	vars := map[string]interface{}{
		"region":  env.AWSRegion,
		"profile": env.AWSProfile,
		"name":    name,
	}

	if len(overrideVars) > 0 {
		vars = overrideVars[0]
	}

	return &terraform.Options{
		TerraformDir: terraformDir,
		NoColor:      true,
		Vars:         vars,
		// BackendConfig defines where to store the Terraform state files of tests
		BackendConfig: map[string]interface{}{
			"bucket":  tfstateBucket,
			"key":     fmt.Sprintf("%s.tfstate", name),
			"region":  env.AWSRegion,
			"profile": env.AWSProfile,
			"encrypt": true,
		},
	}
}

func AssertVpcExists(t *testing.T, actualVpcID string, env EnvVars) {
	_, err := aws.GetVpcByIdE(t, actualVpcID, env.AWSRegion)
	assert.NoError(t, err, "resource has been unexpectedly deleted")
}

func SetMultiEnvs(variables map[string]string) error {
	for envName, envValue := range variables {
		err := os.Setenv(envName, envValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func UnsetAWSEnvs() error {
	return UnsetMultiEnvs([]string{
		"AWS_DEFAULT_REGION",
		"AWS_REGION",
		"AWS_PROFILE",
		"AWS_SDK_LOAD_CONFIG",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
	})
}

func UnsetMultiEnvs(variables []string) error {
	for _, envName := range variables {
		err := os.Unsetenv(envName)
		if err != nil {
			return err
		}
	}

	return nil
}
